package main

import (
	"PortScanner/port"
	"fmt"
)

func main() {
	fmt.Println("PortScanner in Go")

	open := port.ScanPort("tcp", "localhost", 5500)
	fmt.Printf("Port Open: %t\n", open)

}

.