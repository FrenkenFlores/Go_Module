package main

import (
	"fmt"
	"sort"
	"strconv"
)


func main(){
	slice := make([]int, 0)

	for {
		var input string

		fmt.Println("Add number to the slice. Enter X or x to exit: ")
		if _, err := fmt.Scan(&input); err != nil {
			panic(err)
		}

		if input == "X" || input == "x" {
			fmt.Println("Exiting the programm...")
		 	break
		}

		if num, err := strconv.Atoi(input); err == nil {
			slice = append(slice, num)
			sort.Ints(slice)
			fmt.Println(slice)
		} else {
			panic(err)
		}
	}
}