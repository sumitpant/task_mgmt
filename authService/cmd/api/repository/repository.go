package repository

import (
	"fmt"
	"github/sumitpant/authService/cmd/api/entities"

	"gorm.io/gorm"
)

type Repo struct {
	conn *gorm.DB
}

type AuthRepo interface {
	CreateUser(*entities.Auth) (error, bool)
	Login(string) (error, *entities.Auth)
}

func NewConn(conn *gorm.DB) *Repo {
	return &Repo{conn: conn}
}

func (repo *Repo) CreateUser(user *entities.Auth) (error, bool) {
	//repo.conn.Create();
	result := repo.conn.Table("auth").Create(user)
	return result.Error, true
}

func (repo *Repo) Login(email string) (*entities.Auth, error) {
	dest := entities.Auth{}

	result:= repo.conn.Table("auth").First(&dest, "email = ?",email);
	fmt.Println(result);
	fmt.Println(dest.Email);

	return &dest, nil
}
