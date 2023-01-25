package helper

import (
	"errors"
	"log"
	"projects/features/cart"
	"projects/features/item"
	"projects/features/user"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterValidate struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string
}

type PhoneNumberValidate struct {
	PhoneNumber string `validate:"numeric"`
}

type LoginEmailValidate struct {
	Email string `validate:"required,email"`
}

type ItemNameValidate struct {
	Nama_Barang string `validate:"required"`
}
type ItemStokValidate struct {
	Stok int `validate:"required,numeric"`
}
type ItemHargaValidate struct {
	Harga float64 `validate:"required"`
}

type QtyValidate struct {
	Qty int `validate:"required,numeric"`
}

func ToQtyInt(data int) QtyValidate {
	return QtyValidate{
		Qty: data,
	}
}
func ToQty(data cart.Core) QtyValidate {
	return QtyValidate{
		Qty: data.Qty,
	}
}

func ToRegister(data user.Core) RegisterValidate {
	return RegisterValidate{
		Name:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}
}

func ToItemHarga(data item.Core) ItemHargaValidate {
	return ItemHargaValidate{
		Harga: data.Harga,
	}
}
func ToItemStok(data item.Core) ItemStokValidate {
	return ItemStokValidate{
		Stok: data.Stok,
	}
}
func ToItemName(data item.Core) ItemNameValidate {
	return ItemNameValidate{
		Nama_Barang: data.Nama_Barang,
	}
}
func ToPhoneNumber(data user.Core) PhoneNumberValidate {
	return PhoneNumberValidate{
		PhoneNumber: data.HP,
	}
}

func ToEmailLogin(data user.Core) LoginEmailValidate {
	return LoginEmailValidate{
		Email: data.Email,
	}
}

func ToEmailLoginString(data string) LoginEmailValidate {
	return LoginEmailValidate{
		Email: data,
	}
}

func Validasi(data interface{}) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		msg := ""
		if strings.Contains(err.Error(), "required") {
			msg = "field required wajib diisi"
		} else if strings.Contains(err.Error(), "email") {
			msg = "format email salah"
		} else if strings.Contains(err.Error(), "Username") {
			msg = "format username salah"
		} else if strings.Contains(err.Error(), "PhoneNumber") {
			msg = "format phone_number salah"
		}
		return errors.New(msg)
	}
	return nil
}
