package main

import (
	"encoding/json"
	"os"
)

// NewApi decodes the given json file into an Api object.
func NewApi(fileName string) (*Api, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var api Api
	err = decoder.Decode(&api)
	if err != nil {
		return nil, err
	}
	return &api, nil
}
