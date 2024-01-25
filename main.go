package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	fmt.Println("starting server")
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	if bytesWritten, err := io.WriteString(w, "This is my website!\n"); err != nil {
		fmt.Printf("bytesWritten %d, err %s", bytesWritten, err.Error())
	}
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	if bytesWritten, err := io.WriteString(w, "Hello, There!\n"); err != nil {
		fmt.Printf("bytesWritten %d, err %s", bytesWritten, err.Error())

	}
}
