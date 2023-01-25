package services

import (
	"errors"
	"log"
	"projects/features/cart"
	"projects/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type cartUseCase struct {
	qry cart.CartData
	vld *validator.Validate
}

func New(ud cart.CartData) cart.CartService {
	return &cartUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (cuc *cartUseCase) Add(token interface{}, productId uint, qty int) (cart.Core, error) {
	err := helper.Validasi(helper.ToQtyInt(qty))
	if err != nil {
		return cart.Core{}, err
	}
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return cart.Core{}, errors.New("user tidak ditemukan")
	}

	res, err := cuc.qry.Add(userID, productId, qty)
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else if strings.Contains(err.Error(), "stock") {
			msg = "stok tidak cukup"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return cart.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cuc *cartUseCase) GetMyCart(token interface{}) ([]cart.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return []cart.Core{}, errors.New("user tidak ditemukan")
	}
	res, err := cuc.qry.GetMyCart(userID)
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else if strings.Contains(err.Error(), "stock") {
			msg = "stok tidak cukup"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []cart.Core{}, errors.New(msg)
	}

	return res, nil
}
