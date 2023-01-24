package handler

import "projects/features/user"

type UserReponse struct {
	Nama   string `json:"name"`
	Email  string `json:"email"`
	Foto   string `json:"photo"`
	HP     string `json:"phone_number"`
	Alamat string `json:"address"`
}

type RegisterResponse struct {
	Nama  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
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
