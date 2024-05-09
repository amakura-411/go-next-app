package user

import (
	"errors"
	"fmt"
	userEntity "my-app/domain/entity/user"

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
func (ur *UserRepository) GetUserByID(id string) (*userEntity.User, error) {
	var user *userEntity.User
	if err := ur.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		fmt.Println("user not found")
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetAllUsersは全てのユーザーを取得する
func (ur *UserRepository) GetAllUsers() ([]*userEntity.User, error) {
	var users []*userEntity.User
	if err := ur.DB.Find(&users).Error; err != nil {
		fmt.Println("errorです")
		return nil, err
	}
	return users, nil
}

// CreateUserは新しいユーザーを作成する
func (ur *UserRepository) CreateUser(u userEntity.User) error {
	if err := ur.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUserはユーザー情報を更新する
func (ur *UserRepository) UpdateUser(id string, u *userEntity.User) error {
	var user userEntity.User
	fmt.Println("UpdateUserId:", id)
	if err := ur.DB.First(id).Error; err != nil {
		// userIdが存在しない場合、エラーを返す
		fmt.Println(("IDが存在しません"))
		return err
	}
	// userデータを更新する
	if err := ur.DB.Model(&user).Updates(u).Error; err != nil {
		return err
	}
	return nil

}

// DeleteUserByIDは指定されたIDのユーザーを削除する
func (ur *UserRepository) DeleteUserByID(id string) error {
	if err := ur.DB.Delete(&userEntity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserRepository) CheckUserID(id string) bool {
	if err := us.DB.First(id).Error; err != nil {
		return true
	}
	return false
}
