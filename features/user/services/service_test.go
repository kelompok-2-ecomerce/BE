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
)

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t) // mock data
	password := "be1422"
	t.Run("Berhasil login", func(t *testing.T) {
		// input dan respond untuk mock data
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		// res dari data akan mengembalik password yang sudah di hash
		hashed := helper.HashPassword(password)
		resData := user.Core{ID: uint(1), Password: hashed}

		repo.On("Login", inputData).Return(resData, nil).Once() // simulasi method login pada layer data

		srv := New(repo)
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("email belum terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		repo.On("Login", inputData).Return(user.Core{}, errors.New("record not found")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "belum terdaftar")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("token belum terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "belum terdaftar")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("Masalah pada server", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		repo.On("Login", inputData).Return(user.Core{}, errors.New("login query error :")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "masalah pada server")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Salah password", func(t *testing.T) {
		// input dan respond untuk mock data
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		// res dari data akan mengembalik password yang sudah di hash
		hashed := helper.HashPassword("asdasdasdad")
		resData := user.Core{ID: uint(1), Password: hashed}

		repo.On("Login", inputData).Return(resData, nil).Once() // simulasi method login pada layer data

		srv := New(repo)
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password tidak sesuai")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerralterra.id",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		token, res, err := srv.Login(inputData.Email, inputData.Password)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
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
	password := "be1422"
	hash := helper.HashPassword(password)
	t.Run("Berhasil update user tanpa image", func(t *testing.T) {
		inputData := user.Core{
			ID:       4,
			Nama:     "Rizal4",
			Email:    "zaki@gmail.com",
			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:       "08123022342",
			Alamat:   "KOTA SURABAYA",
			Password: hash,
		}
		resData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		repo.On("Update", uint(4), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		inputData.Password = password
		res, err := srv.Update(pToken, inputData, nil)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama, res.Nama)
		assert.Equal(t, resData.Email, res.Email)
		assert.Equal(t, resData.HP, res.HP)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Update(token, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		inputData := user.Core{
			ID:       4,
			Nama:     "Rizal4",
			Email:    "zaki@gmail.com",
			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:       "08123022342",
			Alamat:   "KOTA SURABAYA",
			Password: hash,
		}
		repo.On("Update", uint(4), inputData).Return(user.Core{}, errors.New("record not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		inputData.Password = password
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := user.Core{
			ID:       4,
			Nama:     "Rizal4",
			Email:    "zaki@gmail.com",
			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:       "08123022342",
			Alamat:   "KOTA SURABAYA",
			Password: hash,
		}
		repo.On("Update", uint(4), inputData).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		inputData.Password = password
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("format username salah", func(t *testing.T) {
		inputData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("format phone number salah", func(t *testing.T) {
		inputData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("user tidak ditemukan", func(t *testing.T) {
		inputData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("email sudah terdaftar", func(t *testing.T) {
		inputData := user.Core{
			ID:       4,
			Nama:     "Rizal4",
			Email:    "zaki@gmail.com",
			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:       "08123022342",
			Alamat:   "KOTA SURABAYA",
			Password: hash,
		}
		resData := user.Core{
			ID:     4,
			Nama:   "Rizal4",
			Email:  "zaki@gmail.com",
			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			HP:     "08123022342",
			Alamat: "KOTA SURABAYA",
		}
		repo.On("Update", uint(4), inputData).Return(resData, errors.New("Duplicate email or password")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		inputData.Password = password
		res, err := srv.Update(pToken, inputData, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
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
	password := "be1422"
	hash := helper.HashPassword(password)
	t.Run("Berhasil registrasi", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}

		resData := user.Core{
			ID:       uint(1),
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: hash,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(resData, nil).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama, res.Nama)
		assert.Equal(t, resData.Email, res.Email)
		repo.AssertExpectations(t)
	})

	t.Run("email/username sudah terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(user.Core{}, errors.New("data is duplicated")).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah pada server", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "masalah pada server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("field required wajib diisi", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wajib diisi")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format username salah", func(t *testing.T) {
		inputData := user.Core{
			Nama:     "jerry",
			Email:    "jerr@alterra.id",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

}
