package models

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_time"`
	ModifiedBy string `json:"modified_time"`
	AuthorId  string `json:"author_id"`
}


func GetList(pageNum int, pageSize int, maps interface{}) ([]Game, error) {
	var (
		games []Game
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&games).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&games).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return games, nil
}