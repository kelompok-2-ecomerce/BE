package data

import (
	"errors"
	"log"
	"projects/features/item"
	"strings"

	"gorm.io/gorm"
)

type itemData struct {
	db *gorm.DB
}

func New(db *gorm.DB) item.ItemData {
	return &itemData{
		db: db,
	}
}

// Add implements item.ItemData
func (pd *itemData) Add(userID int, newItem item.Core) (item.Core, error) {
	cnv := CoreToData(newItem)
	cnv.UserID = uint(userID)

	err := pd.db.Create(&cnv).Error
	if err != nil {
		log.Println("add post query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return item.Core{}, errors.New(msg)
	}

	newItem.ID = cnv.ID

	return newItem, nil
}

// Delete implements item.ItemData
func (*itemData) Delete(userID int, itemID int) error {
	panic("unimplemented")
}

// GetAllPost implements item.ItemData
func (*itemData) GetAllPost() ([]item.Core, error) {
	panic("unimplemented")
}

// MyPost implements item.ItemData
func (*itemData) MyPost(userID int) ([]item.Core, error) {
	panic("unimplemented")
}

// Update implements item.ItemData
func (pd *itemData) Update(userID int, itemID int, updatedData item.Core) (item.Core, error) {
	cnv := CoreToData(updatedData)
	// if cnv.UserID != uint(userID) {
	// 	log.Println("ini bukan barang anda")
	// }

	// DB Update(value)
	tx := pd.db.Where("id = ? AND user_id = ?", itemID, userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update barang query error :", tx.Error)
		return item.Core{}, tx.Error

	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return item.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}
