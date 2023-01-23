package data

import (
	"errors"
	"log"
	"projects/features/item"
	"strings"

	"gorm.io/gorm"
)

type itemData struct {
	db *gorm.DB
}

func New(db *gorm.DB) item.ItemData {
	return &itemData{
		db: db,
	}
}

// Add implements item.ItemData
func (pd *itemData) Add(userID int, newItem item.Core) (item.Core, error) {
	cnv := CoreToData(newItem)
	cnv.UserID = uint(userID)

	err := pd.db.Create(&cnv).Error
	if err != nil {
		log.Println("add post query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return item.Core{}, errors.New(msg)
	}

	newItem.ID = cnv.ID

	return newItem, nil
}

// Delete implements item.ItemData
func (*itemData) Delete(userID int, itemID int) error {
	panic("unimplemented")
}

// GetAllPost implements item.ItemData
func (*itemData) GetAllPost() ([]item.Core, error) {
	panic("unimplemented")
}

// MyPost implements item.ItemData
func (*itemData) MyPost(userID int) ([]item.Core, error) {
	panic("unimplemented")
}

// Update implements item.ItemData
func (*itemData) Update(userID int, itemID int, updatedData item.Core) (item.Core, error) {
	panic("unimplemented")
}

// func (pd *postingData) Add(userID int, newPosting posting.Core) (posting.Core, error) {
//

// func (pd *postingData) Update(userID int, postID int, updatedData posting.Core) (posting.Core, error) {

// 	cnv := CoreToData(updatedData)

// 	// DB Update(value)
// 	tx := pd.db.Where("id = ? AND user_id = ?", postID, userID).Updates(&cnv)
// 	if tx.Error != nil {
// 		log.Println("update book query error :", tx.Error)
// 		return posting.Core{}, tx.Error

// 	}

// 	// Rows affected checking
// 	if tx.RowsAffected <= 0 {
// 		log.Println("update book query error : data not found")
// 		return posting.Core{}, errors.New("not found")
// 	}

// 	// return result converting cnv to book.Core
// 	return ToCore(cnv), nil
// }

// func (pd *postingData) GetAllPost() ([]posting.Core, error) {
// 	var komentar []PostUser
// 	tx := pd.db.Raw("SELECT postings.id, postings.postingan, postings.image_url, users.username FROM postings JOIN users ON users.id = postings.user_id  WHERE postings.deleted_at IS NULL").Find(&komentar)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	var dataCore = ListModelTOCore(komentar)

// 	return dataCore, nil
// }
// func (pd *postingData) Delete(userID int, postID int) error {
// 	post := Posting{}
// 	err := pd.db.Where("id = ? AND user_id = ?", postID, userID).Delete(&post, postID)
// 	if err.Error != nil {
// 		log.Println("delete book query error :", err.Error)
// 		return err.Error
// 	}
// 	if err.RowsAffected <= 0 {
// 		log.Println("delete book query error : data not found")
// 		return errors.New("not found")
// 	}

// 	return nil
// }
// func (pd *postingData) MyPost(userID int) ([]posting.Core, error) {
// 	var myBooks []PostUser
// 	err := pd.db.Raw("SELECT postings.id, postings.postingan, postings.image_url, users.username FROM postings JOIN users ON users.id = postings.user_id WHERE postings.user_id = ?", userID).Find(&myBooks).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	var dataCore = ListModelTOCore(myBooks)

// 	return dataCore, nil
// }
