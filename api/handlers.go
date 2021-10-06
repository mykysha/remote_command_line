package api

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type API struct {
	Log      *log.Logger
	Listener net.Listener
}

func (a API) Init() {
	a.connecter()
}

func (a API) connecter() {
	for i := 1; i > 0; i++ {
		clientName := "Client" + strconv.Itoa(i)

		conn, err := a.Listener.Accept()
		if err != nil {
			a.Log.Println(err)

			return
		}

		a.Log.Println("new client connected: ", clientName)

		go func() {
			err = a.handler(conn, clientName)
			if err != nil {
				a.Log.Println(err)
			}
		}()
	}
}

func (a API) handler(conn net.Conn, name string) error {
	defer func() {
		a.Log.Println(name, " disconnected")
	}()

	for {
		bufSize := 256 // 1 header + 255 commands
		buf := make([]byte, bufSize)

		_, err := conn.Read(buf)
		if err != nil {
			err = fmt.Errorf("reading: %w", err)

			return err
		}

		a.Log.Print(name, " -> ", string(buf))

		_, err = conn.Write([]byte("Message received."))
		if err != nil {
			err = fmt.Errorf("error sending: %w", err)

			return err
		}

		if strings.TrimSpace(string(buf)) == "STOP" {
			a.Log.Println("Exiting TCP server!")

			return nil
		}
	}
}
