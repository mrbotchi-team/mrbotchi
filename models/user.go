package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrbotchi-team/mrbotchi/config"

	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	IUserModel interface {
		FindByName(name string) *User
		Insert() error
	}
	User struct {
		Name      string    `db:"name"`
		Password  string    `db:"password"`
		IsDeleted bool      `db:"is_deleted"`
		CreatedAt time.Time `db:"created_at"`
	}
	UserModel struct {
		db *sqlx.DB
	}
)

func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{db}
}

func (um *UserModel) FindByName(name string) *User {
	user := &User{}
	um.db.Get(user, "SELECT * FROM users WHERE name = $1 LIMIT 1", name)

	return user
}

func (um *UserModel) Insert(name, password string, param config.Argon2Config) error {
	hashedPassword, err := utils.GenerateHashedPassword(password, param)
	if nil != err {
		return err
	}

	um.db.MustExec("INSERT INTO users (name, password, is_deleted, created_at) VALUES ($1, $2, $3, $4)", name, hashedPassword, false, time.Now())

	return nil
}
