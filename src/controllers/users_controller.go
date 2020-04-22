package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"woden/src/auth"

	"woden/src/models"
	"woden/src/responses"
	"woden/src/utils"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, email, password string) (string, error) {
	var err error
	user := models.User{}
	if len(username) == 0 {
		err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	} else {
		err = server.DB.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
	}
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", err
	}
	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return "", err
	}
	user.Token = token
	return token, nil
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()

	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%d", userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
}
