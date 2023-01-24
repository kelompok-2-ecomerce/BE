package handler

import (
	"fmt"
	"log"
	"net/http"
	"projects/features/cart"
	"projects/helper"

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

// Add implements cart.CartHandler
func (ch *cartHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdateCartRequest{}
		// idParam := c.Param("posting_id")
		// id, _ := strconv.Atoi(idParam)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := ch.srv.Add(c.Get("user"), *cnv)
		fmt.Println(res)
		if err != nil {
			fmt.Println(err)
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		cart := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan barang ke cart", cart))
	}
}
