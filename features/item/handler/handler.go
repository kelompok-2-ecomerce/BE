package handler

import (
	"errors"
	"fmt"
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
		file, errPath := c.FormFile("image")

		fmt.Print("error get path handler, err = ", errPath)

		if file != nil {
			res, err := helper.UploadImage(c)
			// fmt.Println(res)
			if err != nil {
				fmt.Println(err)
				return errors.New("create gambar failed cannot upload data")
			}
			input.Image_url = res
			// fmt.Println(input.Image_url)
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := input.reqToCore()

		res, err := ph.srv.Add(c.Get("user"), cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		item := ToResponse("add", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan barang", item))
	}
}

// Delete implements item.ItemHandler
func (*itemHandle) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAllPost implements item.ItemHandler
func (*itemHandle) GetAllPost() echo.HandlerFunc {
	panic("unimplemented")
}

// GetID implements item.ItemHandler
func (*itemHandle) GetID() echo.HandlerFunc {
	panic("unimplemented")
}

// MyPost implements item.ItemHandler
func (*itemHandle) MyPost() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements item.ItemHandler
func (ph *itemHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddUpdatePostingRequest{}
		file, errPath := c.FormFile("file")

		fmt.Print("error get path handler, err = ", errPath)

		if file != nil {
			res, err := helper.UploadImage(c)
			// fmt.Println(res)
			if err != nil {
				fmt.Println(err)
				return errors.New("create gambar failed cannot upload data")
			}
			input.Image_url = res
			// fmt.Println(input.Image_url)
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		cnv := ToCore(input)

		ItemID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		res, err := ph.srv.Update(c.Get("user"), ItemID, *cnv)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		item := ToResponse("update", res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses mengubah barang", item))
	}

}
