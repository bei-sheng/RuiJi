package models

import (
	"goRJ/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

func GetUser(id int) (userModel User) {
	db.Limit(1).Where("ID", id).Find(&userModel)
	return
}

func Get(idstr string) (userModel User) {
	db.Where("id", idstr).First(&userModel)
	return
}
