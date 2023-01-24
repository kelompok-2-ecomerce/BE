package handler

import (
	"projects/features/cart"
)

type CartResponse struct {
	Qty         string `json:"qty"`
	ItemID      uint   `json:"itemid"`
	Nama_Barang string `json:"nama_barang"`
	Nama        string `json:"nama_user"`
}

type AddCartResponse struct {
	Qty    string `json:"qty"`
	ItemID uint   `json:"itemid"`
}
type updateCartResponse struct {
	Qty    string `json:"qty"`
	ItemID uint   `json:"itemid"`
}

func ToResponse(feature string, cart cart.Core) interface{} {
	switch feature {
	case "add":
		return AddCartResponse{
			Qty:    cart.Qty,
			ItemID: cart.ItemID,
		}
	case "update":
		return updateCartResponse{
			Qty:    cart.Qty,
			ItemID: cart.ItemID,
		}
	default:
		return CartResponse{

			Qty:         cart.Qty,
			ItemID:      cart.ItemID,
			Nama_Barang: cart.Nama_Barang,
			Nama:        cart.Nama,
		}
	}
}

func ListCartCoreToCartRespon(dataCore cart.Core) CartResponse { // data user core yang ada di controller yang memanggil user repository
	return CartResponse{
		// ID:          dataCore.ID,
		Nama_Barang: dataCore.Nama_Barang,
		Qty:         dataCore.Qty,
		ItemID:      dataCore.ItemID,
		Nama:        dataCore.Nama,
	}
}

func ListItemsCoreToItemsRespon(dataCore []cart.Core) []CartResponse {
	var ResponData []CartResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListCartCoreToCartRespon(value))
	}
	return ResponData
}
