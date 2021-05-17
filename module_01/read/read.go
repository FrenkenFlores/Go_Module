package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func readFile (file *os.File) []name{
	reader := bufio.NewReader(file)
	var nameStruct name;
	var nameList []name;
	var tmp []string

	for {
		buffer, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err);
		}
		tmp = strings.Split(string(buffer), " ")
		if len(tmp[0]) > 20 {
			nameStruct.fname = tmp[0][:20]
		} else {
			nameStruct.fname = tmp[0][:len(tmp[0])]
		}
		if len(tmp[1]) > 20 {
			nameStruct.lname = tmp[1][:20]
		} else {
			nameStruct.lname = tmp[1][:len(tmp[1])]
		}
		nameList = append(nameList, nameStruct);
	}
	return nameList
}

func printList(nameList []name) {
	for _, nameStruct := range nameList{
		fmt.Println(nameStruct.fname)
		fmt.Println(nameStruct.lname)
	}
}

func main()  {
	var filename string
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if isMissing := os.IsNotExist(err); isMissing == true {
		log.Fatal(err)
	}
	nameList := readFile(file)
	printList(nameList)
	file.Close()
}