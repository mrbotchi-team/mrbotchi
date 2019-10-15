package models

import (
	"time"

	"github.com/mrbotchi-team/mrbotchi/config"

	"github.com/mrbotchi-team/mrbotchi/utils"
)

type User struct {
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt time.Time `db:"created_at"`
}

func NewUser(name, password string, param config.Argon2Config) (*User, error) {
	hashedPassword, err := utils.GenerateHashedPassword(password, param)
	if nil != err {
		return nil, err
	}

	return &User{
		Name:      name,
		Password:  hashedPassword,
		IsDeleted: false,
		CreatedAt: time.Now(),
	}, nil
}
