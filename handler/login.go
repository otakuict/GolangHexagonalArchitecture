package handler

import (
	//"database/sql"
	//"encoding/json"
	"fmt"
	"hexa/service"
	"net/http"

	"github.com/labstack/echo/v4"
	//"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUser(c echo.Context) error {
	userTest := "johndoe" //map request ด้วยkey customerIDโดยใช้mux  แล้วแปลงเป็นintegerเพราะ customer ด้านล่างต้องการint

	user, err := h.userSrv.GetUser(userTest)
	if err != nil {

		fmt.Fprintln(nil, err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
