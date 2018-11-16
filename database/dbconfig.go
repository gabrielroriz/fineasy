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

//constants
var configPath = os.Getenv("HOME") + "/.fineasy"
var configFilePath = configPath + "/database.conf"

var dbConfig *DBConfig

type DBConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string

	DB *gorm.DB
}

func createConfigFolder() {
	os.Mkdir(configPath, 0777)
}

func SetDBConfig(new *DBConfig) {
	dbConfig = new
	writeDBConfig()
}

func writeDBConfig() error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createConfigFolder()
	}
	content := []byte(fmt.Sprintf("host=%s\nport=%s\ndatabase=%s\nuser=%s\npassword=%s", (*dbConfig).Host, (*dbConfig).Port, (*dbConfig).Database, (*dbConfig).User, (*dbConfig).Password))

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

			statements := make(map[string]string)

			for _, line := range data {

				//remove outside spaces and \t's
				line = strings.TrimSpace(strings.Replace(line, "\t", "", -1))

				if line == "" || line[0] == ';' || line[0] == '#' {
					continue
				}

				//at that moment statement[0] is key, statement[1] is value
				statement := strings.Split(line, "=")

				for i := 0; i < len(statement); i++ {
					//remove outside spaces to key and value
					statement[i] = strings.TrimSpace(statement[i])
				}

				//add to map
				statements[statement[0]] = statement[1]

			}

			config := DBConfig{
				Host:     statements["host"],
				Port:     statements["port"],
				Database: statements["database"],
				User:     statements["user"],
				Password: statements["password"],
			}

			dbConfig = &config
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
