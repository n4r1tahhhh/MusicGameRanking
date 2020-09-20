package handler

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() (*gorm.DB, error) {
	// 環境変数の設定
	EnvLoad()
	user := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")
	endpoint := os.Getenv("ENDPOINT")
	dbName := os.Getenv("DATABASE_NAME")

	// DBと接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, endpoint, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
