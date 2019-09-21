package main

import (
	"fmt"
)

const (
	POLLIN = 0x0001
	//POLLPRI = 0x0002
	//POLLOUT = 0x0004
	//POLLERR = 0x0008
	//POLLHUP = 0x0010
	//POLLNVA = 0x0020
)

func main() {
	// Write
	localVariable := POLLIN
	// Read
	fmt.Println("localVariable:", localVariable)
}
