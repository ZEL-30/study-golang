package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	str_type := reflect.TypeOf(str)

	for i := 0; i < str_type.NumField(); i++ {
		tag := str_type.Field(i).Tag.Get("info")
		fmt.Println("info = ", tag)
		tag = str_type.Field(i).Tag.Get("doc")
		fmt.Println("doc = ", tag)

	}

}

func main() {

	var resume Resume

	findTag(resume)

}
