package main

import (
	"github/maksudkhanov/todo-app"
	"github/maksudkhanov/todo-app/pkg/handler"
	"log"
)
 
func main() {
	handler:=new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("3000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error while runing server: %s", err.Error())
	}
}
