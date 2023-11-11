package main

import (
	"flag"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
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

// func main() {
// 	file := jen.NewFile("rl")
// 	file.Var().Id("initWindow").Func().Params(jen.Uintptr()).Uintptr()
// 	file.Var().Id("initWindow").Func().Params(jen.Uintptr())
// 	fmt.Printf("%#v", file)
// }

func main() {
	api, err := NewApi(inputJson)
	if err != nil {
		panic(err)
	}

	file := jen.NewFile("rl")

	for _, function := range api.Functions {
		GenerateReturn(file.Var().Id(strcase.ToLowerCamel(function.Name)).Func().Params(GenerateParams(function.Params, true)...), function.ReturnType)
	}

	file.Func().Id("init").Params().Block(
		// Id("cs").Op(":=").Qual("C", "CString").Call(Lit("Hello from stdio\n")),
		// Qual("C", "myprint").Call(Id("cs")),
		// Qual("C", "free").Call(Qual("unsafe", "Pointer").Parens(Id("cs"))),
		// for _, function := range api.Functions {
		// 	GenerateReturn(file.Var().Id(strcase.ToLowerCamel(function.Name)).Func().Params(GenerateParams(function.Params, true)...), function.ReturnType)
		// }

		(func() []jen.Code {
			statements := make([]jen.Code, len(api.Functions))
			for i, function := range api.Functions {
				statements[i] = jen.Qual("github.com/ebitengine/purego", "RegisterLibFunc").Call(jen.Id("&"+strcase.ToLowerCamel(function.Name)), jen.Id("raylibDll"), jen.Lit(function.Name))
			}
			return statements
		}())...,

	//purego.RegisterLibFunc(&initWindow, raylibDll, "InitWindow")
	// jen.Qual("purego", "RegisterLibFunc").Call(jen.Id("&"+strcase.ToLowerCamel("InitWindow")), jen.Id("raylibDll"), jen.Lit("InitWindow")),
	)

	for _, function := range api.Functions {
		file.Commentf("%s - %s", function.Name, function.Description)
		GenerateReturn(file.Func().Id(function.Name).Params(GenerateParams(function.Params, false)...), function.ReturnType).Block()
	}

	err = file.Save(outputFile)
	if err != nil {

		f, _ := os.Create("error.log")
		f.WriteString(err.Error())
		f.Close()
		panic(err)
	}
}
