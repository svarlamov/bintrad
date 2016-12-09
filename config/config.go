package config

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/smtp"
)

type Config struct {
	ApiURL           string          `json:"apiUrl"`
	DBPath           string          `json:"dbPath"`
	DBName           string          `json:"dbName"`
	DBSetupSQLScript string          `json:"dbSetupSQLScript"`
	GeoLiteMMDBPath  string          `json:"geoLiteMMDBPath"`
	Home             string          `json:"home"`
	SMTPAuth         smtp.Auth       `json:"-"`
	SMTPCreds        SMTPCredentials `json:"smtpCredentials"`
	TokenCookieName  string          `json:"tokenCookieName"`
	Mode             string          `json:"mode"`
}

type SMTPCredentials struct {
	DisplayName      string `json:"displayName"`
	EmailAddress     string `json:"emailAddress"`
	Password         string `json:"password"`
	Hostname         string `json:"hostname"`
	Port             int64  `json:"port"`
	ConnectionString string `json:"-"`
}

var Conf Config

const (
	DebugMode      string = "debug"
	ProductionMode string = "production"
)

func init() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(fmt.Sprint("Could not parse config:", err))
	}
	err = json.Unmarshal(configFile, &Conf)
	if err != nil {
		panic(fmt.Sprint("Error unmarshalling config:", err))
	}
	Conf.SMTPCreds.ConnectionString = fmt.Sprint(Conf.SMTPCreds.Hostname, ":", Conf.SMTPCreds.Port)

	Conf.SMTPAuth = smtp.PlainAuth("",
		Conf.SMTPCreds.EmailAddress,
		Conf.SMTPCreds.Password,
		Conf.SMTPCreds.Hostname)
}

func (cfg *Config) GetLogger() *logrus.Logger {
	var l = logrus.New()
	l.Formatter = &logrus.JSONFormatter{}
	return l
}

func (cfg *Config) IsDebugMode() bool {
	return cfg.Mode == DebugMode
}
