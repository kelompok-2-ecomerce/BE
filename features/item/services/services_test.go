package services

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"projects/features/item"
	"projects/helper"
	"projects/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewItemData(t)
	userID := 1

	t.Run("success add data", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		resData := item.Core{
			ID:          4,
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		repo.On("Add", userID, inputData).Return(resData, nil).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama_Barang, res.Nama_Barang)
		assert.Equal(t, resData.Harga, res.Harga)
		repo.AssertExpectations(t)
	})

	t.Run("format input file tidak dapat dibuka", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		srv := New(repo)
		header := &multipart.FileHeader{}

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak dapat dibuka")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("format input file size tidak diizinkan", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		srv := New(repo)
		f, err := os.Open("./files/wallpaper.jpg")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/wallpaper.jpg")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file size")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("format input file type", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		srv := New(repo)
		f, err := os.Open("./files/OpenAPI.txt")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/OpenAPI.txt")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file type")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("field required wajib diisi", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
		}
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wajib diisi")
		assert.Equal(t, res, item.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		repo.On("Add", userID, inputData).Return(item.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.Equal(t, res, item.Core{})
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("record not found", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		repo.On("Add", userID, inputData).Return(item.Core{}, errors.New("record not found")).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.Equal(t, res, item.Core{})
		assert.ErrorContains(t, err, "item not found")
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Add(token, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewItemData(t)
	itemID := 1
	t.Run("Item tidak ditemukan", func(t *testing.T) {
		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(pToken, itemID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("Item tidak ditemukan diquery", func(t *testing.T) {

		repo.On("Delete", 1, itemID).Return(errors.New("record not found")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(pToken, itemID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {

		repo.On("Delete", 1, itemID).Return(errors.New("query error")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(pToken, itemID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Berhasil menonaktifkan akun", func(t *testing.T) {
		repo.On("Delete", 1, itemID).Return(nil).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Delete(pToken, itemID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllProducts(t *testing.T) {
	repo := mocks.NewItemData(t)
	t.Run("berhasil melihat list products", func(t *testing.T) {
		resData := []item.Core{
			{
				ID:          1,
				Nama_Barang: "Bibit Anggur",
				Harga:       12000,
				Stok:        99,
				Deskripsi:   "Bibit unggul dari kota asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA SURABAYA",
				NamaUser:    "Rumah Hidroponik",
			},
			{
				ID:          2,
				Nama_Barang: "Bibit Bunga Matahari",
				Harga:       1000,
				Stok:        99,
				Deskripsi:   "Bibit unggul dari daerah asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA MALANG",
				NamaUser:    "Toko Benih Tanaman Online",
			},
			{
				ID:          3,
				Nama_Barang: "Bibit Cerry",
				Harga:       10000,
				Stok:        79,
				Deskripsi:   "Bibit unggul dari daerah asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA MALANG",
				NamaUser:    "Toko Benih Tanaman Online",
			},
		}
		repo.On("GetAllProducts").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetAllProducts()
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, res[0].ID)
		assert.Equal(t, resData[0].Nama_Barang, res[0].Nama_Barang)
		assert.Equal(t, resData[0].NamaUser, res[0].NamaUser)
		assert.Equal(t, resData[0].Stok, res[0].Stok)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := []item.Core{}
		repo.On("GetAllProducts").Return(resData, errors.New("record not found")).Once()

		srv := New(repo)
		res, err := srv.GetAllProducts()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []item.Core{})
		repo.AssertExpectations(t)
	})
	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := []item.Core{}
		repo.On("GetAllProducts").Return(resData, errors.New("query error")).Once()

		srv := New(repo)
		res, err := srv.GetAllProducts()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []item.Core{})
		repo.AssertExpectations(t)
	})
}

func TestMyProducts(t *testing.T) {
	repo := mocks.NewItemData(t)
	userID := 1
	t.Run("berhasil melihat list myproducts", func(t *testing.T) {
		resData := []item.Core{
			{
				ID:          1,
				Nama_Barang: "Bibit Anggur",
				Harga:       12000,
				Stok:        99,
				Deskripsi:   "Bibit unggul dari kota asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA SURABAYA",
				NamaUser:    "Rumah Hidroponik",
			},
			{
				ID:          2,
				Nama_Barang: "Bibit Bunga Matahari",
				Harga:       1000,
				Stok:        99,
				Deskripsi:   "Bibit unggul dari daerah asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA MALANG",
				NamaUser:    "Toko Benih Tanaman Online",
			},
			{
				ID:          3,
				Nama_Barang: "Bibit Cerry",
				Harga:       10000,
				Stok:        79,
				Deskripsi:   "Bibit unggul dari daerah asal",
				Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Alamat:      "KOTA MALANG",
				NamaUser:    "Toko Benih Tanaman Online",
			},
		}
		repo.On("MyProducts", userID).Return(resData, nil).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyProducts(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, res[0].ID)
		assert.Equal(t, resData[0].Nama_Barang, res[0].Nama_Barang)
		assert.Equal(t, resData[0].NamaUser, res[0].NamaUser)
		assert.Equal(t, resData[0].Stok, res[0].Stok)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := []item.Core{}
		repo.On("MyProducts", userID).Return(resData, errors.New("record not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyProducts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []item.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := []item.Core{}
		repo.On("MyProducts", userID).Return(resData, errors.New("query error")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyProducts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []item.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		res, err := srv.MyProducts(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []item.Core{})
	})
}

func TestGetProductByID(t *testing.T) {
	repo := mocks.NewItemData(t)
	userID := 1
	productID := 1
	t.Run("berhasil melihat list products", func(t *testing.T) {
		resData := item.Core{

			ID:          1,
			Nama_Barang: "Bibit Anggur",
			Harga:       12000,
			Stok:        99,
			Deskripsi:   "Bibit unggul dari kota asal",
			Image_url:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
			Alamat:      "KOTA SURABAYA",
			NamaUser:    "Rumah Hidroponik",
		}
		repo.On("GetProductByID", userID, productID).Return(resData, nil).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProductByID(pToken, productID)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama_Barang, res.Nama_Barang)
		assert.Equal(t, resData.NamaUser, res.NamaUser)
		assert.Equal(t, resData.Stok, res.Stok)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := item.Core{}
		repo.On("GetProductByID", userID, productID).Return(resData, errors.New("record not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProductByID(pToken, productID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, item.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := item.Core{}
		repo.On("GetProductByID", userID, productID).Return(resData, errors.New("query error")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProductByID(pToken, productID)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, item.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		res, err := srv.MyProducts(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []item.Core{})
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewItemData(t)
	userID := 1
	productID := 1
	t.Run("update berhasil", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		resData := item.Core{
			ID:          4,
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		repo.On("Update", userID, productID, inputData).Return(resData, nil).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, productID, inputData, header)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama_Barang, res.Nama_Barang)
		assert.Equal(t, resData.Harga, res.Harga)
		repo.AssertExpectations(t)
	})

	t.Run("format input file tidak dapat dibuka", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		srv := New(repo)
		header := &multipart.FileHeader{}

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak dapat dibuka")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("format input file size tidak diizinkan", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		srv := New(repo)
		f, err := os.Open("./files/wallpaper.jpg")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/wallpaper.jpg")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, productID, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file size")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("format input file type", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		srv := New(repo)
		f, err := os.Open("./files/OpenAPI.txt")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/OpenAPI.txt")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, productID, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file type")
		assert.Equal(t, res, item.Core{})
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		resData := item.Core{}
		repo.On("Update", userID, productID, inputData).Return(resData, errors.New("server error")).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, productID, inputData, header)
		assert.NotNil(t, err)
		assert.Equal(t, res, item.Core{})
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("record not found", func(t *testing.T) {
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}

		resData := item.Core{}
		repo.On("Update", userID, productID, inputData).Return(resData, errors.New("record not found")).Once()
		srv := New(repo)
		f, err := os.Open("./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", "./files/ImYoonAh.JPG")
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = io.Copy(part, f)
		if err != nil {
			log.Fatal(err.Error())
		}
		writer.Close()
		req, _ := http.NewRequest("PUT", "/users", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, productID, inputData, header)
		assert.NotNil(t, err)
		assert.Equal(t, res, item.Core{})
		assert.ErrorContains(t, err, "item not found")
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := item.Core{
			Nama_Barang: "Bibit Anggur",
			Stok:        100,
			Harga:       5000,
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Update(token, productID, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "id user not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}
