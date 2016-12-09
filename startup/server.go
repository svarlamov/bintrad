package startup

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/controllers"
	"github.com/svarlamov/bintrad/models" /* MySQL */
	"github.com/svarlamov/logrus-middleware"
	"net/http"
	/* SQLite  _ "github.com/mattn/go-sqlite3"*/)

var Log = config.Conf.GetLogger()

func StartServer() {
	Log.Info("Server Loaded")
	lm := logrusmiddleware.Middleware{
		Name:   "bintrad",
		Logger: Log,
	}

	err := models.Setup()
	if err != nil {
		Log.WithField("error", err).Fatal("Couldn't setup models")
	}

	Log.WithField("hostname", config.Conf.ApiURL).Info("Server starting")
	http.ListenAndServe(config.Conf.ApiURL, lm.Handler(controllers.CreateRouter(), ""))
}
