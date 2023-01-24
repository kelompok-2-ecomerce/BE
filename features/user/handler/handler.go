package handler

import (
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

func (uc *userControll) AllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := uc.srv.AllUser()
		if err != nil {
			return c.JSON((helper.PrintErrorResponse(err.Error())))
		}

		dataResp := fromCoreList(result)
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil menampilkan data user", dataResp))
	}
}

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
		dataResp := ToResponses(res)
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("success login", dataResp, token))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		_, err := uc.srv.Register(*ToCore(input))
		// res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		// dataResp := ToResponses(res)
		// return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", dataResp))
		return c.JSON(http.StatusCreated, helper.PrintSuccessReponse("success add data"))
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
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Berhasil menampilkan profil", dataResp))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		ex := c.Get("user")

		input := UpdateRequest{}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		dataCore := *ToCore(input)

		_, err = uc.srv.Update(ex, dataCore, file)
		// res, err := uc.srv.Update(ex, dataCore, file)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		// dataResp := ToResponse(res)
		// return c.JSON(http.StatusOK, helper.PrintSuccessReponse("berhasil mengubah data", dataResp))
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("update berhasil"))
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("user")
		_, err := uc.srv.Delete(tx)
		// res, err := uc.srv.Delete(tx)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		// result := ToResponse(res)
		// return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Delete user berhasil", result))
		return c.JSON(http.StatusOK, helper.PrintSuccessReponse("Delete user berhasil"))
	}
}
