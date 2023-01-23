package handler

import "projects/features/user"

type UserReponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Foto     string `json:"foto"`
	HP       string `json:"hp"`
	Alamat   string `json:"alamat"`
}
type LoginReponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		Nama:   data.Nama,
		Email:  data.Email,
		Foto:   data.Foto,
		Alamat: data.Alamat,
		HP:     data.HP,
	}
}

func ToResponses(data user.Core) RegisterResponse {
	return RegisterResponse{

		Nama:  data.Nama,
		Email: data.Email,
	}
}
func fromCoreList(dataCore []user.Core) []UserReponse {
	var dataResponse []UserReponse

	for _, v := range dataCore {
		dataResponse = append(dataResponse, ToResponse(v))
	}
	return dataResponse
}
