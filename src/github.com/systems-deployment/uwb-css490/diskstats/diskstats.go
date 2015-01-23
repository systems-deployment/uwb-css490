package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func stats(response http.ResponseWriter, request *http.Request) {
	response.Header()["content-type"] = []string{"text/plain"}
	bytes, _ := exec.Command("df", "-h").Output()
	response.Write(bytes)
}

func main() {
	http.HandleFunc("/stats", stats)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("Server fail: %s\n", err)
}
