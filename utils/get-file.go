package utils

import (
	"fmt"
	"os"
	"io/ioutil"
)

func GetFile(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	return data
}
