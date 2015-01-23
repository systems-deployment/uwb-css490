// diskstats is a simple demo web server that takes content from the
// output of an external command.
//
// Copyright 2015 Systems Deployment, LLC
// Author: Morris Bernstein (morris@systems-deployment.com)
package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

const (
	timeout = 2 * time.Second
)

func stats(response http.ResponseWriter, request *http.Request) {
	response.Header()["content-type"] = []string{"text/plain"}

	var (
		done    = make(chan error)
		bytes   []byte
		err     error
		command = exec.Command("df", "-h")
	)

	// Execute the command in a separate goroutine so the main
	// thread can timeout gracefully.
	go func() {
		var err error
		bytes, err = command.Output()
		done <- err
	}()

	// Start the clock.
	timer := time.NewTimer(timeout)

	// Wait for the command to complete or time out.
	select {
	case err = <-done:
		// The goroutine completed or failed.  Stop the timer
		// so its goroutine can terminate.
		timer.Stop()
	case <-timer.C:
		// Too late.  Kill the process so the goroutine can
		// complete.  Otherwise, we'd have a potential leak.
		if err := command.Process.Kill(); err == nil {
			// Catch the done message without blocking, so
			// it won't leak.
			go func() { <-done }()
		}
		err = fmt.Errorf("timeout")
	}

	// Did it work?
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Write(bytes)
}

func main() {
	http.HandleFunc("/stats", stats)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("Server fail: %s\n", err)
}
