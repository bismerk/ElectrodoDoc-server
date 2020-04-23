package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"io/ioutil"
	"net/http"
	"woden/src/auth"
	"woden/src/models"

	"woden/src/responses"
	"woden/src/utils"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	if err := checkmail.ValidateFormat(r.FormValue("login")); err != nil {
		user.Username = r.FormValue("login")
	} else {
		user.Email = r.FormValue("login")
	}
	user.Password = r.FormValue("password")
	user.Prepare()
	err := user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
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
		return "", errors.New(fmt.Sprintf("%s", err))
	}
	if user.Password != password {
		return "", errors.New("Wrong password")
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	user.Token = token
	return token, nil
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	user.Username = r.FormValue("login")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	user.Prepare()

	err := user.Validate("")
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
	w.Header().Set("id", fmt.Sprintf("%d", userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	tokenDB, _ := models.GetToken(token)
	if tokenDB != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
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
	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, tokenId)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	tokenId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	result, err := models.Logout(tokenId, token)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(fmt.Sprintf("%s", err)))
		return
	}

	if result == true {
		r.Header.Del("Authorization")
		responses.JSON(w, http.StatusOK, "Logout successfully")
		return
	}
	if result == false {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Bad Authorization"))
		return
	}
}
