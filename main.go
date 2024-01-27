package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KamenDimitrov97/final-project-telerik/config"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	cfg, err := config.Get()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("starting server: ")

	svcError := http.ListenAndServe(cfg.BindAddr, nil)
	if errors.Is(svcError, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", svcError)
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
