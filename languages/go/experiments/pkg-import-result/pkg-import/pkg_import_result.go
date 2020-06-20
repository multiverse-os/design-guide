package main

import (
	"fmt"

	unix "golang.org/x/sys/unix"
)

func main() {
	// Write
	localVariable := unix.POLLIN
	// Read
	fmt.Println("localVariable:", localVariable)
}
