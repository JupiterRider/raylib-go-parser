package main

import (
	"flag"
	"fmt"
)

var (
	inputJson string
)

func init() {
	flag.StringVar(&inputJson, "input", "raylib_api.json", "json file to decode")
	flag.Parse()
}

func main() {
	api, err := NewApi(inputJson)
	if err != nil {
		panic(err)
	}

	for _, v := range api.Functions {
		fmt.Println(v.ReturnType)
	}
}
