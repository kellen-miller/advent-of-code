package io

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(relativePath string) (*os.File, func()) {
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

	closeFunc := func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}

	return file, closeFunc
}

func GetScanner(relativePath string) (*bufio.Scanner, func()) {
	file, closeFile := OpenFile(relativePath)
	return bufio.NewScanner(file), closeFile
}
