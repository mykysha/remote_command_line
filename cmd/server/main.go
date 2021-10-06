package main

import (
	"github.com/nndergunov/AOCKM_Lab1/cmd/info"
	"github.com/nndergunov/AOCKM_Lab1/cmd/server/config"
)

// main is Server app entry point.
func main() {
	Starter()
}

func Starter() {
	srv := config.Server{
		Host: info.Host,
		Port: info.Port,
		Type: info.Type,
	}

	srv.Init()
}
