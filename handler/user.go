package handler

import (
	"net/http"

	"github.com/felipecveiga/crud_estudo_teste/model"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Falha ao buscar usuários"})

	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUsers(c echo.Context) error {

	userID := c.Param("id")
	user, err := h.UserService.GetUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Usuário não encontrado"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveUser(c echo.Context) error {

	userID := c.Param("id")
	user, err := h.UserService.RemoveUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "usuário não encontrado"})
	}
	return c.JSON(http.StatusOK, user)

}

func (h *UserHandler) CreateUser(c echo.Context) error {

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "dados invalidos"})
	}
	err := h.UserService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UserUpdate(c echo.Context) error {

	userID := c.Param("id")
	var userData model.User
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "erro ao ler os dados da requisição"})
	}

	userUpdate, err := h.UserService.UserUpdate(userID, userData)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuário não encontrado"})
	}
	return c.JSON(http.StatusOK, userUpdate)
}
