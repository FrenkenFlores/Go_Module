package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func BubbleSort(numberArray []int) {
	for i := 0; i < len(numberArray) - 1; i++ {
		for j := 0; j < len(numberArray) - 1 - i; j++ {
			if numberArray[j] > numberArray[j + 1] {
				swap(&numberArray[j], &numberArray[j + 1])
			}
		}
	}
}


func ParseLine() []int {
	var numberArray []int
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var tmp_slice []string = strings.Split(scanner.Text(), " ")
	var input_slice []string
	for j := 0; j < len(tmp_slice); j++{			// get rid of spaces
		if len(tmp_slice[j]) == 0{
			continue
		}
		input_slice = append(input_slice, tmp_slice[j])
	}
	if len(input_slice) > 10 {
		log.Fatal("type in a sequence of up to 10 integers")
	}
	for i := 0; i < len(input_slice); i++ {						//convert string numbers to int
		n, _ := strconv.Atoi(input_slice[i])
		numberArray = append(numberArray, n)
	}
	return numberArray
}

func main () {
	var numberArray []int = ParseLine();
	BubbleSort(numberArray)
	fmt.Println(numberArray)
}