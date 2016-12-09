package setup

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/svarlamov/bintrad/config"
	"github.com/svarlamov/bintrad/models" /* MySQL */
	/* SQLite  _ "github.com/mattn/go-sqlite3"*/)

var Log = config.Conf.GetLogger()

func StartSetup(resetDatabase bool) {
	Log.Info("Server Loaded")

	err := models.Setup()
	if err != nil {
		Log.WithField("error", err).Fatal("Couldn't setup models")
	}

	if resetDatabase {
		Log.Info("Resetting database")
		err = models.WipeDatabase()
		if err != nil {
			Log.WithField("error", err).Fatal("Couldn't wipe database for reset")
		}
		err = models.InitializeDatabaseFromFile(config.Conf.DBSetupSQLScript)
		if err != nil {
			Log.WithField("error", err).Fatal("Couldn't wipe database for reset")
		}
		Log.Info("Finished resetting database")
	}

	setupUsers()
	setupTickers()

}
