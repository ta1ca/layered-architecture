package config

import (
	"awesomeProject/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {
	// 実行環境取得
	env := os.Getenv("HAKEN_ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// 環境変数取得
	godotenv.Load(".env." + env)
	godotenv.Load()
	fmt.Println(os.Getenv("CONNECT"))

	// DB接続
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}

	autoMigration()

	return db
}

// DB終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// Item マイグレーション
func autoMigration() {
	db.AutoMigrate(&model.Item{})
}
