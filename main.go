package main

import (
	"fmt"

	"github.com/sundargautam18/Go-with-go/pkg/date"
	"github.com/sundargautam18/Go-with-go/pkg/wifi"
)

func main() {
	fmt.Println("Hello World")
	date := date.Now()
	wifiInfo := wifi.Connect()
	fmt.Println(wifiInfo)
	fmt.Println(date)
}
