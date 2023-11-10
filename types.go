package main

type Api struct {
	Defines   []map[string]any `json:"defines"`
	Structs   []map[string]any `json:"structs"`
	Aliases   []map[string]any `json:"aliases"`
	Enums     []map[string]any `json:"enums"`
	Callbacks []map[string]any `json:"callbacks"`
	Functions []Function       `json:"functions"`
}

type Function struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ReturnType  string      `json:"returnType"`
	Params      []Parameter `json:"params"`
}

type Parameter struct {
	Datatype string `json:"type"`
	Name     string `json:"name"`
}
