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

func (cd *cartData) CartExits(userID int) uint {
	newCart := Cart{}
	row := cd.db.Raw(`
	SELECT c.id 
	FROM carts c 
	JOIN users u ON u.id = c.user_id 
	WHERE c.user_id = ?
	AND is_ordered = 0
	LIMIT 1;
	`, userID).Row()
	row.Scan(&newCart.ID)
	return newCart.ID
}

func (cd *cartData) CheckStock(productId uint, qty int) int64 {
	productUpdate := cd.db.Exec(`
	UPDATE items i 
	SET i.stok = stok - ?
	WHERE i.id  = ?
	AND i.stok >= ?
	AND i.deleted_at IS NULL;
	`, qty, productId, qty).RowsAffected

	return productUpdate
}

func (cd *cartData) Add(userID int, productId uint, qty int) (cart.Core, error) {
	// check stok product and decrease if enough stock
	if cd.CheckStock(productId, qty) == 0 {
		return cart.Core{}, errors.New("not enough stock")
	}

	// check exist cart
	newCart := Cart{}
	newCart.ID = cd.CartExits(userID)
	if newCart.ID <= 0 {
		newCart.ItemID = productId
		newCart.UserID = uint(userID)
		err := cd.db.Create(&newCart).Error
		if err != nil {
			log.Println("add cart query error :", err.Error())
			return cart.Core{}, err
		}
	}

	// check exist product on cart product
	newCartProduct := CartItem{}
	row := cd.db.Raw(`
	SELECT ci.qty  
	FROM cart_items ci 
	WHERE ci.cart_id = ?
	AND ci.item_id = ?;
	`, newCart.ID, productId).Row()
	row.Scan(&newCart.Qty)

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

func (cd *cartData) GetMyCart(userID int) ([]cart.Core, error) {
	// check exist cart
	newCart := Cart{}
	newCart.ID = cd.CartExits(userID)
	if newCart.ID <= 0 {
		msg := "data not found"
		log.Println("query get ny cart " + msg)
		return []cart.Core{}, errors.New(msg)
	}

	listProduct := []Cart{}
	err := cd.db.Raw(`
	SELECT  i.id "ItemID" , i.nama_barang "ProductName", i.image_url, i.harga "Price", ci.qty, COALESCE(SUM(i.harga*ci.qty),0) "Total"
	FROM items i 
	JOIN cart_items ci ON ci.item_id = i.id 
	WHERE ci.cart_id = ?
	GROUP BY ci.cart_id , ci.item_id
	ORDER BY i.updated_at;
	`, newCart.ID).Scan(&listProduct).Error
	if err != nil {
		log.Println("list cart product query error :", err.Error())
		return []cart.Core{}, err
	}

	return ToCoreArr(listProduct), nil
}

func (cd *cartData) CartProductExits(userID int, productId uint) int {
	// check exist cart
	newCart := Cart{}
	newCart.ID = cd.CartExits(userID)
	if newCart.ID <= 0 {
		return int(newCart.ID)
	}
	row := cd.db.Raw(`
	SELECT ci.qty  
	FROM cart_items ci 
	WHERE ci.cart_id = ?
	AND ci.item_id = ?;
	`, newCart.ID, productId).Row()
	row.Scan(&newCart.Qty)
	return newCart.Qty
}

func (cd *cartData) UpdateProductCart(userID int, productId uint, qty int) error {
	// check exist cart
	newCart := Cart{}
	newCart.ID = cd.CartExits(userID)
	if newCart.ID <= 0 {
		msg := "data not found"
		log.Println("query get ny cart " + msg)
		return errors.New(msg)
	}

	// get qty current
	QtyCurrent := cd.CartProductExits(userID, productId)
	if qty > QtyCurrent {
		QtyOnGoing := qty - QtyCurrent
		log.Println(QtyOnGoing)
		// check stok product and decrease if enough stock
		if cd.CheckStock(productId, QtyOnGoing) == 0 {
			return errors.New("not enough stock on update product cart")
		}
	} else {
		// increase stock product
		QtyOnGoing := QtyCurrent - qty
		productUpdate := cd.db.Exec(`
		UPDATE items i 
		SET i.stok = stok + ?
		WHERE i.id  = ?
		AND i.deleted_at IS NULL;
		`, QtyOnGoing, productId).RowsAffected
		if productUpdate == 0 {
			return errors.New("no row affected")
		}
	}

	cartProductUpdate := cd.db.Exec(`
		UPDATE cart_items ci
		SET ci.qty = ?
		WHERE ci.cart_id = ?
		AND ci.item_id = ?;
		`, qty, newCart.ID, productId).RowsAffected

	if cartProductUpdate == 0 {
		return errors.New("no row affected")
	}
	return nil
}
