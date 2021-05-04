package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json: "name"`
	Address string `json: "address"`
	Phone string `json: "phone"`
}

func main() {
	u := User {
		Name : "fflores",
		Address : "Moscow",
		Phone : "+79******305",
	}
	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
        fmt.Println(string(out))
}
