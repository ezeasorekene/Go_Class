package main

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
	"os"
)
/*
* @Author: github.com/ezeasorekene
* @Title: Experimenting Go Slice and Go Sort
* @Description
* This program prompts the user to enter integers
* and stores the integers in a sorted slice. The program
* is written as a loop. Before entering the loop,
* the program creates an empty integer slice of size (length) 3.
* During each pass through the loop, the program prompts
* the user to enter an integer to be added to the slice.
* The program then adds the integer to the slice, sorts the slice,
* and prints the contents of the slice in a sorted order.
* The slice grows in size to accommodate any number of integers
* which the user decides to enter. The program only quits
* (exits the loop) when the user enters the character ‘x’ instead of an integer.
*/
func main() {
	var input string
	integer := make([]int,3)
	fmt.Printf("Enter an integer: \n")
	for true {
		_, err := fmt.Scan(&input)
		exit := strings.ToLower(input)
		if err != nil {
			fmt.Println("An error occured while reading your input. Try again.")
			os.Exit(0)
		}
		if exit == "x" {
			fmt.Println("You ended the program!")
			os.Exit(0)
		}

		store, err := strconv.Atoi(input)
		integer = append(integer, store)
		sort.Ints(integer)
		fmt.Println(integer)
		fmt.Printf("Enter another integer: \n")
	}
}
