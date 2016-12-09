package api

import (
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
)

var Log = config.Conf.GetLogger()

func V0_API(w http.ResponseWriter, r *http.Request) {
	utils.JSONSuccess(w, nil, "")
}
