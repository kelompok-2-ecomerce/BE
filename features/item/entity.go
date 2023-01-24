package item

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Nama_Barang string
	Image_url   string
	NamaUser    string
	Alamat      string
	Deskripsi   string
	Harga       float64
	Stok        int
}

type ItemHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// MyPost() echo.HandlerFunc
	// GetID() echo.HandlerFunc
}

type ItemService interface {
	Add(token interface{}, newItem Core, image *multipart.FileHeader) (Core, error)
	Update(token interface{}, itemID int, updatedData Core, image *multipart.FileHeader) (Core, error)
	GetAllProducts() ([]Core, error)
	// Delete(token interface{}, itemID int) error
	// MyPost(token interface{}) ([]Core, error)
}

type ItemData interface {
	Add(userID int, newItem Core) (Core, error)
	Update(userID int, itemID int, updatedData Core) (Core, error)
	GetAllProducts() ([]Core, error)
	// Delete(userID int, itemID int) error
	// MyPost(userID int) ([]Core, error)
}
