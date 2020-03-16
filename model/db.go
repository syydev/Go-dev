package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/SeungyoungYang/Go-dev/utils"
)

var db *database

type database struct {
	local *gorm.DB
}

func Init(c utils.Config) {
	config := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Addr,
		c.Database.DbName,
		true,
		"Local")
	d, err := gorm.Open("mysql", config)
	if err != nil {
		panic(err)
	}
	d.DB().SetMaxIdleConns(1)
	d.LogMode(true)
	db = &database{
		local: d,
	}
}

func Close() {
	db.local.Close()
}

func AutoMigrate() {
	db.local.AutoMigrate(
		User{},
	)
}
