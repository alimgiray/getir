package memory

import "gorm.io/gorm"

type Record struct {
	gorm.Model `json:"-"`
	Key        string `json:"key" gorm:"primary_key"`
	Value      string `json:"value"`
}
