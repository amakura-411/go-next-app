package database

import (
	"fmt"
	user "my-app/domain/entity"

	"gorm.io/gorm"
)

// UserRepositoryはユーザーデータの永続化を担当する
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepositoryはユーザーリポジトリを初期化する
func NewUserRepository(db *gorm.DB) *UserRepository {
	fmt.Println("NewUserRepository")
	return &UserRepository{DB: db}
}

// GetUserByIDは指定されたIDのユーザーを取得する
func (ur *UserRepository) GetUserByID(id int) (*user.User, error) {
	var u user.User
	if err := ur.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// GetAllUsersは全てのユーザーを取得する
func (ur *UserRepository) GetAllUsers() ([]*user.User, error) {
	var users []*user.User

	// dbの接続確認
	fmt.Println(ur.DB, "ur.DB")

	// username Test　password Test
	// ur.DB.Create(&user.User{
	// 	User_id:   "1",
	// 	UserName:  "Test",
	// 	Password:  "Test",
	// 	Icon:      "default",
	// 	CreatedAt: time.Now(),
	// 	UpdateAt:  time.Now(),
	// })

	if err := ur.DB.Find(&users).Error; err != nil {
		fmt.Println("errorです")
		return nil, err
	}
	return users, nil
}

// CreateUserは新しいユーザーを作成する
func (ur *UserRepository) CreateUser(u *user.User) error {
	if err := ur.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUserはユーザー情報を更新する
func (ur *UserRepository) UpdateUser(u *user.User) error {
	if err := ur.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserByIDは指定されたIDのユーザーを削除する
func (ur *UserRepository) DeleteUserByID(id int) error {
	if err := ur.DB.Delete(&user.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// 他のメソッド...
