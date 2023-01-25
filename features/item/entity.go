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
	MyProducts() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetProductByID() echo.HandlerFunc
}

type ItemService interface {
	Add(token interface{}, newItem Core, image *multipart.FileHeader) (Core, error)
	Update(token interface{}, itemID int, updatedData Core, image *multipart.FileHeader) (Core, error)
	GetAllProducts() ([]Core, error)
	MyProducts(token interface{}) ([]Core, error)
	GetProductByID(token interface{}, productID int) (Core, error)
	Delete(token interface{}, itemID int) error
}

type ItemData interface {
	Add(userID int, newItem Core) (Core, error)
	Update(userID int, itemID int, updatedData Core) (Core, error)
	GetAllProducts() ([]Core, error)
	MyProducts(userID int) ([]Core, error)
	GetProductByID(userID int, productID int) (Core, error)
	Delete(userID int, itemID int) error
}
