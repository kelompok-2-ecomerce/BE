package data

import (
	"projects/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	IsOrdered   bool `gorm:"default:false"`
	ProductName string
	ImageUrl    string
	Price       float64
	Qty         int
	Total       float64
	Items       []Item `gorm:"many2many:cart_items;"`
	UserID      uint
	ItemID      uint
}

type CartItem struct {
	CartID    uint
	ItemID    uint
	Qty       int
	DeletedAt gorm.DeletedAt
}

type User struct {
	gorm.Model
	Carts []Cart
}

type Item struct {
	gorm.Model
}

func ToCore(data Cart) cart.Core {
	return cart.Core{
		ID:          data.ID,
		ProductName: data.ProductName,
		ImageUrl:    data.ImageUrl,
		Price:       data.Price,
		Qty:         data.Qty,
		Total:       data.Total,
	}
}

func CoreToData(data cart.Core) Cart {
	return Cart{
		ProductName: data.ProductName,
		ImageUrl:    data.ImageUrl,
		Price:       data.Price,
		Qty:         data.Qty,
		ItemID:      data.ItemID,
		Total:       data.Total,
	}
}

func ToCoreArr(data []Cart) []cart.Core {
	arrRes := []cart.Core{}
	for _, v := range data {
		tmp := ToCore(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
