package models

import "errors"

type User struct {
	Id    uint   `gorm:"column:id; sql:"primary_key:yes; type:int unsigned NOT NULL AUTO_INCREMENT" json:"-"`
	Level int    `gorm:"column:level" sql:"type:int(11) NOT NULL" json"level"`
	Job   string `sql:"type:ENUM('Barbarian', 'Hunter', 'Mage') NOT NULL" json:"job"`
	Name  string `gorm:"column:name" json:"name"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave() (err error) {
	if !(1 <= u.Level && u.Level <= 20) {
		err = errors.New("User Level should be between 0 and 21.")
	}
	return
}
