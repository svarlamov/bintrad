package views

import (
	"github.com/svarlamov/bintrad/context"
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

func V0_VIEWS_Trading_Desk(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetUserSilently(r)
	if err != nil || user.Id == 0 {
		utils.TemporaryRedirect(w, r, "/")
		return
	}

	completeUser := models.CompleteUser{}
	err = completeUser.PopulateFromUser(user)
	if err != nil {
		utils.TemporaryRedirect(w, r, "/")
		return
	}

	utils.RenderSuccessfulTemplateFromFile(w, completeUser, "templates/trading_desk.html")

}
