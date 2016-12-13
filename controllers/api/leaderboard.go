package api

import (
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

func V0_API_Get_Leaderboard(w http.ResponseWriter, r *http.Request) {
	leadUsers, err := models.GetCurrentLeaderboard(10)
	if err != nil {
		utils.JSONInternalError(w, "Error getting user data", "")
		return
	}

	utils.JSONSuccess(w, leadUsers, "Successfully got leaderboard")
}
