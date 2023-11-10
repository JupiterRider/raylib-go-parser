package main

import "github.com/dave/jennifer/jen"

func GenerateParams(params []Parameter) []jen.Code {
	var jenParams []jen.Code

	for _, parameter := range params {
		datatype, err := ConvertToGoType(parameter.Datatype)
		if err != nil {
			panic(err)
		}
		name := parameter.Name
		if IsReserved(name) {
			name = "_" + name
		}

		jenParams = append(jenParams, jen.Id(name).Id(datatype))

	}

	return jenParams
}

func GenerateReturn(statement *jen.Statement, ReturnType string) *jen.Statement {
	goType, err := ConvertToGoType(ReturnType)
	if err != nil {
		panic(err)
	}

	if len(goType) == 0 {
		return statement
	}

	return statement.Id(goType)
}

// IsReserved returns true if the word is a go keyword or name of a package.
func IsReserved(word string) bool {
	if jen.IsReservedWord(word) {
		return true
	}

	switch word {
	case "color":
		return true
	}

	return false
}
