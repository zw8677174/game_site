package models

type Auth struct {
	Model
	UserName string `gorm:"column:username" json:"UserName"`
	Password string `json:"password"`
	NickName string `gorm:"column:nick_name" json:"NickName"`
}

func (u *Auth) GetOne(maps interface{}) *Auth {
	db.Where(maps).Last(u)
	return u
}
