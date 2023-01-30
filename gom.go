package gom

import (
	"fmt"
	"github.com/zngzlg/gom/test"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func Hello() bool {
	fmt.Println("hello world")
	test.TestGom()
	return true
}

type UserInfo struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"index"`
}

type UserInfoExp struct {
	UserInfo UserInfo  `gorm:"embedded"`
	Create   time.Time `gorm:"autoCreateTime"`
}

func Create() bool {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("error when init db")
	}
	// db pool config
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)

	userInfo := UserInfo{Name: "zzl"}
	userInfoExp := UserInfoExp{UserInfo: userInfo}

	// sync schema first
	db.AutoMigrate(&UserInfo{})
	db.Create(&userInfo)
	db.AutoMigrate(&UserInfoExp{})
	db.Create(&userInfoExp)
	return true
}
