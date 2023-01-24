package handler

import (
	"projects/features/item"
)

type ItemResponse struct {
	Nama_Barang string  `json:"nama_barang"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"deskripsi"`
	Image_url   string  `json:"image"`
	Nama_User   string  `json:"nama_user"`
}
type ItemResponses struct {
	ID          uint    `json:"id"`
	Nama_Barang string  `json:"nama_barang"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"deskripsi"`
	Image_url   string  `json:"image"`
	Nama_User   string  `json:"nama_user"`
}
type AddItemResponse struct {
	Nama_Barang string  `json:"nama_barang"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"deskripsi"`
	Image_url   string  `json:"image"`
}
type updateItemResponse struct {
	Nama_Barang string  `json:"nama_barang"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"deskripsi"`
	Image_url   string  `json:"image"`
}

func ToResponse(feature string, item item.Core) interface{} {
	switch feature {
	case "add":
		return AddItemResponse{
			Nama_Barang: item.Nama_Barang,
			Harga:       item.Harga,
			Stok:        item.Stok,
			Deskripsi:   item.Deskripsi,
			Image_url:   item.Image_url,
		}
	case "update":
		return updateItemResponse{
			Nama_Barang: item.Nama_Barang,
			Harga:       item.Harga,
			Stok:        item.Stok,
			Deskripsi:   item.Deskripsi,
			Image_url:   item.Image_url,
		}
	default:
		return ItemResponse{

			Nama_Barang: item.Nama_Barang,
			Harga:       item.Harga,
			Stok:        item.Stok,
			Deskripsi:   item.Deskripsi,
			Image_url:   item.Image_url,
			Nama_User:   item.Nama,
		}
	}
}

func ListItemCoreToItemRespon(dataCore item.Core) ItemResponse { // data user core yang ada di controller yang memanggil user repository
	return ItemResponse{
		// ID:          dataCore.ID,
		Nama_Barang: dataCore.Nama_Barang,
		Image_url:   dataCore.Image_url,
		Harga:       dataCore.Harga,
		Stok:        dataCore.Stok,
		Deskripsi:   dataCore.Deskripsi,
		Nama_User:   dataCore.Nama,
	}
}
func ListItemsCoreToItemRespon(dataCore item.Core) ItemResponses { // data user core yang ada di controller yang memanggil user repository
	return ItemResponses{
		ID:          dataCore.ID,
		Nama_Barang: dataCore.Nama_Barang,
		Image_url:   dataCore.Image_url,
		Harga:       dataCore.Harga,
		Stok:        dataCore.Stok,
		Deskripsi:   dataCore.Deskripsi,
		Nama_User:   dataCore.Nama,
	}
}
func ListItemsCoreToItemsRespon(dataCore []item.Core) []ItemResponse {
	var ResponData []ItemResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListItemCoreToItemRespon(value))
	}
	return ResponData
}
