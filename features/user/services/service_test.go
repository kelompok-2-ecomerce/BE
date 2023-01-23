package services

import (
	"errors"
	"projects/features/user"
	helper "projects/helper"
	"projects/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("Berhasil Login", func(t *testing.T) {
		inputEmail := "rischi@gmail.com"
		hashed, _ := helper.GeneratePassword("rischi12345")
		resData := user.Core{ID: uint(1), Nama: "rischi", Email: "rischi@gmail.com", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil)
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "rischi12345")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		inputEmail := "riau@gmail.com"

		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found"))
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "riau12345")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("password tidak sesuai", func(t *testing.T) {
		inputEmail := "rischi@gmail.com"
		hashed, _ := helper.GeneratePassword("rischi12345")
		resData := user.Core{ID: uint(1), Nama: "rischi", Email: "rischi@gmail.com", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil)
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "rischi11111")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password tidak sesuai")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}
func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Nama: "rischi", Email: "rischi@gmail.com"}
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}
func TestAllUser(t *testing.T) {
	repo := mocks.NewUserData(t)
	user := []user.Core{{
		ID:    1,
		Nama:  "Herdy",
		Email: "herdy@gmail.com",
	}}
	t.Run("Sukses lihat data", func(t *testing.T) {
		repo.On("AllUser", mock.Anything).Return(user, nil).Once()

		srv := New(repo)
		res, err := srv.AllUser()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)

	})
	t.Run("not found", func(t *testing.T) {
		repo.On("AllUser").Return(nil, errors.New("not found")).Once()

		srv := New(repo)

		res, err := srv.AllUser()
		assert.NoError(t, err)
		assert.Equal(t, res, res)
		repo.AssertExpectations(t)
	})

}
func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("sukses update data", func(t *testing.T) {
		input := user.Core{Nama: "herdi", Email: "herdiladania11@gmail.com"}
		updatedData := user.Core{ID: uint(1), Nama: "herdi", Email: "herdiladania11@gmail.com"}
		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NoError(t, err)
		assert.Equal(t, updatedData.ID, res.ID)
		assert.Equal(t, input.Nama, res.Nama)
		assert.Equal(t, input.Email, res.Email)

		repo.AssertExpectations(t)
	})
	t.Run("Input tidak sesuai format", func(t *testing.T) {
		input := user.Core{
			Email: "herdi",
		}
		repo.On("Update", uint(1), input).Return(user.Core{}, errors.New("input tidak sesuai")).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("data tidak ditemukan", func(t *testing.T) {
		input := user.Core{Nama: "herdi", Email: "herdiladania11@gmail.com"}
		repo.On("Update", uint(5), input).Return(user.Core{}, errors.New("data not found")).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		input := user.Core{Nama: "herdi", Email: "herdiladania11@gmail.com"}
		repo.On("Update", uint(1), input).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()

		service := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := service.Update(pToken, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("sukses menghapus profile", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(user.Core{}, nil).Once()

		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(token)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		_, err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(5)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(5)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		_, err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}
func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)

	srv := New(repo)

	// Case: user melakukan pendaftaran akun baru
	t.Run("Register successfully", func(t *testing.T) {
		// Prgramming input and return repo

		type SampleUsers struct {
			ID       int
			Name     string
			Email    string
			Username string
			Password string
		}
		sample := SampleUsers{
			ID:       1,
			Name:     "reza",
			Email:    "reza@gmail.com",
			Username: "reza06",
			Password: "12345",
		}
		input := user.Core{
			Nama:  sample.Name,
			Email: sample.Email,

			Password: sample.Password,
		}

		// Program service

		hashed, _ := helper.GeneratePassword(input.Password)
		resData := user.Core{
			Nama:     input.Nama,
			Email:    input.Email,
			Password: hashed,
		}
		repo.On("Register", mock.Anything).Return(resData, nil).Once()
		data, err := srv.Register(input)

		assert.Nil(t, err)
		errCompare := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(input.Password))
		assert.NoError(t, errCompare)
		assert.Equal(t, data.ID, resData.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {

		user := user.Core{
			Nama:  "Herdy",
			Email: "herdy@gmail.com",
		}

		actual, err := srv.Register(user)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation error")
		assert.Empty(t, actual)
	})

	t.Run("Register error data duplicate", func(t *testing.T) {
		type SampleUsers struct {
			ID       int
			Name     string
			Email    string
			Username string
			Password string
		}
		sample := SampleUsers{
			ID:       1,
			Name:     "reza",
			Email:    "reza@gmail.com",
			Username: "reza06",
			Password: "12345",
		}
		input := user.Core{
			Nama:  sample.Name,
			Email: sample.Email,

			Password: sample.Password,
		}

		// Programming input and return repo
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()

		// Program service
		data, err := srv.Register(input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "data sudah terdaftar")
		assert.Empty(t, data)
		repo.AssertExpectations(t)
	})
	t.Run("Masalah server", func(t *testing.T) {
		type SampleUsers struct {
			ID       int
			Name     string
			Email    string
			Username string
			Password string
		}
		sample := SampleUsers{
			ID:       1,
			Name:     "reza",
			Email:    "reza@gmail.com",
			Username: "reza06",
			Password: "12345",
		}
		input := user.Core{
			Nama:     sample.Name,
			Email:    sample.Email,
			Password: sample.Password,
		}

		// Programming input and return repo
		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()

		// Program service
		data, err := srv.Register(input)

		// Test
		assert.NotNil(t, err)
		assert.EqualError(t, err, "terdapat masalah pada server")
		assert.Empty(t, data)
		repo.AssertExpectations(t)
	})
}
