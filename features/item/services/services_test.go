package services

// import (
// 	"errors"
// 	"projects/features/item"
// 	"projects/helper"
// 	"projects/mocks"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAdd(t *testing.T) {
// 	data := mocks.NewItemData(t)
// 	t.Run("Berhasil Menambahkan Item", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID   int
// 			Name string
// 		}
// 		sample := SampleUsers{
// 			ID:   1,
// 			Name: "fajar1411",
// 		}
// 		inputData := item.Core{
// 			ID:          1,
// 			Nama_Barang: "Baju",
// 			Image_url:   "www.google.com",
// 			Deskripsi:   "biru asik",
// 			Stok:        1,
// 			Harga:       20000,
// 			NamaUser:    sample.Name,
// 		}

// 		Respon := item.Core{
// 			ID:          1,
// 			Nama_Barang: inputData.Nama_Barang,
// 			Image_url:   inputData.Image_url,
// 			Deskripsi:   inputData.Deskripsi,
// 			Stok:        inputData.Stok,
// 			Harga:       inputData.Harga,
// 		}

// 		_, token := helper.GenerateJWT(sample.ID)
// 		useToken := token.(*jwt.Token)
// 		useToken.Valid = true

// 		data.On("Add", sample.ID, inputData).Return(Respon, nil).Once()
// 		srv := New(data)

// 		res, err := srv.Add(useToken, inputData)
// 		assert.Nil(t, err)
// 		assert.Equal(t, Respon.ID, res.ID)
// 		assert.Equal(t, Respon.Nama, res.Nama)
// 		data.AssertExpectations(t)
// 	})
// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID   int
// 			Name string
// 		}
// 		sample := SampleUsers{
// 			ID:   1,
// 			Name: "fajar1411",
// 		}
// 		inputData := item.Core{
// 			ID:          1,
// 			Nama_Barang: "Baju",
// 			Image_url:   "www.google.com",
// 			Deskripsi:   "biru asik",
// 			Stok:        1,
// 			Harga:       20000,
// 			NamaUser:    sample.Name,
// 		}
// 		srv := New(data)
// 		_, token := helper.GenerateJWT(sample.ID)

// 		res, err := srv.Add(token, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "user not found")
// 		assert.Equal(t, uint(0), res.ID) //perbandingan
// 	})

// 	t.Run("data not found", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID   int
// 			Name string
// 		}
// 		sample := SampleUsers{
// 			ID:   4,
// 			Name: "fajar1411",
// 		}
// 		inputData := item.Core{
// 			ID:          1,
// 			Nama_Barang: "Baju",
// 			Image_url:   "www.google.com",
// 			Deskripsi:   "biru asik",
// 			Stok:        1,
// 			Harga:       20000,
// 			Nama:        sample.Name,
// 		}
// 		data.On("Add", sample.ID, inputData).Return(item.Core{}, errors.New("data not found")).Once()

// 		srv := New(data)

// 		_, token := helper.GenerateJWT(sample.ID)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Add(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "Items not found")
// 		assert.Equal(t, uint(0), res.ID)
// 		data.AssertExpectations(t)
// 	})
// 	t.Run("masalah di server", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID   int
// 			Name string
// 		}
// 		sample := SampleUsers{
// 			ID:   4,
// 			Name: "fajar1411",
// 		}
// 		inputData := item.Core{
// 			ID:          1,
// 			Nama_Barang: "Baju",
// 			Image_url:   "www.google.com",
// 			Deskripsi:   "biru asik",
// 			Stok:        1,
// 			Harga:       20000,
// 			Nama:        sample.Name,
// 		}
// 		data.On("Add", sample.ID, inputData).Return(item.Core{}, errors.New("internal server error")).Once()
// 		srv := New(data) //new service

// 		_, token := helper.GenerateJWT(sample.ID)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Add(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "server")
// 		assert.Equal(t, uint(0), res.ID)
// 		data.AssertExpectations(t)
// 	})
// }

// func TestDeletePost(t *testing.T) {
// 	data := mocks.NewItemData(t)

// 	srv := New(data)
// 	t.Run("Delete Success", func(t *testing.T) {
// 		data.On("Delete", 1, 1).Return(nil).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(token, 1)

// 		assert.Nil(t, err)

// 		data.AssertExpectations(t)
// 	})

// 	t.Run("Delete Error", func(t *testing.T) {
// 		data.On("Delete", 1, 1).Return(errors.New("user not found")).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(token, 1)

// 		assert.NotNil(t, err)

// 		data.AssertExpectations(t)
// 	})
// 	t.Run("Delete Error", func(t *testing.T) {
// 		data.On("Delete", 1, 1).Return(errors.New("not found")).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(token, 1)

// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "item not found")
// 		data.AssertExpectations(t)
// 	})
// 	t.Run("Delete server error", func(t *testing.T) {
// 		data.On("Delete", 1, 1).Return(errors.New("internal server error")).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(token, 1)

// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "internal server error")
// 		data.AssertExpectations(t)
// 	})
// }

// func TestMyItems(t *testing.T) {
// 	data := mocks.NewItemData(t)

// 	srv := New(data)

// 	// Case: user ingin melihat list buku yang dimilikinya
// 	t.Run("Post succesfully", func(t *testing.T) {
// 		resData := []item.Core{
// 			{
// 				ID:          1,
// 				Nama_Barang: "Baju",
// 				Image_url:   "www.google.com",
// 				Deskripsi:   "biru asik",
// 				Stok:        1,
// 				Harga:       20000,
// 			},
// 			{
// 				ID:          2,
// 				Nama_Barang: "celana",
// 				Image_url:   "www.google.com",
// 				Deskripsi:   "biru asik",
// 				Stok:        1,
// 				Harga:       40000,
// 			},
// 		}

// 		// Programming input and return repo
// 		data.On("MyItem", 1).Return(resData, nil).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		actual, err := srv.MyItem(token)

// 		// Test
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData[0].ID, actual[0].ID)
// 		assert.Equal(t, resData[0].Nama_Barang, actual[0].Nama_Barang)
// 		assert.Equal(t, resData[1].ID, actual[1].ID)
// 		assert.Equal(t, resData[1].Harga, actual[1].Harga)
// 	})
// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID   int
// 			Name string
// 		}
// 		sample := SampleUsers{
// 			ID:   4,
// 			Name: "fajar1411",
// 		}
// 		srv := New(data)
// 		_, token := helper.GenerateJWT(sample.ID)

// 		res, err := srv.MyItem(token)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "user not found")
// 		assert.Equal(t, uint(0), res) //perba
// 	})
// }

// func TestGetAllItems(t *testing.T) {
// 	data := mocks.NewItemData(t)

// 	srv := New(data)

// 	t.Run("Item succesfully", func(t *testing.T) {
// 		resData := []item.Core{
// 			{
// 				ID:          1,
// 				Nama_Barang: "Baju",
// 				Image_url:   "www.google.com",
// 				Deskripsi:   "biru asik",
// 				Stok:        1,
// 				Harga:       20000,
// 			},
// 			{
// 				ID:          2,
// 				Nama_Barang: "celana",
// 				Image_url:   "www.google.com",
// 				Deskripsi:   "biru asik",
// 				Stok:        1,
// 				Harga:       40000,
// 			},
// 		}
// 		data.On("GetAllItems").Return(resData, nil).Once()

// 		res, err := srv.GetAllItems()
// 		assert.NoError(t, err)
// 		assert.Equal(t, res, res)
// 		data.AssertExpectations(t)

// 	})
// 	t.Run("not found", func(t *testing.T) {
// 		data.On("GetAllItems").Return(nil, errors.New("Products not found")).Once()

// 		res, err := srv.GetAllItems()
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "Products not found")
// 		assert.Nil(t, res)
// 	})
// 	t.Run("Get all book error server", func(t *testing.T) {
// 		// Programming input and return repo
// 		data.On("GetAllItems").Return([]item.Core{}, errors.New("internal server error")).Once()

// 		// Program service
// 		actual, err := srv.GetAllItems()

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "internal server error")
// 		assert.Nil(t, actual)

// 	})
// }

// func TestUpdatePost(t *testing.T) {

// 	input := item.Core{Deskripsi: "biru asik"}
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

// 	t.Run("Update successfully", func(t *testing.T) {
// 		data.On("Update", 1, 1).Return(resData, nil).Once()

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		actual, err := srv.Update(token, 1, input)

// 		assert.Nil(t, err)
// 		assert.NotEqual(t, resData.Deskripsi, actual.Deskripsi)
// 		assert.Equal(t, resData.ID, actual.ID)
// 		assert.Equal(t, resData.Nama, actual.Nama)

// 		data.AssertExpectations(t)
// 	})
// 	t.Run("Update error user not found", func(t *testing.T) {

// 		token := jwt.New(jwt.SigningMethodHS256)
// 		actual, err := srv.Update(token, 1, input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "id user not found")
// 		assert.Empty(t, actual)
// 	})
// 	t.Run("Update error item not found", func(t *testing.T) {
// 		// Programming input and return repo
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
// 		data.On("Update", 1, 1, input).Return(item.Core{}, errors.New("internal server error")).Once()

// 		// Program service
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		actual, err := srv.Update(token, 1, input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "internal server error")
// 		assert.Empty(t, actual)
// 		data.AssertExpectations(t)
// 	})
// }
