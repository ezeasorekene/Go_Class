package main

import (
	"fmt"
  "encoding/json"
  "bufio"
  "strings"
	"os"
)
/*
* @Author: github.com/ezeasorekene
* @Title: Experimenting Go Map and JSON Marshal()
* @Description
* This program prompts the user to first enter a name, and then enter an address.
* The program creates a map and adds the name and address to the map using the keys
* “name” and “address”, respectively. The program uses Marshal() to create a JSON object
* from the map, and then prints the JSON object.
*/

// type BioData struct {
//   name string
//   address string
// }

func main() {
	var yourname string
  var youraddress string

  //Initialize the input reader
  reader1 := bufio.NewReader(os.Stdin)
  //Request for an inout - name
  fmt.Printf("Enter your name: ")
  //Read input
  name, err := reader1.ReadString('\n')
  //Remove trailing \n
  yourname = strings.TrimSuffix(name, "\n")
  //Check if there is any errors
  if err != nil {
    fmt.Println("An error occured while reading your name. Try again.")
    os.Exit(0)
  }

  //Initialize the second input reader
  reader2 := bufio.NewReader(os.Stdin)
  //Request for an inout - address
  fmt.Printf("Enter your address: ")
  //Read input
  address, err := reader2.ReadString('\n')
  //Remove trailing \n
  youraddress = strings.TrimSuffix(address, "\n")
  //Check if there is any errors
  if err != nil {
    fmt.Println("An error occured while reading your address. Try again.")
    os.Exit(0)
  }

  bio := make(map[string]string)
  bio["name"] = yourname
  bio["address"] = youraddress

  biodata, _ := json.Marshal(bio)

  fmt.Println(string(biodata))
}
