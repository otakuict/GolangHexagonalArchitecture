package handler

import (
	//"encoding/json"
	"fmt"
	"hexa/service"
	"net/http"
	"strconv"

	//"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

// func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
// 	customers, err := h.custSrv.GetCustomers()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, err)
// 		return
// 	}

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(customers)
// }

func (h customerHandler) GetCustomers(c echo.Context) error {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {

		fmt.Fprintln(nil, err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"err": "invalid data type",
		})
	}
	return c.JSON(http.StatusOK, customers)
}

// func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
// 	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"]) //map request ด้วยkey customerIDโดยใช้mux  แล้วแปลงเป็นintegerเพราะ customer ด้านล่างต้องการint

// 	customer, err := h.custSrv.GetCustomer(customerID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, err)

// 		return
// 	}

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(customer)
// }

func (h customerHandler) GetCustomer(c echo.Context) error {
	customerID, _ := strconv.Atoi(c.Param("id")) //map request ด้วยkey customerIDโดยใช้mux  แล้วแปลงเป็นintegerเพราะ customer ด้านล่างต้องการint

	customer, err := h.custSrv.GetCustomer(customerID)
	if err != nil {

		fmt.Fprintln(nil, err)
		return err
	}

	return c.JSON(http.StatusOK, customer)
}
