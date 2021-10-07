package api

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"

	cline "github.com/nndergunov/AOCKM_Lab1/pkg"
)

type API struct {
	Log      *log.Logger
	Listener net.Listener
}

// Init starts API work.
func (a API) Init() {
	a.connecter()
}

// connecter manages new client connections.
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

// handler deals with one particular connection.
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

		buf = bytes.Trim(buf, "\x00")

		a.Log.Print("received:", name, " -> ", string(buf))

		buf = cline.Route(string(buf))

		_, err = conn.Write(buf)
		if err != nil {
			err = fmt.Errorf("error sending: %w", err)

			return err
		}

		a.Log.Println("sent:\t", string(buf))
	}
}
