package io

import (
	"fmt"
	"os"
)

func OpenInput(relativePath string) (*os.File, func(file *os.File)) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fullPath := pwd + "/" + relativePath
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}

	closeFunc := func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}

	return file, closeFunc
}
