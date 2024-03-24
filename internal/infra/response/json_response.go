package json_response

import (
	"encoding/json"
	"log"
	"net/http"

	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
)

func JSON(w http.ResponseWriter, statusCode response_status.StatusCode, data any) {
	w.WriteHeader(int(statusCode))

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ERROR(w http.ResponseWriter, statusCode response_status.StatusCode, err error) {
	JSON(w, statusCode, struct {
		Error string
	}{
		Error: err.Error(),
	})
}
