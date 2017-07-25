package models

import "errors"

type Item struct {
	Id           uint   `gorm:"column:id" sql:"primary_key:yes; type:int unsigned NOT NULL AUTO_INCREMENT" json:"-"`
	Level        int    `gorm:"column:level" sql:"type:int(11) NOT NULL" json"level"`
	Strength     int    `gorm:"column:strength" sql:"type:int(11) NOT NULL" json"strength"`
	Dexterity    int    `gorm:"column:dexterity" sql:"type:int(11) NOT NULL" json"dexterity"`
	Intelligence int    `gorm:"column:intelligence" sql:"type:int(11) NOT NULL" json"intelligence"`
	Vitality     int    `gorm:"column:vitality" sql:"type:int(11) NOT NULL" json"vitality"`
	Name         string `gorm:"column:name" sql:"type:varchar(45)" json:"name"`
}

func (i *Item) TableName() string {
	return "items"
}

func (i *Item) BeforeSave() (err error) {
	if !(1 <= i.Level && i.Level <= 20) {
		err = errors.New("Item Level should be between 0 and 21.")
	}
	return
}
