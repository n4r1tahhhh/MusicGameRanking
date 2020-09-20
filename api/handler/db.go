package handler

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// ConnectDB データベースに接続
func ConnectDB(isTest bool) (*gorm.DB, error) {
	// 開発環境かテスト環境かで開くenvファイルを選ぶ
	environment := ""
	if isTest {
		environment = "test"
	} else {
		environment = "development"
	}
	err := godotenv.Load(fmt.Sprintf(".env.%s", environment))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数をfetch
	user := os.Getenv("USER_NAME")
	password := os.Getenv("USER_PASSWORD")
	endpoint := os.Getenv("ENDPOINT")
	dbName := os.Getenv("DATABASE_NAME")

	// DBと接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, endpoint, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
