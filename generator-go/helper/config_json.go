package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ConfigJson struct {
	FolderName string   `json:"folder_name,omitempty"`
	Endpoint   string   `json:"endpoint,omitempty"`
	Entity     []Entity `json:"entity"`
}

type Entity struct {
	Label     string `json:"label"`
	DataType  string `json:"data_type,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

func ReadFileJson(namefile string) ConfigJson {
	jsonFile, err := os.Open(namefile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	var configJson ConfigJson
	json.Unmarshal(byteValue, &configJson)
	return configJson
}
