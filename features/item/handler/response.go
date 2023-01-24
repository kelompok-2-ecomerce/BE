package handler

import (
	"projects/features/item"
)

type ItemResponse struct {
	ID          uint    `json:"id"`
	Nama_Barang string  `json:"name"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"description"`
	Image_url   string  `json:"image"`
	Alamat      string  `json:"address"`
	Nama_User   string  `json:"penjual"`
}
type AddItemResponse struct {
	Nama_Barang string  `json:"name"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"description"`
	Image_url   string  `json:"image"`
}
type updateItemResponse struct {
	Nama_Barang string  `json:"name"`
	Harga       float64 `json:"harga"`
	Stok        int     `json:"stok"`
	Deskripsi   string  `json:"description"`
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
			ID:          item.ID,
			Nama_Barang: item.Nama_Barang,
			Harga:       item.Harga,
			Stok:        item.Stok,
			Deskripsi:   item.Deskripsi,
			Image_url:   item.Image_url,
			Nama_User:   item.NamaUser,
			Alamat:      item.Alamat,
		}
	}
}

func ListPostCoreToPostRespon(dataCore item.Core) ItemResponse { // data user core yang ada di controller yang memanggil user repository
	return ItemResponse{
		ID:          dataCore.ID,
		Nama_Barang: dataCore.Nama_Barang,
		Image_url:   dataCore.Image_url,
		Harga:       dataCore.Harga,
		Stok:        dataCore.Stok,
		Alamat:      dataCore.Alamat,
		Deskripsi:   dataCore.Deskripsi,
		Nama_User:   dataCore.NamaUser,
	}
}
func ListPostCoreToPostsRespon(dataCore []item.Core) []ItemResponse {
	var ResponData []ItemResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ListPostCoreToPostRespon(value))
	}
	return ResponData
}
