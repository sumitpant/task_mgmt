package main

import (
	"net/http"
	"github/sumitpant/authService/cmd/api/service"
	"github/sumitpant/authService/cmd/api/middleware"
	"github.com/gorilla/mux"
)

// create Router and add  common headers headers
func Router(s *service.Service) *mux.Router {
	r := mux.NewRouter()
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	r.Use(middleware.AddHeaders(headers))

	r.HandleFunc("/sign-up",s.CreateUser).Methods("POST");
	r.HandleFunc("/login",s.Login).Methods("POST");
	




	return r

}
