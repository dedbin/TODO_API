package main

import (
	"TODO_API/pkg/handler"
	todo "TODO_API/smth"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("ошибка сервера: %v", err.Error())
	}
}
