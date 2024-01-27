package main

import (
	"github.com/felipecveiga/crud_estudo_teste/handler"
	"github.com/felipecveiga/crud_estudo_teste/model"
	"github.com/felipecveiga/crud_estudo_teste/repository"
	"github.com/felipecveiga/crud_estudo_teste/service"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//CONFIGURAÇÃO DO BANCO DE DADOS
	dsn := "root:root@tcp(localhost:3306)/crud_estudo_teste?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha com a conexao com o database")
	}

	db.AutoMigrate(&model.User{})

	//INICIAR O HANDLER
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//INICIAR O ECHO
	e := echo.New()

	//ROTAS
	e.GET("users", userHandler.GetAllUsers)
	e.GET("user/:id", userHandler.GetUsers)
	e.DELETE("user/:id", userHandler.RemoveUser)

	//CONEXAO COM O SERVIDOR/ INICIAR
	e.Logger.Fatal(e.Start(":8080"))

}
