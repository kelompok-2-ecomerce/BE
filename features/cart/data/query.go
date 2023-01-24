package data

import (
	"errors"
	"log"
	"projects/features/cart"
	"strings"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartData{
		db: db,
	}
}

// Add implements cart.CartData
func (cd *cartData) Add(userID int, newCart cart.Core) (cart.Core, error) {
	cnv := CoreToData(newCart)
	cnv.UserID = uint(userID)

	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("add cart query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return cart.Core{}, errors.New(msg)
	}

	newCart.ID = cnv.ID

	return newCart, nil
}
