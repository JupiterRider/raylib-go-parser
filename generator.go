package main

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

func GenerateParams(params []Parameter, useUintptr bool) []jen.Code {
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

		if useUintptr && !IsPrimitive(datatype) {
			datatype = "uintptr"
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

func IsPrimitive(typ string) bool {
	if strings.HasPrefix(typ, "[]") {
		return true
	}
	switch typ {
	case "bool":
		return true
	case "uint8":
		return true
	case "uint16":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "int8":
		return true
	case "int16":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "float32":
		return true
	case "float64":
		return true
	case "int":
		return true
	case "uint":
		return true
	case "byte":
		return true
	case "rune":
		return true
	case "string":
		return true
	default:
		return false
	}
}
