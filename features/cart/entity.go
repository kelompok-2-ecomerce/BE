package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	ProductName string
	ImageUrl    string
	Price       float64
	Qty         int
	ItemID      uint
}

type CartHandler interface {
	Add() echo.HandlerFunc
}

type CartService interface {
	Add(token interface{}, productId uint, qty int) (Core, error)
}

type CartData interface {
	Add(userID int, productId uint, qty int) (Core, error)
}
