package data

import (
	"errors"
	"log"
	"projects/features/cart"

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

func (cd *cartData) Add(userID int, productId uint, qty int) (cart.Core, error) {
	// item := Item{
	// 	Model: gorm.Model{ID: productId},
	// }
	// qry := cd.db.Model(&item).Where("stok > ? AND deleted_at IS NULL", userID).Updates("stok")
	productUpdate := cd.db.Exec(`
	UPDATE items i 
	SET i.stok = stok - ?
	WHERE i.id  = ?
	AND i.stok > ?
	AND i.deleted_at IS NULL;
	`, qty, productId, qty).RowsAffected

	if productUpdate == 0 {
		return cart.Core{}, errors.New("not enough stock")

	}
	newCart := Cart{}
	newCartProduct := CartItem{}
	row := cd.db.Raw(`
	SELECT c.id 
	FROM carts c 
	JOIN users u ON u.id = c.user_id 
	WHERE c.user_id = ?
	AND is_ordered = 0
	LIMIT 1;
	`, userID).Row()
	row.Scan(&newCart.ID)
	log.Println(newCart.ID)
	if newCart.ID <= 0 {
		newCart.ItemID = productId
		newCart.UserID = uint(userID)
		newCart.Qty = qty
		err := cd.db.Create(&newCart).Error
		if err != nil {
			log.Println("add cart query error :", err.Error())
			return cart.Core{}, err
		}
	}
	newCartProduct.CartID = newCart.ID
	newCartProduct.ItemID = productId
	newCartProduct.Qty = qty
	err := cd.db.Create(&newCartProduct).Error
	if err != nil {
		log.Println("add cart query error :", err.Error())
		return cart.Core{}, err
	}

	return ToCore(newCart), nil
}
