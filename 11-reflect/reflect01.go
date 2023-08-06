package main

import "fmt"

func main() {
	var a string // pair<statictype:string, value:"zel">
	a = "zel"

	var all_type interface{}
	// pair<type:string, value:"zel">
	all_type = a

	str, _ := all_type.(string)
	fmt.Println(str)

}
