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
		log.Println("add item query error", err.Error())
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
func (id *itemData) Delete(userID int, itemID int) error {
	product := Item{
		Model: gorm.Model{ID: uint(itemID)},
	}
	qry := id.db.Where("user_id = ?", userID).Delete(&product)
	if qry.RowsAffected <= 0 {
		log.Println("delete product query error : data not found")
		return errors.New("not found")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete product query error :", err.Error())
		return err
	}
	return nil
}

// GetAllProducts implements item.ItemData
func (id *itemData) GetAllProducts() ([]item.Core, error) {
	res := []Item{}

	err := id.db.Raw(`
	SELECT i.id , i.nama_barang , i.image_url , u.nama "NamaUser", u.alamat , i.deskripsi ,i.harga , i.stok 
	FROM items i 
	JOIN users u ON u.id = i.user_id
	WHERE i.deleted_at IS NULL
	ORDER BY i.id DESC;
	`).Scan(&res).Error
	if err != nil {
		log.Println("list products query error :", err.Error())
		return []item.Core{}, err
	}

	return ToCoreArr(res), nil
}

// MyPost implements item.ItemData
func (id *itemData) MyProducts(userID int) ([]item.Core, error) {
	res := []Item{}
	err := id.db.Raw(`
	SELECT i.id , i.nama_barang , i.image_url , u.nama "NamaUser", u.alamat , i.deskripsi ,i.harga , i.stok 
	FROM items i 
	JOIN users u ON u.id = i.user_id
	WHERE i.deleted_at IS NULL
	AND u.id = ?
	ORDER BY i.updated_at;
	`, userID).Scan(&res).Error
	if err != nil {
		log.Println("list myproducts query error :", err.Error())
		return []item.Core{}, err
	}

	return ToCoreArr(res), nil
}

func (id *itemData) GetProductByID(userID int, productID int) (item.Core, error) {
	res := Item{}

	err := id.db.Raw(`
	SELECT i.id , i.nama_barang , i.image_url , u.nama "NamaUser", u.alamat , i.deskripsi ,i.harga , i.stok 
	FROM items i 
	JOIN users u ON u.id = i.user_id
	WHERE i.deleted_at IS NULL
	AND u.id = ?
	AND i.id = ?
	ORDER BY i.id DESC;
	`, userID, productID).Scan(&res).Error
	if err != nil {
		log.Println("list myproducts query error :", err.Error())
		return item.Core{}, err
	}
	return ToCore(res), nil
}

// Update implements item.ItemData
func (pd *itemData) Update(userID int, itemID int, updatedData item.Core) (item.Core, error) {
	cnv := CoreToData(updatedData)

	// DB Update(value)
	tx := pd.db.Where("id = ? AND user_id = ?", itemID, userID).Updates(&cnv)
	if tx.Error != nil {
		log.Println("update product query error :", tx.Error)
		return item.Core{}, tx.Error

	}

	// Rows affected checking
	if tx.RowsAffected <= 0 {
		log.Println("update product query error : data not found")
		return item.Core{}, errors.New("not found")
	}

	// return result converting cnv to book.Core
	return ToCore(cnv), nil
}
