package db

import (
	"github.com/AlanKabolov/prpr-v1/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error{
	return db.AutoMigrate(&models.Event{})
}