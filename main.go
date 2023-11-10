package main

import (
	"flag"
	"os"

	"github.com/dave/jennifer/jen"
)

var (
	inputJson  string
	outputFile string
)

func init() {
	flag.StringVar(&inputJson, "input", "raylib_api.json", "json file to decode")
	flag.StringVar(&outputFile, "output", "raylib_purego", "generated result file")
	flag.Parse()
}

func main() {
	api, err := NewApi(inputJson)
	if err != nil {
		panic(err)
	}

	file := jen.NewFile("rl")

	for _, function := range api.Functions {
		file.Commentf("%s - %s", function.Name, function.Description)
		GenerateReturn(file.Func().Id(function.Name).Params(GenerateParams(function.Params)...), function.ReturnType).Block()
	}

	err = file.Save(outputFile)
	if err != nil {

		f, _ := os.Create("error.log")
		f.WriteString(err.Error())
		f.Close()
		panic(err)
	}
}
