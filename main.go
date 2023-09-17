package main

import (
	"log"
	"net/http"

	"github.com/bluewd111/go-template/app/user/handler"
	"github.com/bluewd111/go-template/app/user/repository"
	"github.com/bluewd111/go-template/app/user/service"
)

func main() {
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	uHandler := handler.GetUserHandlerInstance(userService)
	http.Handle(uHandler.GetPathAndHandler())

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}
