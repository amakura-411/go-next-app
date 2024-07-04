package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func InitConnection() (*gorm.DB, error) {
	fmt.Println("Init Connection Start")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// ひとまず環境変数に設定しているが、後で.envファイルに移行する
		"root",
		"password",
		"database",
		"3306",
		"mydb",
	)
	// 接続するまで、何度もmysqlに接続する（5秒間隔）
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",   // テーブル名のプレフィックスがある場合に指定
				SingularTable: true, // テーブル名を単数形にする
			},
		})
		if err == nil {
			break
		}
		log.Printf("DB接続エラー: %v", err)
		// ５秒待つ
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Printf("DB接続エラー: %v", err)

		return nil, errors.New("DB接続エラー")
	}

	fmt.Println("Database Connecting ...")
	return DB, nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitConnection()
	}

	return DB
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
