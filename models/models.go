package models

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/oschwald/geoip2-golang"
	"github.com/svarlamov/bintrad/config"
	"io"
	"os"
	"strings"
)

var Log = config.Conf.GetLogger()
var db *gorm.DB
var geodb *geoip2.Reader

func Setup() error {
	var err error = nil
	// MySQL
	db, err = gorm.Open("mysql", config.Conf.DBPath)
	// SQLite: db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.SingularTable(true)

	geodb, err = geoip2.Open(config.Conf.GeoLiteMMDBPath)
	if err != nil {
		return err
	}

	return nil
}

func WipeDatabase() error {
	tableCnt := DBBasicCountModel{}
	err := db.Raw("SELECT COUNT(*) as count FROM information_schema.tables WHERE table_schema = '" + config.Conf.DBName + "';").Scan(&tableCnt).Error
	if err != nil {
		return err
	}
	if tableCnt.Count == 0 {
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	rollback := func() {
		tx.Rollback()
	}
	err = tx.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("SET @tables = NULL;").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("SELECT GROUP_CONCAT(table_schema, '.', table_name) INTO @tables FROM information_schema.tables WHERE table_schema = '" + config.Conf.DBName + "';").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("SET @tables = CONCAT('DROP TABLE ', @tables);").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("PREPARE stmt FROM @tables;").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("EXECUTE stmt;").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("DEALLOCATE PREPARE stmt;").Error
	if err != nil {
		rollback()
		return err
	}
	err = tx.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error
	if err != nil {
		rollback()
		return err
	}
	return tx.Commit().Error
}

func InitializeDatabaseFromFile(path string) error {
	buf := bytes.NewBuffer(nil)
	fileHandler, err := os.Open(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(buf, fileHandler)
	if err != nil {
		return err
	}

	sqlString := string(buf.Bytes())
	sqlStrings := strings.Split(sqlString, ";")
	for _, text := range sqlStrings {
		if text == "" || text == "\n" {
			continue
		}
		err = db.Exec(text + ";").Error
		if err != nil {
			return err
		}
	}
	return nil
}
