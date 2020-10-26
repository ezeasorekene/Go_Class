package main

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
        var sentence string
        if name != "" {
                sentence = "One for " + name + ", one for me."
        } else {
                sentence = "One for you, one for me."
        }
        return sentence
}

func main() {
	fmt.Println(ShareWith("Ekene"))
}
