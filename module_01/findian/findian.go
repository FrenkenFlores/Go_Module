package main

import (
	"fmt"
	"strings"
)
func main() {
	var s string
	fmt.Println("Enter a string")
	fmt.Scanln(&s)
	if strings.Compare(string(s[0]), "i") == 0 &&
		strings.Contains(s, "a") == true &&
		strings.Compare(string(s[len(s) - 1]), "n") == 0{
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
