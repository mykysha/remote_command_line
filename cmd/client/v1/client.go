package v1

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	conn    net.Conn
	Host    string
	Port    string
	addr    string
	Type    string
	bufSize int
	Log     *log.Logger
}

func (c *Client) Init() {
	c.addr = fmt.Sprintf("%s:%s", c.Host, c.Port)

	conn, err := net.Dial(c.Type, c.addr)
	if err != nil {
		c.Log.Println(err)

		return
	}

	c.Log.Println("connected to ", c.addr, " via ", c.Type)

	c.conn = conn

	c.bufSize = 256 // 1 header + 255 commands

	err = c.communicator()
	if err != nil {
		c.Log.Println(err)
	}
}

func (c Client) communicator() error {
	for {
		end, err := c.clToServer()
		if err != nil {
			err = fmt.Errorf("communicator: %w", err)

			return err
		}

		if end {
			return nil
		}

		err = c.serverReader()
		if err != nil {
			err = fmt.Errorf("communicator: %w", err)

			return err
		}
	}
}
