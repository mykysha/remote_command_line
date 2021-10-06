package main

import (
	"log"
	"os"

	v1 "github.com/nndergunov/AOCKM_Lab1/cmd/client/v1"
	"github.com/nndergunov/AOCKM_Lab1/cmd/info"
)

func main() {
	Starter()
}

func Starter() {
	logger := log.New(os.Stdout, "client ", log.LstdFlags)
	c := v1.Client{
		Host: info.Host,
		Port: info.Port,
		Type: info.Type,
		Log:  logger,
	}

	c.Init()
}
