package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	f, err := os.Open("list.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		anle := (scanner.Text())
		anle = strings.Replace(anle, "----", "", -1)
		color.Red(anle)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
