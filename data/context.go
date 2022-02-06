package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Context *gorm.DB

func init() {
	var err error
	Context, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
