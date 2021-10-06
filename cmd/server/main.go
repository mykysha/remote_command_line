package main

import (
	"log"

	"github.com/nndergunov/AOCKM_Lab1/cmd/server/config"
)

// main is Server app entry point.
func main() {
	err := config.ServerStarter()
	if err != nil {
		log.Fatalln(err)
	}
}
