package services

import (
	"errors"
	"fmt"
	"log"
	"projects/features/cart"
	"projects/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type cartSrv struct {
	data     cart.CartData
	validasi *validator.Validate
}

func New(cd cart.CartData) cart.CartService {
	return &cartSrv{
		data:     cd,
		validasi: validator.New(),
	}
}

// Add implements cart.CartService
func (cs *cartSrv) Add(token interface{}, newCart cart.Core) (cart.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return cart.Core{}, errors.New("user not found")
	}

	err := cs.validasi.Struct(newCart)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return cart.Core{}, errors.New("validation error")
	}

	res, err := cs.data.Add(userID, newCart)
	fmt.Println(res)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Cart not found"
		} else {
			msg = "internal server error"
		}
		return cart.Core{}, errors.New(msg)
	}

	return res, nil
}
