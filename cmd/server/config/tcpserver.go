package config

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	v1 "github.com/nndergunov/AOCKM_Lab1/cmd/client/v1"
)

func ServerStarter() error {
	addr := fmt.Sprintf("%s:%s", v1.Host, v1.Port)

	l, err := net.Listen(v1.Type, addr)
	if err != nil {
		err = fmt.Errorf("starter: %w", err)

		return err
	}

	defer func(l net.Listener) {
		err = l.Close()
		if err != nil {
			err = fmt.Errorf("starter: %w", err)

			log.Println(err)
		}
	}(l)

	log.Printf("listening on %s", addr)

	for i := 1; i > 0; i++ {
		clientName := "Client" + strconv.Itoa(i)

		conn, err := l.Accept()
		if err != nil {
			err = fmt.Errorf("starter: %w", err)

			return err
		}

		log.Println("new client connected: ", clientName)

		go func() {
			err = connHandler(conn, clientName)
			if err != nil {
				log.Println(err)
			}
		}()
	}

	return nil
}

func connHandler(conn net.Conn, name string) error {
	for {
		bufSize := 256 // 1 header + 255 commands
		buf := make([]byte, bufSize)

		_, err := conn.Read(buf)
		if err != nil {
			err = fmt.Errorf("reading: %w", err)

			return err
		}

		log.Print(name, " -> ", string(buf))

		_, err = conn.Write([]byte("Message received."))
		if err != nil {
			err = fmt.Errorf("error sending: %w", err)

			return err
		}

		if strings.TrimSpace(string(buf)) == "STOP" {
			log.Println("Exiting TCP server!")

			return nil
		}
	}
}
