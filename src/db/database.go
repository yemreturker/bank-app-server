package db

import (
	"bank-app-server/src/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bank.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	// Migration (tabloları oluştur)
	DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
}