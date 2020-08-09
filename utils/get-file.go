package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetFile(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("\nFound file at %s\n", path)
	return data
}
