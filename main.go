package main

import (
	"log"
	"projects/config"

	id "projects/features/item/data"
	ihl "projects/features/item/handler"
	isrv "projects/features/item/services"
	"projects/features/user/data"
	"projects/features/user/handler"
	"projects/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	itemData := id.New(db)
	itemsrv := isrv.New(itemData)
	itemHdl := ihl.New(itemsrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.AllUser())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/products", itemHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/products/:id", itemHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/myproducts", itemHdl.MyItem(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
