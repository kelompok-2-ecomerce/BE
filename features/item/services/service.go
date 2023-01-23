package services

import (
	"errors"
	"fmt"
	"projects/features/item"
	"projects/helper"
	"strings"

	"github.com/go-playground/validator/v10"
)

type itemSrv struct {
	data     item.ItemData
	validasi *validator.Validate
}

func New(pd item.ItemData) item.ItemService {
	return &itemSrv{
		data:     pd,
		validasi: validator.New(),
	}
}

// Add implements item.ItemService
func (ps *itemSrv) Add(token interface{}, newItem item.Core) (item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return item.Core{}, errors.New("user not found")
	}

	res, err := ps.data.Add(userID, newItem)
	fmt.Println(res)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Posting not found"
		} else {
			msg = "internal server error"
		}
		return item.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements item.ItemService
func (*itemSrv) Delete(token interface{}, itemID int) error {
	panic("unimplemented")
}

// GetAllPost implements item.ItemService
func (*itemSrv) GetAllPost() ([]item.Core, error) {
	panic("unimplemented")
}

// MyPost implements item.ItemService
func (*itemSrv) MyPost(token interface{}) ([]item.Core, error) {
	panic("unimplemented")
}

// Update implements item.ItemService
func (ps *itemSrv) Update(token interface{}, itemID int, updatedData item.Core) (item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return item.Core{}, errors.New("id user not found")
	}
	if validasieror := ps.validasi.Struct(updatedData); validasieror != nil {
		return item.Core{}, nil
	}

	res, err := ps.data.Update(userID, itemID, updatedData)
	if err != nil {
		fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "item not found"
		} else {
			msg = "internal server error"
		}
		return item.Core{}, errors.New(msg)
	}

	return res, nil
}
