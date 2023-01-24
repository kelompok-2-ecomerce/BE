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
		log.Println("add items query error", err.Error())
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
func (pd *itemData) Delete(userID int, itemID int) error {
	item := Item{}
	err := pd.db.Where("id = ? AND user_id = ?", itemID, userID).Delete(&item, itemID)
	if err.Error != nil {
		log.Println("delete item query error :", err.Error)
		return err.Error
	}
	if err.RowsAffected <= 0 {
		log.Println("delete item query error : data not found")
		return errors.New("not found")
	}

	return nil
}

// GetAllPost implements item.ItemData
func (pd *itemData) GetAllItems() ([]item.Core, error) {
	var MyItem []ItemUser
	err := pd.db.Raw("SELECT items.id, items.Nama_Barang, items.image_url, items.deskripsi, items.harga, items.stok, users.nama FROM items JOIN users ON users.id = items.user_id WHERE items .deleted_at IS NULL").Find(&MyItem).Error
	if err != nil {
		return nil, err
	}

	var dataCore = ListModelTOCore(MyItem)

	return dataCore, nil
}

// MyPost implements item.ItemData
func (pd *itemData) MyItem(userID int) ([]item.Core, error) {
	var MyItem []ItemUser
	err := pd.db.Raw("SELECT items.id, items.Nama_Barang, items.image_url, items.deskripsi, items.harga, items.stok, users.nama FROM items JOIN users ON users.id = items.user_id WHERE items.user_id = ?", userID).Find(&MyItem).Error
	if err != nil {
		return nil, err
	}

	var dataCore = ListModelTOCore(MyItem)

	return dataCore, nil
}

// Update implements item.ItemData
func (pd *itemData) Update(userID int, itemID int, updatedData item.Core) (item.Core, error) {
	cnv := CoreToData(updatedData)

	// DB Update(value)
	tx := pd.db.Where("id = ? AND user_id = ?", itemID, userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update barang query error :", tx.Error)
		return item.Core{}, tx.Error

	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update barang query error : data not found")
		return item.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}

// GetID implements item.ItemData
func (*itemData) GetID(ItemID int) (item.Core, error) {
	panic("unimplemented")
}
