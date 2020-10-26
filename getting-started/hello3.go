package main

import (
    "fmt"
    // "strings"
  	"bufio"
  	"os"
  )

func main() {
  var whoareyou string
  fmt.Printf("What is your name? \n")
  reader1 := bufio.NewReader(os.Stdin)
  name, err := reader1.ReadString('\n')
  fmt.Printf("What is your age? \n")
  reader2 := bufio.NewReader(os.Stdin)
  age, err := reader2.ReadString('\n')
  fmt.Printf("Where are you from? \n")
  reader3 := bufio.NewReader(os.Stdin)
  town, err := reader3.ReadString('\n')
  whoareyou = ebube(&name,&age,&town)
  fmt.Printf(whoareyou)
}

func ebube(*name,*age,*town) string {
  var name string
  var age string
  var town string
  // age = "20"
  // name = "Ebube"
  // town = "Awka"
  word := "My name is " + name + ". I am " + age + " years old. I come from " + town + "\n"
  // fmt.Printf("My name is %s. I am %d years old. I come from %s \n", name, age, town)
  return word
}
