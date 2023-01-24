package data

import (
	"projects/features/cart"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Qty    string
	ItemID uint
	UserID uint
}

type CartUser struct {
	ID          uint
	Nama_Barang string
	Qty         string
	Nama        string
	CreatedAt   time.Time
}

func CoreToData(data cart.Core) Cart {
	return Cart{
		Model:  gorm.Model{ID: data.ID},
		Qty:    data.Qty,
		ItemID: data.ItemID,
	}
}

func (dataModel *CartUser) ModelsToCore() cart.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return cart.Core{
		ID:          dataModel.ID,
		Nama_Barang: dataModel.Nama_Barang,
		Qty:         dataModel.Qty,
		Nama:        dataModel.Nama,
	}
}

func ListModelTOCore(dataModel []CartUser) []cart.Core { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []cart.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
