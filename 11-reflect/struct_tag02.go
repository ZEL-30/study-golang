package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {

	movie01 := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	// 编码的过程 结构体 --> json
	json_str, err := json.Marshal(movie01)
	if err != nil {
		fmt.Println("encoding json string failed, ", err)
	} else {
		fmt.Printf("%s\n", json_str)
	}

	// 编码的过程 结构体 --> json
	// json_str = "{"title":"喜剧之王","year":2000,"price":10,"actors":["周星驰","张柏芝"]}"
	movie02 := Movie{}
	err = json.Unmarshal(json_str, &movie02)
	if err != nil {
		fmt.Println("encoding struct failed, ", err)
	} else {
		fmt.Println(movie02)
	}

}
