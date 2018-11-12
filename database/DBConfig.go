package database

import "github.com/jinzhu/gorm"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string

	DB *gorm.DB
}
