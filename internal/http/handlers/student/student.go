package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/madihanazir/students-api/internal/types"
	"github.com/madihanazir/students-api/internal/utils/response"
	"github.com/madihanazir/students-api/storage"
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
