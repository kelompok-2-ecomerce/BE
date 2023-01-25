package handler

import "projects/features/cart"

type CartProductRes struct {
	ItemID      uint    `json:"product_id"`
	ProductName string  `json:"name"`
	ImageUrl    string  `json:"image"`
	Price       float64 `json:"harga"`
	Qty         int     `json:"qty"`
	Total       float64 `json:"total_harga"`
}

func ToCartProductRes(data cart.Core) CartProductRes {
	return CartProductRes{
		ItemID:      data.ItemID,
		ProductName: data.ProductName,
		ImageUrl:    data.ImageUrl,
		Price:       data.Price,
		Qty:         data.Qty,
		Total:       data.Total,
	}
}

func ToCartProductResArr(data []cart.Core) []CartProductRes {
	arrRes := []CartProductRes{}
	for _, v := range data {
		tmp := ToCartProductRes(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
