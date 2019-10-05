package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//command line arguments
	args := os.Args
	fileName := args[1]
	//data, err := ioutil.ReadFile(fileName)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	//fmt.Println(string(data))
	io.Copy(os.Stdout, f)
}
