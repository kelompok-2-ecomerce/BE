package handler

import (
	"net/http"
	"projects/features/cart"
	"projects/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartHandle struct {
	srv cart.CartService
}

func New(cs cart.CartService) cart.CartHandler {
	return &cartHandle{
		srv: cs,
	}
}

func (ch *cartHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCartReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		ProductID, err := strconv.Atoi(c.Param("idProduct"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		_, err = ch.srv.Add(c.Get("user"), uint(ProductID), ToCore(input).Qty)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("success add data"))
	}
}
