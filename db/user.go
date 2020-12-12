package db

import (
	"github.com/jinzhu/gorm"
)

func (r *ReaderDB) GetUser(email, password string) (*User, error) {
	var user User
	err := r.gormDB.Where("email = ? and password = ?", email, password).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &UserNotFoundErr{}
		}
		return nil, err
	}
	return &user, nil
}
