package handler

import (
	"net/http"

	"github.com/felipecveiga/crud_estudo_teste/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *service.UserService
	
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
		
	}

}


func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {"error": "Failed to fetch users"})

	}
	return c.JSON(http.StatusOK, users)
}