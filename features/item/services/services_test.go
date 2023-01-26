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
	data := mocks.NewItemData(t)
	t.Run("Berhasil Menambahkan Item", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   1,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "Baju",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}

		Respon := item.Core{
			ID:          1,
			Nama_Barang: inputData.Nama_Barang,
			Image_url:   inputData.Image_url,
			Deskripsi:   inputData.Deskripsi,
			Stok:        inputData.Stok,
			Harga:       inputData.Harga,
		}

		data.On("Add", sample.ID, inputData).Return(Respon, nil).Once()
		srv := New(data)
		f, _ := os.Open("./files/ImYoonAh.JPG")

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
		req, _ := http.NewRequest("POST", "/products", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")
		_, token := helper.GenerateJWT(sample.ID)
		useToken := token.(*jwt.Token)
		useToken.Valid = true
		res, err := srv.Add(useToken, inputData, header)
		assert.Nil(t, err)
		assert.Equal(t, Respon.ID, res.ID)
		assert.Equal(t, Respon.NamaUser, res.NamaUser)
		data.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   1,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "Baju",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}
		srv := New(data)
		_, token := helper.GenerateJWT(sample.ID)
		f, _ := os.Open("./files/ImYoonAh.JPG")

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
		req, _ := http.NewRequest("POST", "/products", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")
		res, err := srv.Add(token, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user tidak ditemukan")
		assert.Equal(t, uint(0), res.ID) //perbandingan
	})
	t.Run("format input file tidak dapat dibuka", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "Baju",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}

		srv := New(data)
		f, _ := os.Open("./files/OpenAPI.txt")

		defer f.Close()

		file := &multipart.FileHeader{}
		_, token := helper.GenerateJWT(sample.ID)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, file)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file tidak dapat dibuka")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("format input file type tidak diizinkan", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "Baju",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}

		srv := New(data)
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

		_, token := helper.GenerateJWT(sample.ID)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file type tidak diizinkan")
		assert.Equal(t, res, item.Core{})
	})
	t.Run("format input file type tidak dapat diupload", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "Baju",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}
		data.On("Add", sample.ID, inputData).Return(item.Core{}, errors.New("format input file type tidak dapat diupload")).Once()

		srv := New(data)
		f, _ := os.Open("./files/ImYoonAh.JPG")

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
		req, _ := http.NewRequest("POST", "/products", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		_, header, _ := req.FormFile("image")
		_, token := helper.GenerateJWT(sample.ID)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, header)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format input file type tidak dapat diupload")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("masalah di server", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		inputData := item.Core{
			ID:          1,
			Nama_Barang: "1",
			Image_url:   "www.google.com",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    sample.Name,
		}
		data.On("Add", sample.ID, inputData).Return(item.Core{}, errors.New("internal server error")).Once()
		srv := New(data) //new service

		_, token := helper.GenerateJWT(sample.ID)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestDeletePost(t *testing.T) {
	data := mocks.NewItemData(t)

	srv := New(data)
	t.Run("Delete Success", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.Nil(t, err)

		data.AssertExpectations(t)
	})

	t.Run("Delete Error", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(errors.New("user not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)

		data.AssertExpectations(t)
	})
	t.Run("Delete Error", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "product tidak ditemukan")
		data.AssertExpectations(t)
	})
	t.Run("Delete server error", func(t *testing.T) {
		data.On("Delete", 1, 1).Return(errors.New("terjadi kesalahan pada server")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(token, 1)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server")
		data.AssertExpectations(t)
	})
}

func TestMyItems(t *testing.T) {
	data := mocks.NewItemData(t)

	srv := New(data)

	// Case: user ingin melihat list buku yang dimilikinya
	t.Run("item succesfully", func(t *testing.T) {
		resData := []item.Core{
			{
				ID:          1,
				Nama_Barang: "Baju",
				Image_url:   "www.google.com",
				Deskripsi:   "biru asik",
				Stok:        1,
				Harga:       20000,
			},
			{
				ID:          2,
				Nama_Barang: "celana",
				Image_url:   "www.google.com",
				Deskripsi:   "biru asik",
				Stok:        1,
				Harga:       40000,
			},
		}

		// Programming input and return repo
		data.On("MyProducts", 1).Return(resData, nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.MyProducts(token)

		// Test
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, actual[0].ID)
		assert.Equal(t, resData[0].Nama_Barang, actual[0].Nama_Barang)
		assert.Equal(t, resData[1].ID, actual[1].ID)
		assert.Equal(t, resData[1].Harga, actual[1].Harga)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		srv := New(data)
		_, token := helper.GenerateJWT(sample.ID)

		res, err := srv.MyProducts(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user tidak ditemukan")
		assert.Equal(t, uint(0), res) //perba
	})

	t.Run("not found", func(t *testing.T) {

		data.On("MyProducts", 1).Return([]item.Core{}, errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyProducts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Equal(t, uint(0), res)
	})

	t.Run("terjadi kesalahan pada server", func(t *testing.T) {

		data.On("MyProducts", 1).Return([]item.Core{}, errors.New("terjadi kesalahan pada server")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.MyProducts(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server")
		assert.Equal(t, uint(0), res)
	})

}

// func TestUpdateItem(t *testing.T) {

// 	data := mocks.NewItemData(t)
// 	srv := New(data)

// 	t.Run("Update successfully", func(t *testing.T) {
// 		input := item.Core{Nama_Barang: "One Piece"}
// 		resData := item.Core{
// 			ID:          1,
// 			Nama_Barang: "One Piece",
// 			Image_url:   "www.google.com",
// 			Nama:        "fajar1411",
// 		}

// 		data.On("update", 1, 1, input).Return(resData, nil).Once()
// 		_, token := helper.GenerateJWT(1)
// 		useToken := token.(*jwt.Token)
// 		useToken.Valid = true
// 		srv := New(data)
// 		RES, err := srv.Update(useToken, 1, input)

// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.Nama_Barang, RES.Nama_Barang)
// 		assert.Equal(t, resData.ID, RES.ID)
// 		assert.Equal(t, resData.Nama, RES.Nama)

// 		data.AssertExpectations(t)
// 	})
// 	t.Run("Update error user not found", func(t *testing.T) {
// 		input := item.Core{Nama_Barang: "One Piece"}
// 		// resData := item.Core{
// 		// 	ID:          1,
// 		// 	Nama_Barang: "One Piece",
// 		// 	Image_url:   "www.google.com",
// 		// 	Nama:        "fajar1411",
// 		// }

// 		token := jwt.New(jwt.SigningMethodHS256)
// 		actual, err := srv.Update(token, 1, input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "id user not found")
// 		assert.Empty(t, actual)
// 	})
// 	t.Run("Update error post not found", func(t *testing.T) {
// 		// Programming input and return repo
// 		input := item.Core{Nama_Barang: "One Piece"}
// 		data.On("Update", 1, 1, input).Return(item.Core{}, errors.New("not found")).Once()

// 		// Program service
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		actual, err := srv.Update(token, 1, input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "item not found")
// 		assert.Empty(t, actual)
// 		data.AssertExpectations(t)
// 	})
// 	t.Run("Update error internal server", func(t *testing.T) {
// 		// Programming input and return repo
// 		input := item.Core{ID: 1, Nama_Barang: "One Piece"}
// 		data.On("Update", 1, 1, input).Return(item.Core{}, errors.New("internal server error")).Once()

// 		// Program service
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		actual, err := srv.Update(pToken, 1, input)

//			// Test
//			assert.NotNil(t, err)
//			assert.ErrorContains(t, err, "internal server error")
//			assert.Empty(t, actual)
//			data.AssertExpectations(t)
//		})
//	}
func TestGetAllItems(t *testing.T) {
	data := mocks.NewItemData(t)

	srv := New(data)

	t.Run("Item succesfully", func(t *testing.T) {
		resData := []item.Core{
			{
				ID:          1,
				Nama_Barang: "Baju",
				Image_url:   "www.google.com",
				Deskripsi:   "biru asik",
				Stok:        1,
				Harga:       20000,
			},
			{
				ID:          2,
				Nama_Barang: "celana",
				Image_url:   "www.google.com",
				Deskripsi:   "biru asik",
				Stok:        1,
				Harga:       40000,
			},
		}
		data.On("GetAllProducts").Return(resData, nil).Once()

		res, err := srv.GetAllProducts()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		data.AssertExpectations(t)

	})
	t.Run("not found", func(t *testing.T) {
		data.On("GetAllProducts").Return(nil, errors.New("not found")).Once()

		res, err := srv.GetAllProducts()
		assert.NotNil(t, err)
		assert.EqualError(t, err, "data tidak ditemukan")
		assert.Nil(t, res)
	})
	t.Run("Get all item error server", func(t *testing.T) {
		// Programming input and return repo
		data.On("GetAllProducts").Return([]item.Core{}, errors.New("internal server error")).Once()

		// Program service
		actual, err := srv.GetAllProducts()

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "terjadi kesalahan pada server")
		assert.Nil(t, actual)

	})
}

func TestGetProductByID(t *testing.T) {
	data := mocks.NewItemData(t)

	srv := New(data)

	// Case: user ingin melihat list buku yang dimilikinya
	t.Run("item succesfully", func(t *testing.T) {
		resData := item.Core{
			ID:          1,
			Nama_Barang: "baju",
			Deskripsi:   "biru asik",
			Stok:        1,
			Harga:       20000,
			NamaUser:    "fajar",
		}

		// Programming input and return repo
		data.On("GetProductByID", 1, 1).Return(resData, nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		actual, err := srv.GetProductByID(pToken, 1)

		// Test
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.Equal(t, resData.ID, actual.ID)
		data.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		type SampleUsers struct {
			ID   int
			Name string
		}
		sample := SampleUsers{
			ID:   4,
			Name: "fajar1411",
		}
		srv := New(data)
		_, token := helper.GenerateJWT(sample.ID)

		res, err := srv.GetProductByID(token, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "user tidak ditemukan")
		assert.Equal(t, uint(0), res) //perba
	})

	t.Run("not found", func(t *testing.T) {

		data.On("GetProductByID", 1, 1).Return(item.Core{}, errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProductByID(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Equal(t, uint(0), res)
	})

	t.Run("terjadi kesalahan pada server", func(t *testing.T) {

		data.On("GetProductByID", 1, 1).Return(item.Core{}, errors.New("terjadi kesalahan pada server")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetProductByID(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server")
		assert.Equal(t, uint(0), res)
	})

}

// func TestIDItem(t *testing.T) {
// 	resData := item.Core{
// 		ID:          1,
// 		Nama_Barang: "baju",
// 		Deskripsi:   "biru asik",
// 		Stok:        1,
// 		Harga:       20000,
// 		Nama:        "fajar",
// 	}

// 	data := mocks.NewItemData(t)

// 	srv := New(data)
// 	t.Run("getid successfully", func(t *testing.T) {
// 		data.On("GetID", 1).Return(resData, nil).Once()

// 		actual, err := srv.GetID(1)

// 		assert.NoError(t, err)
// 		assert.Equal(t, resData.ID, actual.ID)
// 		data.AssertExpectations(t)

// 	})
// 	t.Run("not found", func(t *testing.T) {
// 		// resData := item.Core{
// 		// 	ID:          1,
// 		// 	Nama_Barang: "baju",
// 		// 	Deskripsi:   "biru asik",
// 		// 	Stok:        1,
// 		// 	Harga:       20000,
// 		// 	Nama:        "fajar",
// 		// }
// 		data.On("GetID", 1).Return(item.Core{}, errors.New("not found")).Once()

// 		res, err := srv.GetID(1)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "ID Product not found")
// 		assert.Equal(t, uint(0), res.ID)
// 		data.AssertExpectations(t)
// 	})
// 	t.Run("terdapat masalah pada server", func(t *testing.T) {
// 		// resData := item.Core{
// 		// 	ID:          1,
// 		// 	Nama_Barang: "baju",
// 		// 	Deskripsi:   "biru asik",
// 		// 	Stok:        1,
// 		// 	Harga:       20000,
// 		// 	Nama:        "fajar",
// 		// }
// 		data.On("GetID", 1).Return(item.Core{}, errors.New("terdapat masalah pada server")).Once()

// 		res, err := srv.GetID(1)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "terdapat masalah pada server")
// 		assert.Equal(t, uint(0), res.ID)
// 		data.AssertExpectations(t)
// 	})
// }
