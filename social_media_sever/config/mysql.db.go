package config

import (
	"database/sql"
	"fmt"
	"log"
	"social_media_sever/models"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	LoadEnv()

	// First connect without database name to create the database
	rootDSN := GetRootDSN()

	sqlDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		log.Fatal("Failed to connect to MySQL server:", err)
	}
	defer sqlDB.Close()

	// Create the database if it doesn't exist
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", Config.DBName))
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	// Now connect to the database
	dsn := GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
	fmt.Println("Database connected and migrated")
}
