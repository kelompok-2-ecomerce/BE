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

	cd "projects/features/cart/data"
	chl "projects/features/cart/handler"
	csrv "projects/features/cart/services"

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

	cartData := cd.New(db)
	cartsrv := csrv.New(cartData)
	cartHdl := chl.New(cartsrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/products", itemHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products", itemHdl.GetAllProducts())
	e.GET("/myproducts", itemHdl.MyProducts(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/products/:id", itemHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products/:idProduct", itemHdl.GetProductByID(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/products/:idProduct", itemHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/carts/:idProduct", cartHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/carts", cartHdl.GetMyCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/carts/:idProduct", cartHdl.UpdateProductCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/carts/:idProduct", cartHdl.DeleteProductCart(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
