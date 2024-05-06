package user

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// user_idは、同一なものは存在しないようにする
	// その同一性の保持はUser_id.goで担保する
	user_id  string
	username string
	password string
	icon     string
	// time.Timeは、Goの標準ライブラリに含まれる型で、日時を表すための型
	created_at time.Time
	update_at  time.Time
}

func NewUser(username, password, icon string, create_at time.Time) (User, error) {
	user := &User{}
	user.user_id = uuid.New().String()

	// usernameのバリデーション
	err := user.changeUsername(username)
	if err != nil {
		return *user, err
	}

	// passwordのバリデーション
	err = user.changePassword(password)
	if err != nil {
		return *user, err
	}
	// iconが空文字の場合、デフォルトのアイコンを代入
	if icon == "" {
		user.icon = "default"
	} else {
		user.icon = icon
	}

	// created_atに値がない場合、現在時刻を代入
	if create_at.IsZero() {
		user.created_at = time.Now()
	} else {
		user.created_at = create_at
	}
	user.update_at = user.created_at
	return *user, nil
}

func (user *User) changeUsername(username string) (err error) {
	// あとでバリデーションで処理するが、ひとまずここで実装
	if username == "" {
		err = errors.New("username is required")
	} else if utf8.RuneCountInString(username) < 3 {
		err = errors.New("username is too short")
	}
	// errorがnilでない場合、エラーが発生しているので、そのエラーを返す
	if err != nil {
		return err
	}
	user.username = username
	return nil
}

func (user *User) changePassword(password string) error {
	// passwordのバリデーションも同様に実装するが、ひとまずここで実装する
	if password == "" {
		return errors.New("password is required")
	} else if utf8.RuneCountInString(password) < 8 {
		return errors.New("password is too short")
	}
	// パスワードをハッシュ化して保存
	hashedPassword, err := user.createHashedPassword(password)
	if err != nil {
		// エラーログを出力
		fmt.Println("failed to create hashed password: %w", err)
		return err
	}
	user.password = hashedPassword
	return nil
}

func (user *User) createHashedPassword(password string) (string, error) {
	// 与えられたパスワードをハッシュ化して返す
	// bcrypt.GenerateFromPassword()は、引数に与えられたパスワードをハッシュ化して返す
	// この関数は、ハッシュ化されたパスワードとエラーを返すため、hashed_passwordと_で受け取る
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// エラーログを出力
		fmt.Println("failed to generate hashed password:", err)
		return "", err
	}

	return string(hashBytes), nil
}
