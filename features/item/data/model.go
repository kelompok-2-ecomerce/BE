package data

import (
	"projects/features/item"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Nama_Barang string
	Image_url   string
	Deskripsi   string
	Harga       float64
	Stok        int
	UserID      uint
}

type ItemUser struct {
	ID          uint
	Nama_Barang string
	Image_url   string
	Nama        string
	Deskripsi   string
	Harga       float64
	Stok        int
	CreatedAt   time.Time
}

func CoreToData(data item.Core) Item {
	return Item{
		Model:       gorm.Model{ID: data.ID},
		Nama_Barang: data.Nama_Barang,
		Image_url:   data.Image_url,
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
		Deskripsi:   data.Deskripsi,
		Harga:       data.Harga,
		Stok:        data.Stok,
	}
}

func (dataModel *ItemUser) ModelsToCore() item.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return item.Core{
		ID:          dataModel.ID,
		Nama_Barang: dataModel.Nama_Barang,
		Image_url:   dataModel.Image_url,
		Deskripsi:   dataModel.Deskripsi,
		Harga:       dataModel.Harga,
		Stok:        dataModel.Stok,
		Nama:        dataModel.Nama,
	}
}

func ListModelTOCore(dataModel []ItemUser) []item.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []item.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
