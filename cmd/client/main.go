package main

import (
	"log"
	"os"

	v1 "github.com/nndergunov/AOCKM_Lab1/cmd/client/v1"
	"github.com/nndergunov/AOCKM_Lab1/cmd/info"
)

// main is client entry point.
func main() {
	Starter()
}

// Starter gives client all the needed information.
func Starter() {
	logFile, err := os.Create("log/client/clientLog.txt")
	if err != nil {
		log.Println(err)

		return
	}

	logger := log.New(logFile, "client ", log.LstdFlags)
	c := v1.Client{
		Host: info.Host,
		Port: info.Port,
		Type: info.Type,
		Log:  logger,
	}

	c.Init()
}
