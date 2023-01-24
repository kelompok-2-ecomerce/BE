package data

import (
	"projects/features/item"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Nama_Barang string
	Image_url   string
	NamaUser    string
	Alamat      string
	Deskripsi   string
	Harga       float64
	Stok        int
	UserID      uint
}

type User struct {
	gorm.Model
	Comment []Item
}

func CoreToData(data item.Core) Item {
	return Item{
		Model:       gorm.Model{ID: data.ID},
		Nama_Barang: data.Nama_Barang,
		Image_url:   data.Image_url,
		NamaUser:    data.NamaUser,
		Alamat:      data.Alamat,
		Deskripsi:   data.Deskripsi,
		Harga:       data.Harga,
		Stok:        data.Stok,
	}
}
func ToCore(data Item) item.Core {
	return item.Core{
		ID:          data.ID,
		Nama_Barang: data.Nama_Barang,
		Image_url:   data.Image_url,
		NamaUser:    data.NamaUser,
		Alamat:      data.Alamat,
		Deskripsi:   data.Deskripsi,
		Harga:       data.Harga,
		Stok:        data.Stok,
	}
}
