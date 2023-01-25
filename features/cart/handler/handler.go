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
		productID, err := strconv.Atoi(c.Param("idProduct"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		_, err = ch.srv.Add(c.Get("user"), uint(productID), ToCore(input).Qty)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("success add data"))
	}
}

func (ch *cartHandle) GetMyCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ch.srv.GetMyCart(c.Get("user"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil melihat list products dari keranjang", ToCartProductResArr(res)))
	}
}

func (ch *cartHandle) UpdateProductCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCartReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		productID, err := strconv.Atoi(c.Param("idProduct"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		err = ch.srv.UpdateProductCart(c.Get("user"), uint(productID), ToCore(input).Qty)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("update berhasil"))
	}
}

func (ch *cartHandle) DeleteProductCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID, err := strconv.Atoi(c.Param("idProduct"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		err = ch.srv.DeleteProductCart(c.Get("user"), uint(productID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Delete product berhasil"))
	}
}
