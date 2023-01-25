package services

import (
	"errors"
	"log"
	"mime/multipart"
	"projects/features/user"
	"projects/helper"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) AllUser() ([]user.Core, error) {
	data, err := uuc.qry.AllUser()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "terdapat masalah pada server"

		}
		return nil, errors.New(msg)

	}
	return data, nil
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {

	err := helper.Validasi(helper.ToEmailLoginString(email))
	if err != nil {
		return "", user.Core{}, err
	}

	res, err := uuc.qry.Login(email)
	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows") {
			msg = "email belum terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	//Token expires after 1 hour
	token, _ := helper.GenerateJWT(int(res.ID))

	return token, res, nil

}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {

	err := helper.Validasi(helper.ToRegister(newUser))
	if err != nil {
		return user.Core{}, err
	}
	hashed := helper.HashPassword(newUser.Password)

	newUser.Password = string(hashed)
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user tidak ditemukan harap login lagi")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user tidak ditemukan harap login lagi"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(token interface{}, updateData user.Core, file *multipart.FileHeader) (user.Core, error) {
	if updateData.Password != "" {
		hashed := helper.HashPassword(updateData.Password)
		updateData.Password = string(hashed)
	}

	id := helper.ExtractToken(token)

	if len(updateData.Email) > 0 {
		err := helper.Validasi(helper.ToEmailLogin(updateData))
		if err != nil {
			return user.Core{}, err
		}
	}

	if len(updateData.HP) > 0 {
		err := helper.Validasi(helper.ToPhoneNumber(updateData))
		if err != nil {
			return user.Core{}, err
		}
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return user.Core{}, errors.New("format input file tidak dapat dibuka")
		}
		err = helper.CheckFileSize(file.Size)
		if err != nil {
			idx := strings.Index(err.Error(), ",")
			msg := err.Error()
			return user.Core{}, errors.New("format input file size tidak diizinkan, size melebihi" + msg[idx+1:])
		}
		extension, err := helper.CheckFileExtension(file.Filename)
		if err != nil {
			return user.Core{}, errors.New("format input file type tidak diizinkan")
		}
		filename := "images/profile/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + extension

		photo, err := helper.UploadImageToS3(filename, src)
		if err != nil {
			return user.Core{}, errors.New("format input file type tidak dapat diupload")
		}

		updateData.Foto = photo

		defer src.Close()
	}

	res, err := uuc.qry.Update(uint(id), updateData)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else if strings.Contains(err.Error(), "Duplicate") {
			msg = "email sudah terdaftar"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Delete(token interface{}) (user.Core, error) {

	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user tidak ditemukan harap login lagi")
	}
	data, err := uuc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return data, nil

}
