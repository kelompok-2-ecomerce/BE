package cart

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Nama_Barang string
	ItemID      uint
	Qty         string
	Nama        string
	CreatedAt   time.Time
}

type CartHandler interface {
	Add() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// GetAllItems() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// MyItem() echo.HandlerFunc
	// GetID() echo.HandlerFunc
}

type CartService interface {
	Add(token interface{}, newCart Core) (Core, error)
	// Update(token interface{}, itemID int, updatedData Core) (Core, error)
	// GetAllItems() ([]Core, error)
	// GetID(ItemID int) (Core, error)
	// Delete(token interface{}, itemID int) error
	// MyItem(token interface{}) ([]Core, error)
}

type CartData interface {
	Add(userID int, newCart Core) (Core, error)
	// Update(userID int, itemID int, updatedData Core) (Core, error)
	// GetAllItems() ([]Core, error)
	// GetID(ItemID int) (Core, error)
	// Delete(userID int, itemID int) error
	// MyItem(userID int) ([]Core, error)
}
