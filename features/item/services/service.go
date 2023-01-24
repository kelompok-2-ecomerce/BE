package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"projects/features/item"
	"projects/helper"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type itemSrv struct {
	data     item.ItemData
	validasi *validator.Validate
}

func New(pd item.ItemData) item.ItemService {
	return &itemSrv{
		data:     pd,
		validasi: validator.New(),
	}
}

// Add implements item.ItemService
func (ps *itemSrv) Add(token interface{}, newItem item.Core, file *multipart.FileHeader) (item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return item.Core{}, errors.New("user tidak ditemukan")
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return item.Core{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			idx := strings.Index(err.Error(), ",")
			msg := err.Error()
			return item.Core{}, errors.New("format input file size tidak diizinkan, size melebihi" + msg[idx+1:])
		}
		extension, err := helper.CheckFileExtension(file.Filename)
		if err != nil {
			return item.Core{}, errors.New("format input file type tidak diizinkan")
		}
		filename := "images/product/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

		photo, err := helper.UploadImageToS3(filename, src)
		if err != nil {
			return item.Core{}, errors.New("format input file type tidak dapat diupload")
		}

		newItem.Image_url = photo

		defer src.Close()
	}

	err := helper.Validasi(helper.ToItemName(newItem))
	if err != nil {
		return item.Core{}, err
	}

	err = helper.Validasi(helper.ToItemStok(newItem))
	if err != nil {
		return item.Core{}, err
	}

	err = helper.Validasi(helper.ToItemHarga(newItem))
	if err != nil {
		return item.Core{}, err
	}

	res, err := ps.data.Add(userID, newItem)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "item not found"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return item.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements item.ItemService
func (*itemSrv) Delete(token interface{}, itemID int) error {
	panic("unimplemented")
}

// GetAllPost implements item.ItemService
func (is *itemSrv) GetAllProducts() ([]item.Core, error) {
	res, err := is.data.GetAllProducts()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []item.Core{}, errors.New(msg)
	}
	return res, nil
}

// MyPost implements item.ItemService
func (is *itemSrv) MyProducts(token interface{}) ([]item.Core, error) {
	res, err := is.data.GetAllProducts()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []item.Core{}, errors.New(msg)
	}
	return res, nil
}

func (is *itemSrv) GetProductByID(token interface{}, productID int) (item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return item.Core{}, errors.New("user tidak ditemukan")
	}
	res, err := is.data.GetProductByID(userID, productID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return item.Core{}, errors.New(msg)
	}
	return res, nil
}

// Update implements item.ItemService
func (ps *itemSrv) Update(token interface{}, itemID int, updatedData item.Core, file *multipart.FileHeader) (item.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return item.Core{}, errors.New("id user not found")
	}
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return item.Core{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			idx := strings.Index(err.Error(), ",")
			msg := err.Error()
			return item.Core{}, errors.New("format input file size tidak diizinkan, size melebihi" + msg[idx+1:])
		}
		extension, err := helper.CheckFileExtension(file.Filename)
		if err != nil {
			return item.Core{}, errors.New("format input file type tidak diizinkan")
		}
		filename := "images/product/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

		photo, err := helper.UploadImageToS3(filename, src)
		if err != nil {
			return item.Core{}, errors.New("format input file type tidak dapat diupload")
		}

		updatedData.Image_url = photo

		defer src.Close()
	}

	if len(updatedData.Nama_Barang) > 0 {
		err := helper.Validasi(helper.ToItemName(updatedData))
		if err != nil {
			return item.Core{}, err
		}
	}

	if updatedData.Stok > 0 {
		err := helper.Validasi(helper.ToItemStok(updatedData))
		if err != nil {
			return item.Core{}, err
		}
	}

	if updatedData.Harga > 0 {
		err := helper.Validasi(helper.ToItemHarga(updatedData))
		if err != nil {
			return item.Core{}, err
		}
	}

	res, err := ps.data.Update(userID, itemID, updatedData)
	if err != nil {
		fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "item not found"
		} else {
			msg = "internal server error"
		}
		return item.Core{}, errors.New(msg)
	}

	return res, nil
}
