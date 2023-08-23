package main

import (
	"log"

	"github.com/dedbin/TODO_API/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv = new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("ошибка сервера: %v", err.Error())
	}
}
