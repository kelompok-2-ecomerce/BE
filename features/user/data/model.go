package data

import (
	"projects/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string
	Foto     string
	HP       string
	Alamat   string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
		Foto:     data.Foto,
		Alamat:   data.Alamat,
		HP:       data.Alamat,
	}
}

func (dataModel *User) ModelsToCore() user.Core {
	return user.Core{
		ID:    dataModel.ID,
		Nama:  dataModel.Nama,
		Email: dataModel.Email,
	}
}
func listModelToCore(dataModel []User) []user.Core {
	var dataCore []user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func CoreToData(data user.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
		Alamat:   data.Alamat,
		HP:       data.HP,
		Foto:     data.Foto,
	}
}
