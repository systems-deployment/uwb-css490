package main

import (
	"fmt"
	"github.com/systems-deployment/uwb-css490/diskmon2/manager"
	"os"
)

var (
	disks = []string{"/dev/disk0s2"}
)

func main() {
	if err := manager.Monitor(disks); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
