package handler

import "projects/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Alamat   string `json:"alamat" form:"alamat"`
	Foto     string
	Hp       string `json:"handphone" form:"hp"`
}

type DeleteRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email

		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Alamat = cnv.Alamat
		res.Foto = cnv.Foto
		res.HP = cnv.Hp
	case DeleteRequest:
		cnv := data.(DeleteRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
