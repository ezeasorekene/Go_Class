package main

import (
	"fmt"
	"sort"
	"strconv"
)

var p = fmt.Println
var s = fmt.Scanln

func main() {
	var inp string
	var num []int = make([]int, 3)
	p("Enter an int to append or press X to exit!")

	for true {
		s(&inp)

		if inp == "X" {
			break;
		}

		number, err := strconv.Atoi(inp)
		if err != nil {
			p("Kindly enter proper input")
			continue
		}
		num = append(num, number)
		sort.Ints(num[:])
		p(num)
		p("Wanna enter one more time? (X to exit)")
	}

}

