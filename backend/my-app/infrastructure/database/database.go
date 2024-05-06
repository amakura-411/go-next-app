package database

import (
	"errors"
	"fmt"
	"log"
	"my-app/domain/entity"
	"time"

	"github.com/google/uuid"
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

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // テーブル名のプレフィックスがある場合に指定
			SingularTable: true, // テーブル名を単数形にする
		},
	})

	if err != nil {
		log.Printf("DB接続エラー: %v", err)
		return nil, errors.New("DB接続エラー")
	}

	// userテーブルに値追加
	DB.Create(&entity.User{
		User_id:    uuid.New().String(),
		Username:   "Test",
		Password:   "Test",
		Icon:       "default",
		Created_at: time.Now(),
		Updated_at: time.Now(), // UpdateAtフィールドの値を追加
	})

	// userテーブルに値を確認
	var users []*entity.User
	if err := DB.Find(&users).Error; err != nil {
		fmt.Println("errorです")
		return nil, err
	}

	fmt.Println(DB, "Database Connected")
	fmt.Println("Database Connecting ...")
	return DB, nil
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
