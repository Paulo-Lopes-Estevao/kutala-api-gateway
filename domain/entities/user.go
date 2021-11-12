package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	State    bool      `json:"state"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func NewUser(name string, username string, password string) (*User, error) {
	User := &User{
		Name:     name,
		Username: username,
		Password: password,
		State:    true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	err := User.passwordEncrypt()

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (User *User) passwordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	User.Password = string(password)

	err = User.validate()

	if err != nil {
		return err
	}

	return nil

}

func (User *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password))
	return err == nil
}

func (User *User) validate() error {

	_, err := govalidator.ValidateStruct(User)

	if err != nil {
		return err
	}

	return nil
}
