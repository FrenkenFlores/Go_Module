package main
import "fmt"

func printSlice(sli []int) {
	fmt.Printf("Slice order: ")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v ", sli[i])
	}
	fmt.Printf("\n")
}

func sortSlice(sli []int) {
	for i := 0; i < len(sli) - 1; i++ {
		for j := 0; j < len(sli) - 1; j++ {
			if sli[j] > sli[j + 1] {
				tmp := sli[j + 1]
				sli[j + 1] = sli[j]
				sli[j] = tmp
			}
		}
	}
}

func main () {
	var i int = 0
	var sli []int = make([]int, 3, 3)
	for {
		fmt.Println("Add number to the slice: ")
		fmt.Scan(&sli[i])
		i = (i + 1) % 3
		if i == 0 {
			sortSlice(sli)
			printSlice(sli)
		}
	}
}
