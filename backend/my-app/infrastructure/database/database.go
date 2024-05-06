package database

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection() (*gorm.DB, error) {
	fmt.Println("Init Connection Start")
	dsn := "root:password@tcp(database:3306)/mydb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB接続エラー: %v", err)
		return nil, errors.New("DB接続エラー")
	}

	fmt.Println("Database Connecting ...")
	return db, nil
}

func CloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("データベースクローズエラー: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("データベースクローズエラー: %v", err)
	}
}
