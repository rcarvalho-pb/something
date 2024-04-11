package auth_controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/todo-app-go/internal/authentication"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
	encoder "github.com/rcarvalho-pb/todo-app-go/internal/infra/password_encoder"
	json_response "github.com/rcarvalho-pb/todo-app-go/internal/infra/response"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {
	*sqlx.DB
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	var auth Auth
	if err = json.Unmarshal(requestBody, &auth); err != nil {
		json_response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user user_dto.User
	if err = a.Get(&user, "SELECT * FROM users WHERE email=$1", auth.Username); err != nil {
		json_response.ERROR(w, response_status.NotFound, err)
		return
	}

	if err = encoder.CheckPassword(auth.Password, user.Password); err != nil {
		json_response.ERROR(w, response_status.InternalError, err)
		return
	}

	token, err := authentication.CreateToken(user.ID)
	if err != nil {
		json_response.ERROR(w, response_status.InternalError, err)
		return
	}

	json_response.JSON(w, response_status.Ok, map[string]any{
		"token":  token,
		"userId": user.ID,
	})
}
