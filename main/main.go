package main

import (
	"hexa/handler"
	"hexa/repository"
	"hexa/service"

	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := repository.Dbconn()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	userRepositoryDB := repository.NewUserRepositoryDB(db)

	_ = customerRepositoryMock

	userSerivice := service.NewUserService(userRepositoryDB)
	customerService := service.NewCustomerService(customerRepositoryDB)
	userHandler := handler.NewUserHandler(userSerivice)
	customerHandler := handler.NewCustomerHandler(customerService)

	//Router

	e.GET("/customers", customerHandler.GetCustomers)
	e.GET("/customers/:id", customerHandler.GetCustomer)
	e.GET("/users", userHandler.GetUser)
	//e.GET("/login", handler.Login)
	//router := mux.NewRouter()

	//router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	//http.ListenAndServe(":8000", router)
	e.Logger.Fatal(e.Start(":1323"))
}
