package main

import (
	"fmt"
	"github/sumitpant/authService/cmd/api/repository"
	"github/sumitpant/authService/cmd/api/service"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/task_mgmt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to db")
	}
	fmt.Println("db",db)
	
	repo:= repository.NewConn(db)
	
	ser := service.InjectRepo(repo);


	// router instance
	router := Router(ser);


	// Custom server struct
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} 
		
	


}
