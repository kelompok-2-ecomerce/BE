package data

import (
	"errors"
	"log"
	"projects/features/user"
	"strings"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	res := User{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cnv := CoreToData(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}
func (uq *userQuery) AllUser() ([]user.Core, error) {
	var user []User

	tx := uq.db.Raw("SELECT users.id, users.name, users.email, users.username  From users WHERE users.deleted_at IS NULL").Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = listModelToCore(user)
	return dataCore, nil
}

func (uq *userQuery) Update(id uint, updateData user.Core) (user.Core, error) {
	userModel := CoreToData(updateData)
	userModel.ID = id

	Input := uq.db.Model(&userModel).Where("id = ?", id).Updates(userModel)

	if Input.Error != nil {
		log.Println("Get By ID query error", Input.Error.Error())
		return user.Core{}, Input.Error
	}
	if Input.RowsAffected <= 0 {
		return user.Core{}, errors.New("not found")
	}

	return ToCore(userModel), nil
}

func (uq *userQuery) Delete(id uint) (user.Core, error) {
	users := User{}

	delete := uq.db.Delete(&users, id)

	if delete.Error != nil {
		log.Println("Get By ID query error", delete.Error.Error())
		return user.Core{}, delete.Error
	}

	if delete.RowsAffected < 0 {
		log.Println("Rows affected delete error")
		return user.Core{}, errors.New("user not found")
	}

	return ToCore(users), nil
}
