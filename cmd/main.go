package main

import (
	"github/maksudkhanov/todo-app"
	"github/maksudkhanov/todo-app/pkg/handler"
	"github/maksudkhanov/todo-app/pkg/repository"
	"github/maksudkhanov/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing server: %s", err.Error())
	}
}
