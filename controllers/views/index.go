package views

import (
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/context"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

var Log = config.Conf.GetLogger()

func V0_VIEWS_Index(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetUserSilently(r)
	if err != nil || user.Id == 0 {
		utils.RenderSuccessfulTemplateFromFile(w, nil, "templates/login.html")
		return
	}

	utils.TemporaryRedirect(w, r, "/tradingDesk")
}
