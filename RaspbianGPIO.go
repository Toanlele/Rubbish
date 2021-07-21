package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	GPIOsun := [8]int{18, 23, 24, 25, 12, 16, 20, 21}

	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	var i, k int
	for k = 0; k < 100; k++ {
		for i = 0; i < 8; i++ {
			fmt.Println("balanc = ", GPIOsun[i])
			dome1 := GPIOsun[i]
			gpio(dome1)
		}
	}
}
func gpio(sum int) {
	pinsum := rpio.Pin(sum)
	pinsum.Output()
	pinsum.Toggle()
	time.Sleep(time.Second / 10)
	pinsum.Low()

}
