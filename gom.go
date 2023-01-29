package gom

import (
	"fmt"
	"time"
	"github.com/zngzlg/gom/test"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Hello() bool {
	fmt.Println("hello world")
	test.TestGom()
	return true
}

type UserInfo struct {
	ID		uint		`gorm:"PrimaryKey"`
	Name	string		`gorm:"index"`
}

type UserInfoExp struct {
	UserInfo UserInfo	`gorm:"embedded"`
	Create time.Time 	`gorm:"autoCreateTime"`
}

func Create() bool {
  	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("error when init db")
	}
	userInfo := UserInfo{Name:"zzl"}
	userInfoExp := UserInfoExp{UserInfo:userInfo}

	db.Create(&userInfo)
	return true
}
