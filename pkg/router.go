package cline

import (
	"strings"
)

// Route arranges commands with their handler functions.
func Route(req string) []byte {
	var info []byte

	request := strings.Split(req, " ")

	request[0] = strings.ToLower(request[0])

	switch request[0] {
	case "date":
		info = date()

	case "echo":
		info = echo(request[1:])

	case "timeout":
		info = timeout(request[1:])

	case "help", "h":
		info = help()

	case "who":
		info = who()

	default:
		info = []byte("unknown command")
	}

	return info
}
