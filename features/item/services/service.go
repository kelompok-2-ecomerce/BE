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
			msg = "Items not found"
		} else {
			msg = "internal server error"
		}
		return item.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements item.ItemService
func (ps *itemSrv) Delete(token interface{}, itemID int) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := ps.data.Delete(userID, itemID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "item not found"
		} else {
			msg = "internal server error"

		}
		return errors.New(msg)
	}
	return nil
}

// GetAllPost implements item.ItemService
func (ps *itemSrv) GetAllItems() ([]item.Core, error) {
	All, err := ps.data.GetAllItems()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Products not found"
		} else {
			msg = "internal server error"
		}
		return nil, errors.New(msg)
	}

	return All, nil
}

// MyPost implements item.ItemService
func (ps *itemSrv) MyItem(token interface{}) ([]item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return nil, errors.New("user not found")
	}

	res, _ := ps.data.MyItem(userID)

	return res, nil
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

// GetID implements item.ItemService
func (ps *itemSrv) GetID(ItemID int) (item.Core, error) {
	data, err := ps.data.GetID(ItemID)

	if err != nil {
		fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "ID Product not found"
		} else {
			msg = "internal server error"
		}
		return item.Core{}, errors.New(msg)
	}
	return data, nil

}
