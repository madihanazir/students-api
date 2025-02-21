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
)

// New returns an HTTP handler function for the students API
func New() http.HandlerFunc {
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

		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "OK"})
		//request validation

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Students API is working"))
	}
}
