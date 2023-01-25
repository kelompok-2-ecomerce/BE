package services

// import (
// 	"errors"
// 	"projects/features/user"
// 	helper "projects/helper"
// 	"projects/mocks"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestLogin(t *testing.T) {
// 	repo := mocks.NewUserData(t) // mock data

// 	t.Run("Berhasil login", func(t *testing.T) {
// 		// input dan respond untuk mock data
// 		inputEmail := "fajar@gmail.com"
// 		// res dari data akan mengembalik password yang sudah di hash
// 		hashed, _ := helper.GeneratePassword("be1422")
// 		resData := user.Core{ID: uint(1), Nama: "fajar", Email: "fajar@gmail.com", Password: hashed}

// 		repo.On("Login", inputEmail).Return(resData, nil).Once() // simulasi method login pada layer data

// 		srv := New(repo)
// 		token, res, err := srv.Login(inputEmail, "be1422")
// 		assert.Nil(t, err)
// 		assert.NotEmpty(t, token)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Tidak ditemukan", func(t *testing.T) {
// 		inputEmail := "putra@alterra.id"
// 		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()

// 		srv := New(repo)
// 		token, res, err := srv.Login(inputEmail, "be1422")
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "data tidak ditemukan")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID) ///expected value nilanya 0 , actual resid//apakah expected dan actual sama?
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("server error", func(t *testing.T) {
// 		inputEmail := "fajar@gmail.com"

// 		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
// 		srv := New(repo)
// 		token, res, err := srv.Login(inputEmail, "be1422")

// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "terdapat masalah pada server")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID)
// 		// assert.NotEqual(t, uint(0), res.ID)//hasilnya harus berbeda
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Login error password doesnt match", func(t *testing.T) {
// 		inputEmail := "fajar@gmail.com"
// 		hashed, _ := helper.GeneratePassword("be1422")
// 		resData := user.Core{ID: uint(1), Nama: "fajar", Email: "fajar@gmail.com", Password: hashed}
// 		repo.On("Login", inputEmail).Return(resData, nil).Once()
// 		srv := New(repo)
// 		token, res, err := srv.Login(inputEmail, "asal")

// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "password tidak sesuai")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// }
// func TestProfile(t *testing.T) {
// 	repo := mocks.NewUserData(t)
// 	t.Run("Sukses lihat profile", func(t *testing.T) {
// 		resData := user.Core{ID: uint(1), Nama: "fajar", Email: "fajar@gmail.com"}
// 		repo.On("Profile", uint(1)).Return(resData, nil).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Profile(pToken)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(1)

// 		res, err := srv.Profile(token)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 	})
// 	t.Run("data tidak ditemukan", func(t *testing.T) {
// 		repo.On("Profile", uint(1)).Return(user.Core{}, errors.New("not found")).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Profile(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "data tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("terdapat masalah pada server", func(t *testing.T) {
// 		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Profile(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "terdapat masalah pada server")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	repo := mocks.NewUserData(t)

// 	t.Run("sukses update data", func(t *testing.T) {
// 		input := user.Core{Nama: "fajar", Email: "fajar@gmail.com"}
// 		updatedData := user.Core{ID: uint(1), Nama: "fajar", Email: "fajar@gmail.com"}
// 		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

// 		service := New(repo)
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := service.Update(pToken, input)
// 		assert.NoError(t, err)
// 		assert.Equal(t, updatedData.ID, res.ID)
// 		assert.Equal(t, input.Nama, res.Nama)
// 		assert.Equal(t, input.Email, res.Email)

// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("data tidak ditemukan", func(t *testing.T) {
// 		input := user.Core{Nama: "fajar", Email: "fajar@gmail.com"}
// 		repo.On("Update", uint(5), input).Return(user.Core{}, errors.New("data not found")).Once()

// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		srv := New(repo)
// 		inputData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		_, token := helper.GenerateJWT(1)

// 		res, err := srv.Update(token, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("Data tidak ditemukan", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:       4,
// 			Nama:     "Rizal4",
// 			Email:    "zaki@gmail.com",
// 			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:       "08123022342",
// 			Alamat:   "KOTA SURABAYA",
// 			Password: hash,
// 		}
// 		repo.On("Update", uint(4), inputData).Return(user.Core{}, errors.New("record not found")).Once()

// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		inputData.Password = password
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "data tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Input tidak sesuai format", func(t *testing.T) {
// 		input := user.Core{
// 			Email: "fajar",
// 		}
// 		repo.On("Update", uint(1), input).Return(user.Core{}, errors.New("not valid")).Once()

// 		service := New(repo)
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := service.Update(pToken, input)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format tidak sesuai")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("masalah di server", func(t *testing.T) {
// 		input := user.Core{Nama: "fajar", Email: "fajar@gmail.com"}
// 		repo.On("Update", uint(1), input).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()

// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		inputData.Password = password
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "terdapat masalah pada server")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("format email salah", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("format username salah", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("format phone number salah", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("user tidak ditemukan", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(0)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "tidak ditemukan")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// 	t.Run("email sudah terdaftar", func(t *testing.T) {
// 		inputData := user.Core{
// 			ID:       4,
// 			Nama:     "Rizal4",
// 			Email:    "zaki@gmail.com",
// 			Foto:     "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:       "08123022342",
// 			Alamat:   "KOTA SURABAYA",
// 			Password: hash,
// 		}
// 		resData := user.Core{
// 			ID:     4,
// 			Nama:   "Rizal4",
// 			Email:  "zaki@gmail.com",
// 			Foto:   "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
// 			HP:     "08123022342",
// 			Alamat: "KOTA SURABAYA",
// 		}
// 		repo.On("Update", uint(4), inputData).Return(resData, errors.New("Duplicate email or password")).Once()

// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(4)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		inputData.Password = password
// 		res, err := srv.Update(pToken, inputData, nil)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "sudah terdaftar")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// }

// func TestDelete(t *testing.T) {
// 	repo := mocks.NewUserData(t)

// 	t.Run("sukses menghapus profile", func(t *testing.T) {
// 		repo.On("Delete", uint(1)).Return(user.Core{}, nil).Once()

// 		srv := New(repo)
// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		_, err := srv.Delete(token)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("jwt tidak valid", func(t *testing.T) {
// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(1)

// 		_, err := srv.Delete(token)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "not found")
// 	})

// 	t.Run("data not found", func(t *testing.T) {
// 		repo.On("Delete", uint(5)).Return(user.Core{}, errors.New("data not found")).Once()

// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(5)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		_, err := srv.Delete(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "data tidak ditemukan")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("masalah di server", func(t *testing.T) {
// 		repo.On("Delete", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()
// 		srv := New(repo)

// 		_, token := helper.GenerateJWT(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		_, err := srv.Delete(pToken)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "internal server error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestRegister(t *testing.T) {
// 	repo := mocks.NewUserData(t)

// 	srv := New(repo)

// 	// Case: user melakukan pendaftaran akun baru
// 	t.Run("Register successfully", func(t *testing.T) {
// 		// Prgramming input and return repo

// 		type SampleUsers struct {
// 			ID       int
// 			Nama     string
// 			Email    string
// 			Password string
// 		}
// 		sample := SampleUsers{
// 			ID:       1,
// 			Nama:     "fajar",
// 			Email:    "fajar@gmail.com",
// 			Password: "12345",
// 		}
// 		input := user.Core{
// 			Nama:     sample.Nama,
// 			Email:    sample.Email,
// 			Password: sample.Password,
// 		}

// 		resData := user.Core{
// 			ID:       uint(1),
// 			Nama:     "jerry",
// 			Email:    "jerr@alterra.id",
// 			Password: hash,
// 		}
// 		inputData.Password = hash
// 		repo.On("Register", inputData).Return(resData, nil).Once()
// 		srv := New(repo)
// 		inputData.Password = password
// 		res, err := srv.Register(inputData)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		assert.Equal(t, resData.Nama, res.Nama)
// 		assert.Equal(t, resData.Email, res.Email)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Validation error", func(t *testing.T) {

// 		user := user.Core{
// 			Nama:  "fajar",
// 			Email: "fajar@gmail.com",
// 		}
// 		inputData.Password = hash
// 		repo.On("Register", inputData).Return(user.Core{}, errors.New("data is duplicated")).Once()
// 		srv := New(repo)
// 		inputData.Password = password
// 		res, err := srv.Register(inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "validation error")
// 		assert.Empty(t, actual)
// 	})

// 	t.Run("Register error data duplicate", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID       int
// 			Nama     string
// 			Email    string
// 			Username string
// 			Password string
// 		}
// 		sample := SampleUsers{
// 			ID:       1,
// 			Nama:     "fajar",
// 			Email:    "fajar@gmail.com",
// 			Password: "12345",
// 		}
// 		input := user.Core{
// 			Nama:     sample.Nama,
// 			Email:    sample.Email,
// 			Password: sample.Password,
// 		}

// 		// Programming input and return repo
// 		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()

// 		// Program service
// 		data, err := srv.Register(input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "data sudah terdaftar")
// 		assert.Empty(t, data)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Masalah server", func(t *testing.T) {
// 		type SampleUsers struct {
// 			ID       int
// 			Nama     string
// 			Email    string
// 			Username string
// 			Password string
// 		}
// 		sample := SampleUsers{
// 			ID:       1,
// 			Nama:     "fajar",
// 			Email:    "fajar@gmail.com",
// 			Password: "12345",
// 		}
// 		input := user.Core{
// 			Nama:  sample.Nama,
// 			Email: sample.Email,

// 			Password: sample.Password,
// 		}

// 		// Programming input and return repo
// 		repo.On("Register", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()

// 		// Program service
// 		data, err := srv.Register(input)

// 		// Test
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "masalah pada server")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("field required wajib diisi", func(t *testing.T) {
// 		inputData := user.Core{
// 			Nama:     "jerry",
// 			Email:    "jerr@alterra.id",
// 			Password: password,
// 		}
// 		srv := New(repo)
// 		inputData.Password = password
// 		res, err := srv.Register(inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "wajib diisi")
// 		assert.Equal(t, uint(0), res.ID)
// 	})
// 	t.Run("format email salah", func(t *testing.T) {
// 		inputData := user.Core{
// 			Nama:     "jerry",
// 			Email:    "jerr@alterra.id",
// 			Password: password,
// 		}
// 		srv := New(repo)
// 		inputData.Password = password
// 		res, err := srv.Register(inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format")
// 		assert.Equal(t, uint(0), res.ID)
// 	})
// 	t.Run("format username salah", func(t *testing.T) {
// 		inputData := user.Core{
// 			Nama:     "jerry",
// 			Email:    "jerr@alterra.id",
// 			Password: password,
// 		}
// 		srv := New(repo)
// 		inputData.Password = password
// 		res, err := srv.Register(inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "format")
// 		assert.Equal(t, uint(0), res.ID)
// 	})

// }
