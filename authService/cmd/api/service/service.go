package service

import (
	"encoding/json"
	"fmt"
	"github/sumitpant/authService/cmd/api/entities"
	"github/sumitpant/authService/cmd/api/middleware"
	"github/sumitpant/authService/cmd/api/modals"
	"github/sumitpant/authService/cmd/api/repository"

	"net/http"
)

type Service struct {
	repo *repository.Repo
}
type AuthService interface {
	CreateUser(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

func InjectRepo(repo *repository.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &modals.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Sprintln(w, "Internal Error")
	}
	fmt.Println("user", user)
	encrypted_pass, err := middleware.Encrypt(user.Password)
	if err != nil {
		fmt.Fprintln(w, "Encrptyion Error")
	}

	userEntity := entities.Auth{}
	userEntity.Email = user.Email
	userEntity.Password = encrypted_pass
	err, created := s.repo.CreateUser(&userEntity)
	if err != nil {
		fmt.Fprintln(w, "Error in user creation")
	}
	fmt.Fprintln(w, created)

}

func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	user := &modals.User{}
	json.NewDecoder(r.Body).Decode(user)
	if user.Email == "" {
		fmt.Fprintln(w, "email should not be empty")
	}
	result, _ := s.repo.Login(user.Email)
	fmt.Println(result)
	response, _ := json.MarshalIndent(result, "", "\t")
	fmt.Fprintln(w, string(response))
}
