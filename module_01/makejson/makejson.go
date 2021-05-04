package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var name string
	var address string
	for {
		fmt.Println("Please enter a name:")
		fmt.Scan(&name)
		fmt.Println("please enter an address:")
		fmt.Scan(&address)
		m := make(map[string]string)
		m[name] = address
		if out, err := json.Marshal(m); err == nil {
				fmt.Println(string(out))
		} else {
			panic(err)
		} 
	}
}

