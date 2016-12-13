package api

import (
	"github.com/svarlamov/bintrad/context"
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

func V0_API_Get_My_User_Data(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetUserAndCatch(w, r)
	if err != nil {
		Log.WithField("error", err).Error("Error getting user from context")
		return
	}

	completeUser := models.CompleteUser{}
	err = completeUser.PopulateFromUser(user)
	if err != nil {
		utils.JSONInternalError(w, "Error getting user data", "")
		return
	}

	utils.JSONSuccess(w, completeUser, "Successfully got user data")
}
