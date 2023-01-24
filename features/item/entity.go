package item

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Nama_Barang string `validate:"required"`
	Image_url   string
	Nama        string
	Deskripsi   string `validate:"required"`
	Harga       float64
	Stok        int
	CreatedAt   time.Time
}

type ItemHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAllPost() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyItem() echo.HandlerFunc
	GetID() echo.HandlerFunc
}

type ItemService interface {
	Add(token interface{}, newItem Core) (Core, error)
	Update(token interface{}, itemID int, updatedData Core) (Core, error)
	GetAllPost() ([]Core, error)
	Delete(token interface{}, itemID int) error
	MyItem(token interface{}) ([]Core, error)
}

type ItemData interface {
	Add(userID int, newItem Core) (Core, error)
	Update(userID int, itemID int, updatedData Core) (Core, error)
	GetAllPost() ([]Core, error)
	Delete(userID int, itemID int) error
	MyItem(userID int) ([]Core, error)
}
