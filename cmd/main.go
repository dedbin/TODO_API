package main

import (
	handler "TODO_API/pkg/handler"
	repository "TODO_API/pkg/repository"
	service "TODO_API/pkg/service"
	todo "TODO_API/smth"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("ошибка сервера: %v", err.Error())
	}
}
