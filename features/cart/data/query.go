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

	// check exist cart
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

	// check exist product on cart product
	row = cd.db.Raw(`
	SELECT ci.qty  
	FROM cart_items ci 
	WHERE ci.cart_id = ?
	AND ci.item_id = ?;
	`, newCart.ID, productId).Row()
	row.Scan(&newCart.Qty)
	log.Println(newCart.Qty)

	if newCart.Qty > 0 {
		cartProductUpdate := cd.db.Exec(`
		UPDATE cart_items ci
		SET ci.qty = ci.qty + ?
		WHERE ci.cart_id = ?
		AND ci.item_id = ?;
	`, qty, newCart.ID, productId).RowsAffected

		if cartProductUpdate == 0 {
			return cart.Core{}, errors.New("no row affected")

		}
	} else {
		newCartProduct.CartID = newCart.ID
		newCartProduct.ItemID = productId
		newCartProduct.Qty = qty
		err := cd.db.Create(&newCartProduct).Error
		if err != nil {
			log.Println("add cart query error :", err.Error())
			return cart.Core{}, err
		}
	}

	return ToCore(newCart), nil
}
