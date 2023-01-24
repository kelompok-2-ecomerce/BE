package handler

import (
	"log"
	"net/http"
	"projects/features/item"
	"projects/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type itemHandle struct {
	srv item.ItemService
}

func New(ps item.ItemService) item.ItemHandler {
	return &itemHandle{
		srv: ps,
	}
}

// Add implements item.ItemHandler
func (ph *itemHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdatePostingRequest{}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := input.reqToCore()

		_, err = ph.srv.Add(c.Get("user"), cnv, file)
		// res, err := ph.srv.Add(c.Get("user"), cnv, file)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		// item := ToResponse("add", res)

		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("uccess add data"))
		// return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("sukses menambahkan barang", item))
	}
}

// Delete implements item.ItemHandler
func (*itemHandle) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAllPost implements item.ItemHandler
func (ih *itemHandle) GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ih.srv.GetAllProducts()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("berhasil melihat list products", ListPostCoreToPostsRespon(res)))
	}
}

// GetID implements item.ItemHandler
func (*itemHandle) GetID() echo.HandlerFunc {
	panic("unimplemented")
}

// MyPost implements item.ItemHandler
func (ih *itemHandle) MyProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ih.srv.MyProducts(c.Get("user"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("berhasil melihat list products", ListPostCoreToPostsRespon(res)))
	}
}

// Update implements item.ItemHandler
func (ph *itemHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdatePostingRequest{}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)

		ItemID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		_, err = ph.srv.Update(c.Get("user"), ItemID, *cnv, file)
		// res, err := ph.srv.Update(c.Get("user"), ItemID, *cnv, file)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		// item := ToResponse("update", res)

		// return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "update berhasil", item))
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("update berhasil"))
	}

}
