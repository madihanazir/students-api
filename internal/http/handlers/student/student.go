package student

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/madihanazir/students-api/internal/types"
	"github.com/madihanazir/students-api/internal/utils/response"
	"github.com/madihanazir/students-api/storage"
	_ "github.com/madihanazir/students-api/storage/sqlite"
)

// New returns an HTTP handler function for the students API
// here we use dependency injection to pass the storage interface
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating new student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		//request validation

		validate := validator.New()
		if err := validate.Struct(student); err != nil {
			validateErrs, ok := err.(validator.ValidationErrors)

			if ok {
				response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			} else {
				response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			}
			return
		}

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))

			return
		}

		slog.Info("user created", slog.String("userId", fmt.Sprint(lastId)))

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// id := r.PathValue("id")
		// slog.Info("getting student by id", slog.String("id", id))
		// intId, err := (strconv.ParseInt(id, 10, 64))
		// if err != nil {
		// 	response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
		// 	return
		// }
		// student, err := storage.GetStudentById(intId)
		// if err != nil {
		// 	slog.Error("error getting student by id", slog.String("id", id))
		// 	response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
		// 	return
		// }
		// response.WriteJSON(w, http.StatusOK, student)

		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok || idStr == "" {
			http.Error(w, `{"status": "StatusError", "error": "missing id"}`, http.StatusBadRequest)
			return
		}

		// Convert ID to integer
		id, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			http.Error(w, `{"status": "StatusError", "error": "invalid id format"}`, http.StatusBadRequest)
			return
		}

		slog.Info("getting student by id", slog.Int64("id", id))

		// Fetch student from storage
		student, err := storage.GetStudentByID(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, `{"status": "StatusError", "error": "student not found"}`, http.StatusNotFound)
				return
			}
			http.Error(w, `{"status": "StatusError", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}

		// Return student as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(student)
	}
}
func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting list of students")
		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, students)
	}
}

func UpdateStudent(storage storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.ParseInt(vars["id"], 10, 64)
        if err != nil {
            response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id format")))
            return
        }

        var student types.Student
        err = json.NewDecoder(r.Body).Decode(&student)
        if err != nil {
            response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }

        err = storage.UpdateStudent(id, student.Name, student.Email, student.Age)
        if err != nil {
            response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
            return
        }

        response.WriteJSON(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
    }
}

func PatchStudent(storage storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.ParseInt(vars["id"], 10, 64)
        if err != nil {
            response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id format")))
            return
        }

        var fields map[string]interface{}
        err = json.NewDecoder(r.Body).Decode(&fields)
        if err != nil {
            response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
            return
        }

        err = storage.PatchStudent(id, fields)
        if err != nil {
            response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
            return
        }

        response.WriteJSON(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
    }
}


func DeleteStudent(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok || idStr == "" {
			http.Error(w, `{"error": "missing id"}`, http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, `{"error": "invalid id format"}`, http.StatusBadRequest)
			return
		}

		err = storage.DeleteStudent(id)
		if err != nil {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
		response.WriteJSON(w, http.StatusOK, map[string]string{"message": "Student deleted successfully"})
	}
}

func StudentExists(storage storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.ParseInt(vars["id"], 10, 64)
        if err != nil {
            response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id format")))
            return
        }

        exists, err := storage.StudentExists(id)
        if err != nil {
            response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
            return
        }

        if exists {
            w.WriteHeader(http.StatusOK) // 200 OK if student exists
        } else {
            w.WriteHeader(http.StatusNotFound) // 404 Not Found if student doesn't exist
        }
    }
}


