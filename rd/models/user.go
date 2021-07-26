package models

type User struct {
	Model
	Uid    int64 `json:""`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Get(maps interface{}) *User {
	db.Where(maps).Last(u)
	return u
}
