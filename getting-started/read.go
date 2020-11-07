package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

/***
@Author: github.com/ezeasorekene
@Title: Experimenting Structs, Slice and File Read
@Description:

This program reads information from a file and represents it in a slice of structs.
A sample text file called 'filetoread' contains a series of names. Each line of the text file has a
first name and a last name, in that order, separated by a single space on the line.

The program defines a name struct which has two fields, fname for the first name,
and lname for the last name. Each field is a string of size 20 (characters).

The program prompts the user for the name of the text file. The program will
successively read each line of the text file and create a struct which contains
the first and last names found in the file. Each struct created will be added
to a slice, and after all lines have been read from the file, the program
now has a slice containing one struct for each line in the file.
After reading all lines from the file, the program iterates through the slice
of structs and print the first and last names found in each struct.
***/

// Define necessary constant
const (
  maxstringlen = 20
)

//Error function
func iserror(e error)  {
  if e != nil {
    panic(e)
  }
}

//Check the length of the a given string
func checklen(v string) {
  if len(v) > maxstringlen {
    fmt.Printf("One of the strings in your file is greater than %d, please reduce it and try again. \n", maxstringlen)
    os.Exit(0)
  }
}

//Create a struct
type Name struct {
  fname string
  lname string
}

//Get input from a user
func UserFilePrompt() string {
  //Get input from user
  fmt.Println("Enter the absolute filename of the file (press X to cancel)")
  input := bufio.NewReader(os.Stdin)
  output, err := input.ReadString('\n')

  //Check for errors
  iserror(err)

  // Remove the trailing \n
  filename := strings.TrimSuffix(output, "\n")

  // Exit the program if the user enters 'x'
  exit := strings.ToLower(filename)
  if exit == "x" {
    fmt.Println("You ended the program. Goodbye!")
    os.Exit(0)
  }
  return filename
}

func main() {
  // Get the filename entered by the user
  filename := UserFilePrompt()

  // Open the file for manipulation
  data, err := os.Open(filename)

  // Check if there is any errors
  if err != nil {
    fmt.Printf("The file '%s' could not be found. Absolute file path name required \n", filename)
    main()
  } else { // Continue if no errors

    // The contents of the file that has ben opened
    filecontents := bufio.NewScanner(data)

    // Read the file line by line
    filecontents.Split(bufio.ScanLines)

    // Create a slice of defined Type above (Name)
    todisplay := make([]Name,0,1)

    // Read the contents of the open file line by line and append the contents to the slice
    for filecontents.Scan() {
      texts := strings.Split(filecontents.Text(), " ")
      lines := Name{fname: texts[0], lname: texts[1]}
      todisplay = append(todisplay, lines)
    }

    // Iterate through the loop and display the contents of the slice
    for _, d := range todisplay {
      checklen(d.fname)
      checklen(d.lname)
      fmt.Printf("First Name: %s, Last Name: %s\n", d.fname, d.lname)
    }
  }
}
