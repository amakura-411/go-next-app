package database

import (
	user "my-app/domain/entity"

	"gorm.io/gorm"
)

// UserRepositoryはユーザーデータの永続化を担当する
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepositoryはユーザーリポジトリを初期化する
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserByIDは指定されたIDのユーザーを取得する
func (ur *UserRepository) GetUserByID(id string) (*user.User, error) {
	var user user.User

	if err := ur.DB.Find(&user, "User_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsersは全てのユーザーを取得する
func (ur *UserRepository) GetAllUsers() ([]*user.User, error) {
	var users []*user.User

	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUserは新しいユーザーを作成する
func (ur *UserRepository) CreateUser(u user.User) error {
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
