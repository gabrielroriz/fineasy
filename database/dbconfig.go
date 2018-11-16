package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string

	DB *gorm.DB
}

var configPath = os.Getenv("HOME") + "/.fineasy"
var configFilePath = configPath + "/database.conf"

var db *DBConfig

func createConfigFolder() {
	os.Mkdir(configPath, 0777)
}

func SetDBConfig(dbConfig *DBConfig) {
	db = dbConfig
	writeDBConfig()
}

func writeDBConfig() error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createConfigFolder()
	}
	content := []byte(fmt.Sprintf("host=%s\nport=%s\nuser=%s\npassword=%s", (*db).Host, (*db).Port, (*db).User, (*db).Password))

	err := ioutil.WriteFile(configFilePath, content, 0777)

	return err
}

func readDBConfig() error {

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
