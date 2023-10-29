package main

import (
	"fmt"
	"generator-go/helper"
)

func main() {

	json := helper.ReadFileJson("simple.json")
	success := helper.CreateIndex(json)
	if success {
		fmt.Println("sukses generate " + json.FolderName)
	}
}
