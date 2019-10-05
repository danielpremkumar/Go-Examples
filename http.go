package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err == nil {
		// bs := make([]byte, 99999)
		// resp.Body.Read(bs)
		// fmt.Println(string(bs))
		lw := logWriter{}
		io.Copy(lw, resp.Body)
	} else {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
