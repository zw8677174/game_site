package models

import "github.com/jinzhu/gorm"

type Auth struct {
	ID       int64    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (int64, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if auth.ID > 0 {
		return auth.ID, nil
	}

	return 0, nil
}
