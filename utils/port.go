package utils

import (
	"fmt"
	"net"
	"os"
)

func FindPortOrFail() string {
	error, portIndex := FindStartingIndex()

	if error != nil {
		fmt.Println("ERROR: Please enter a valid port, folowed by a command to wrap")
		os.Exit(1)
	}

	if !ValidatePort(os.Args[portIndex]) {
		fmt.Println("ERROR: Please enter a valid port, folowed by a command to wrap")
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":" + os.Args[portIndex])

	if err != nil {
		fmt.Println("ERROR: Can't bind to port ", os.Args[portIndex])
		os.Exit(1)
	}

	_ = ln.Close()
	return os.Args[portIndex]
}
