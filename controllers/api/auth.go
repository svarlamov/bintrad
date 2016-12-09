package api

import (
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

func V0_API_Authenticate(w http.ResponseWriter, r *http.Request) {
	reqObj := models.AuthRequest{}
	err := utils.JSONDecodeAndCatch(w, r, &reqObj)
	if err != nil {
		return
	}
	user, err := reqObj.Authenticate()
	if err != nil || user.Id == 0 {
		utils.JSONBadRequestError(w, "Invalid username/passkey combination", "")
		return
	}

	accessToken, err := user.GenerateAndPersistAccessToken()
	if err != nil || accessToken.IsValid() == false {
		utils.JSONInternalError(w, "Error creating access token", "")
		return
	}

	utils.SetHTTPOnlyCookie(w, config.Conf.TokenCookieName, accessToken.Token)

	respObj := models.AuthResponse{AccessToken: accessToken.Token}
	utils.JSONSuccess(w, respObj, "Successfully authenticated")
}
