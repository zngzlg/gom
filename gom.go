package gom

import (
	"fmt"
	"github.com/zngzlg/gom/testgom"
)

func Hello() bool {
	fmt.Println("hello")
	fmt.Println("world")
        testgom.TestGom()
	return true
}
