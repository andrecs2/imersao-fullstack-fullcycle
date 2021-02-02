package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Base      `valid:"required"`
	Name	  string    `gorm:"column:name;type:varchar(255);not null" valid:"notnull"`
	ID        string    `gorm:"column:user_id;type:uuid;not null" valid:"-"`
	Email     string    `gorm:"column:user_email;type:varchar(255);not null" valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*Account, error) {
	user := User{
		Name:      name,
		Email:    email,
	}

	user.ID = uuid.NewV4().String()

	err := user.isValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
