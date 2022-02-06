package repository

import (
	"encoding/json"
	"io/ioutil"
)

const fileName string = "docs/volumen_list.json" //Route and name of the file from directory 'restapiGo/'

/**
GetFile_ is a function responsible for extracting the data from
the JSON file and delivering an interface{}
*/

func GetFile_() interface{} {

	// Reading the file

	fileInByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		print("Error wanting to read the file")
	}

	// Convert []byte to interface{}

	var interfaceFile interface{}
	err = json.Unmarshal(fileInByte, &interfaceFile)
	if err != nil {
		print("Structure file incorrect.")
	} else {
		return interfaceFile
	}

	return nil
}
