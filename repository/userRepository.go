package repository

import (
	"gorm.io/gorm"
	"jemmy-sapta/model"
	"jemmy-sapta/protos/pbuild"
)

type UserRepository interface {
	CreateUser(user pbuild.UserProto) (*pbuild.UserProto, error)
	GetUserBy(username string, password string) (*pbuild.UserProto, error)
	GetUserById(id int64) (*pbuild.UserProto, error)
}

type ResourceUser struct {
	Db *gorm.DB
}

func InitUserRepository(
	db *gorm.DB,
) *ResourceUser {
	return &ResourceUser{
		Db: db,
	}
}

func (r *ResourceUser) CreateUser(user pbuild.UserProto) (*pbuild.UserProto, error) {
	var userData model.User
	if errorInsert := r.Db.Raw(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password",
		user.Username,
		user.Password,
	).First(&userData); errorInsert.Error != nil {
		return nil, errorInsert.Error
	}

	user.Id = userData.Id
	return &user, nil
}

func (r *ResourceUser) GetUserBy(username string, password string) (*pbuild.UserProto, error) {
	var user pbuild.UserProto
	if searchUser := r.Db.Raw(
		"SELECT * FROM users WHERE username=$1 AND password=$2",
		username,
		password,
	).First(&user); searchUser.Error != nil {
		return nil, searchUser.Error
	}

	return &user, nil
}

func (r *ResourceUser) GetUserById(id int64) (*pbuild.UserProto, error) {
	var user pbuild.UserProto
	if searchUser := r.Db.Raw(
		"SELECT * FROM users WHERE id=$1",
		id,
	).First(&user); searchUser.Error != nil {
		return nil, searchUser.Error
	}

	return &user, nil
}
