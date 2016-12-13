package context

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/context"
	"github.com/svarlamov/bintrad/models"
	"net/http"
)

const accessTokenContextKey string = "access_token"

type APIResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Debug   string      `json:"debug,omitempty"`
}

var unauthorizedError APIResponse = APIResponse{Message: "Unauthorized", Success: false, Debug: "Token/account error"}

func SetToken(r *http.Request, session models.AccessToken) {
	context.Set(r, accessTokenContextKey, session)
}

func ClearRequest(r *http.Request) {
	context.Clear(r)
}

func GetToken(r *http.Request) (models.AccessToken, error) {
	session := context.Get(r, accessTokenContextKey)
	switch session.(type) {
	case nil:
		return models.AccessToken{}, errors.New("Token not found")
	default:
		return session.(models.AccessToken), nil
	}
}

func GetUserAndCatch(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	tkn := models.AccessToken{}
	tkn, err := GetToken(r)
	if err != nil || tkn.Id == 0 {
		jsonWriter(w, unauthorizedError, http.StatusUnauthorized)
		return user, err
	}
	user = models.User{Id: tkn.UserId}
	err = user.FindById()
	if err != nil || user.Id == 0 {
		jsonWriter(w, unauthorizedError, http.StatusUnauthorized)
		return user, err
	}
	return user, nil
}

func GetUserSilently(r *http.Request) (models.User, error) {
	var user models.User
	tkn := models.AccessToken{}
	tkn, err := GetToken(r)
	if err != nil || tkn.Id == 0 {
		return user, err
	}
	user = models.User{Id: tkn.UserId}
	err = user.FindById()
	if err != nil || user.Id == 0 {
		return user, err
	}
	return user, nil
}

func jsonWriter(w http.ResponseWriter, d interface{}, c int) {
	//dj, err := json.MarshalIndent(d, "", "  ")
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
