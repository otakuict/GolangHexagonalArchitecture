package main

import (
	"hexa/handler"
	"hexa/repository"
	"hexa/service"

	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepositoryDB

	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)
	//Router

	e.GET("/customers", customerHandler.GetCustomers)
	e.GET("/customers/:id", customerHandler.GetCustomer)
	//router := mux.NewRouter()

	//router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	//http.ListenAndServe(":8000", router)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
