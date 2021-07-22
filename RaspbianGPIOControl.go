package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	var Qexp int
	var Qesp int
	fmt.Println("Please enter the parameters：")
	for true {
		fmt.Scan(&Qexp)

		switch Qexp {
		case 1:
			Qesp = 18
		case 2:
			Qesp = 23
		case 3:
			Qesp = 24
		case 4:
			Qesp = 25
		default:
			if Qexp == 0 {
				os.Exit(3)
			}
			fmt.Println("6、默认 case")
			Qesp = 12

		}

		Qtoo(Qesp)

	}
}

func Qtoo(Sun int) {
	pinsum := rpio.Pin(Sun)
	pinsum.Output()
	pinsum.Toggle()
	time.Sleep(time.Second / 20)
	pinsum.Low()
}
