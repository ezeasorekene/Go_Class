package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


/**
This function get a word/sentence as an input
and searches the word i, a and n in the word/sentence
If the word/sentence start with i, contains a and
ends with n, then the program will print "Found!" otherwise
it will print "Not Found!"
**/
func main() {
   var input string
   fmt.Printf("Enter a word/sentence that starts with i, ends with n and contains a: \n")
//   num, err := fmt.Scanf("%q", &input)
   reader := bufio.NewReader(os.Stdin)
   input, err := reader.ReadString('\n')
   if err != nil {
      fmt.Println("An error occured while reading your input", err)
      return
   }

   word := strings.ToLower(strings.TrimSuffix(input, "\n")) //Remove the space added after the suffix
   if strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") && strings.Contains(word, "a") {
      fmt.Println("Found!")
      return
   } else {
      fmt.Println("Not Found!")
      return
   }
}
