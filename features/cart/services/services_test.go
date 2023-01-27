package services

import (
	"errors"
	"projects/features/cart"
	"projects/helper"
	"projects/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	repo := mocks.NewCartData(t)
	userID := 1
	productID := 1
	Qty := 100
	t.Run("success add data", func(t *testing.T) {
		resData := cart.Core{
			ID:  4,
			Qty: Qty,
		}
		repo.On("Add", userID, uint(1), Qty).Return(resData, nil).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, uint(productID), Qty)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Qty, res.Qty)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Add(token, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, cart.Core{})
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := cart.Core{}
		repo.On("Add", userID, uint(1), Qty).Return(resData, errors.New("record not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, cart.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Stok tidak cukup", func(t *testing.T) {
		resData := cart.Core{}
		repo.On("Add", userID, uint(1), Qty).Return(resData, errors.New("not enough stock")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "stok")
		assert.Equal(t, res, cart.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := cart.Core{}
		repo.On("Add", userID, uint(1), Qty).Return(resData, errors.New("query error")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, cart.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("field required wajib diisi", func(t *testing.T) {

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Add(pToken, uint(productID), 0)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "required")
		assert.Equal(t, res, cart.Core{})
		repo.AssertExpectations(t)
	})

}

func TestGetMyCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	userID := 1
	t.Run("berhasil melihat list products dari keranjang", func(t *testing.T) {
		resData := []cart.Core{
			{
				ItemID:      1,
				ProductName: "Bibit Anggur",
				ImageUrl:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Price:       12000,
				Qty:         10,
				Total:       120000,
			},
			{
				ItemID:      2,
				ProductName: "Bibit Bunga Matahari",
				ImageUrl:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674166643.jpg",
				Price:       1000,
				Qty:         100,
				Total:       100000,
			},
		}
		repo.On("GetMyCart", userID).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetMyCart(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, res[0].ID)
		assert.Equal(t, resData[0].ProductName, res[0].ProductName)
		assert.Equal(t, resData[0].ImageUrl, res[0].ImageUrl)
		assert.Equal(t, resData[0].Total, res[0].Total)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := []cart.Core{}
		repo.On("GetMyCart", userID).Return(resData, errors.New("record not found")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetMyCart(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []cart.Core{})
		repo.AssertExpectations(t)
	})
	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := []cart.Core{}
		repo.On("GetMyCart", userID).Return(resData, errors.New("query error")).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.GetMyCart(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []cart.Core{})
		repo.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		res, err := srv.GetMyCart(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []cart.Core{})
	})

}

func TestUpdateProductCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	userID := 1
	productID := 1
	Qty := 100
	t.Run("update berhasil", func(t *testing.T) {
		repo.On("UpdateProductCart", userID, uint(productID), Qty).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateProductCart(pToken, uint(productID), Qty)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("UpdateProductCart", userID, uint(productID), Qty).Return(errors.New("record not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateProductCart(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})
	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		repo.On("UpdateProductCart", userID, uint(productID), Qty).Return(errors.New("query error")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateProductCart(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		err := srv.UpdateProductCart(token, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
	})
	t.Run("Stok tidak cukup", func(t *testing.T) {
		repo.On("UpdateProductCart", userID, uint(productID), Qty).Return(errors.New("not enough stock")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateProductCart(pToken, uint(productID), Qty)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "stok")
		repo.AssertExpectations(t)
	})
}

func TestDeleteProductCart(t *testing.T) {
	repo := mocks.NewCartData(t)
	userID := 1
	productID := 1
	t.Run("Delete product berhasil", func(t *testing.T) {
		repo.On("DeleteProductCart", userID, uint(productID)).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteProductCart(pToken, uint(productID))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		err := srv.DeleteProductCart(token, uint(productID))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("DeleteProductCart", userID, uint(productID)).Return(errors.New("record not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteProductCart(pToken, uint(productID))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		repo.On("DeleteProductCart", userID, uint(productID)).Return(errors.New("query error")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteProductCart(pToken, uint(productID))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}
