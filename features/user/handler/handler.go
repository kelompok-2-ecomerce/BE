package handler

import (
	"errors"
	"fmt"
	"net/http"
	"projects/features/user"
	helper "projects/helper"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

// func (uc *userControll) AllUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		result, err := uc.srv.AllUser()
// 		if err != nil {
// 			return c.JSON((helper.PrintErrorResponse(err.Error())))
// 		}

// 		dataResp := fromCoreList(result)
// 		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan data user", dataResp))
// 	}
// }

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponsesLogin(res, token)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil login", dataResp))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponses(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", dataResp))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponse(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", dataResp))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		ex := c.Get("user")

		input := UpdateRequest{}
		file, errPath := c.FormFile("foto")

		fmt.Print("error get path handler, err = ", errPath)

		if file != nil {
			res, err := helper.UploadImage(c)
			// fmt.Println(res)
			if err != nil {
				fmt.Println(err)
				return errors.New("create gambar failed cannot upload data")
			}
			input.Foto = res
			// fmt.Println(input.Image_url)
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		dataCore := *ToCore(input)

		res, err := uc.srv.Update(ex, dataCore)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponse(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mengubah data", dataResp))
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("user")
		res, err := uc.srv.Delete(tx)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := ToResponse(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil hapus", result))
	}
}
