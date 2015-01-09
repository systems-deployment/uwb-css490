// Copyright 2015 Morris Bernstein
package main

import (
	"fmt"
	"github.com/systems-deployment/uwb-css490/lib/topten"
	"net/http"
	"os"
	"path/filepath"
)

func handleTopten(response http.ResponseWriter, request *http.Request) {
	response.Header()["content-type"] = []string{"text/plain"}

	// NOTE: THIS IS UNSAFE.  YOU GENERALLY DO NOT WANT TO PROVIDE
	// OPEN ACCESS TO YOUR FILESYSTEM TO ANYONE WHO ASKS FOR IT
	request.ParseForm()
	path := request.URL.Path
	_, filename := filepath.Split(path)
	f, err := os.Open(filename)
	if err != nil {
		http.Error(response, fmt.Sprintf("file \"%s\" not found: %s", filename, err), http.StatusNotFound)
		return
	}

	topten.TopTen(f, response)
}

func main() {
	http.HandleFunc("/", handleTopten)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("Server fail: %s\n", err)
	os.Exit(1)
}
