package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gabrielroriz/cli-fineasy/models"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var configPath = os.Getenv("HOME") + "/.fineasy"
var configFilePath = configPath + "/database.conf"

var db *DBConfig

func createConfigFolder() {
	os.Mkdir(configPath, 0777)
}

func loadDBConfig() error {

	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		//folder exists
		if _, err := os.Stat(configFilePath); !os.IsNotExist(err) {

			file, err := os.Open(configFilePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			b, err := ioutil.ReadAll(file)

			data := strings.Split(string(b), "\n")

			dbConfig := DBConfig{
				Host:     strings.Split(data[0], "=")[1],
				Port:     strings.Split(data[1], "=")[1],
				User:     strings.Split(data[2], "=")[1],
				Password: strings.Split(data[3], "=")[1],
			}

			db = &dbConfig
			//file exists
			return nil
		} else {
			//file doesn't exist
			return errors.New("doesn't have config database file yet")
		}

	} else {
		//folder doesn't exists
		createConfigFolder()
		return errors.New("doesn't have config folder yet")
	}
}

func writeDBConfig() error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createConfigFolder()
	}
	content := []byte(fmt.Sprintf("host=%s\nport=%s\nuser=%s\npassword=%s", (*db).Host, (*db).Port, (*db).User, (*db).Password))

	err := ioutil.WriteFile(configFilePath, content, 0777)

	return err
}

func SetDBConfig(dbConfig *DBConfig) {
	db = dbConfig
	writeDBConfig()
}

//InitDB : Init DB
func InitDB() (*DBConfig, error) {

	var err error

	err = loadDBConfig()

	if err != nil {
		return nil, err
	} else {

		db.DB, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=fineasy password=%s sslmode=disable",
				db.Host, db.Port, db.User, db.Password))

		if err != nil {
			return nil, err
		}

		db.DB.AutoMigrate(
			&models.Source{},
			&models.Wallet{},
			&models.Category{},
			&models.Flow{})

		// db.DB.LogMode(true)

		return db, nil
	}
}
