package main

import (
	"fmt"
)

func output(num int) {
	if num < 2 {
		return
	}

	output(num / 2)
	fmt.Println(num)

}

func main() {
	input := 9
	output(input)
}
