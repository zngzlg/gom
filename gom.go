package gom

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
	"net/http"
	"log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Hello() bool {
	fmt.Println("hello world")
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

func Metrics() {
	log.Printf("start prometheus metrics")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
