package main

import (
	"github.com/svarlamov/bintrad/setup"
	"github.com/svarlamov/bintrad/startup"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	startCmd        = kingpin.Command("start", "Start the server").Default()
	setupCmd        = kingpin.Command("setup", "Setup the data in the database")
	resetDBFlag     = setupCmd.Flag("resetdb", "Wipe database and reset tables first").Bool()
	justUsersDBFlag = setupCmd.Flag("usersonly", "Only add users when writing table data").Bool()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Sasha Varlamov")
	kingpin.CommandLine.Help = "BinTrad -- Binary Options Trading Game"
	switch kingpin.Parse() {
	case "start":
		startup.StartServer()
	case "setup":
		setup.StartSetup(*resetDBFlag, *justUsersDBFlag)
	}
}
