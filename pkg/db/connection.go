package db

import (
	"fmt"
	"log"

	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/config"
	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Initdb(cfg config.Config) Handler {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Db_Host, cfg.Db_User, cfg.Db_Password, cfg.Db_Name, cfg.Db_Port)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	db.AutoMigrate(&models.User{})

	return Handler{db}
}
