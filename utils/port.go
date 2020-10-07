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

func FindCommand() string {
	error, portIndex := FindStartingIndex()

	if error != nil {
		fmt.Println("ERROR: Please enter a valid port, folowed by a command to wrap")
		os.Exit(1)
	}

	if (len(os.Args) -1) <= portIndex {
		fmt.Println("ERROR: Please enter a valid port, folowed by a command to wrap")
		os.Exit(1)
	}

	return os.Args[portIndex + 1]
}

func FindArgs() []string {
	_, portIndex := FindStartingIndex()

	portIndex = portIndex + 2

	args := []string{}
	for i := range os.Args {
		if !(i < portIndex) {
			args = append(args, os.Args[i])
		}
	}

	return args
}