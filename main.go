package main

import (
	"log"
	"net/http"

	"github.com/clean-architecture/connection"
	"github.com/clean-architecture/handler"
	"github.com/clean-architecture/repository"
	"github.com/clean-architecture/service"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := connection.NewPostgreConn()
	defer db.Close()

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	router := httprouter.New()
	router.GET("/employees", handler.Middleware(employeeHandler.Index))
	router.GET("/employees/:id", handler.Middleware(employeeHandler.Get))
	router.POST("/employees", handler.Middleware(employeeHandler.Create))
	router.PUT("/employees/:id", handler.Middleware(employeeHandler.Update))
	router.DELETE("/employees/:id", handler.Middleware(employeeHandler.Delete))

	log.Println("Listen at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
