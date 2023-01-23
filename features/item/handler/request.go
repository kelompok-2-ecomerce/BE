package handler

import (
	"projects/features/item"
)

type AddUpdatePostingRequest struct {
	Nama_Barang string  `form:"nama_barang"`
	Harga       float64 `form:"harga"`
	Stok        int     `form:"stok"`
	Deskripsi   string  `form:"deskripsi"`
	Image_url   string
}

func (data *AddUpdatePostingRequest) reqToCore() item.Core {
	return item.Core{
		Nama_Barang: data.Nama_Barang,
		Image_url:   data.Image_url,
		Harga:       data.Harga,
		Deskripsi:   data.Deskripsi,
		Stok:        data.Stok,
	}
}

func ToCore(data interface{}) *item.Core {
	res := item.Core{}

	switch data.(type) {
	case AddUpdatePostingRequest:
		cnv := data.(AddUpdatePostingRequest)
		res.Nama_Barang = cnv.Nama_Barang
		res.Image_url = cnv.Image_url
		res.Stok = cnv.Stok
		res.Harga = cnv.Harga
		res.Deskripsi = cnv.Deskripsi

	default:
		return nil
	}

	return &res
}
