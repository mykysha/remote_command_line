package v1

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	conn    net.Conn
	Host    string
	Port    string
	addr    string
	Type    string
	bufSize int
	writer  *bufio.Writer
	reader  *bufio.Reader
	scanner *bufio.Scanner
	Log     *log.Logger
}

// Init starts Client work.
func (c *Client) Init() {
	c.addr = fmt.Sprintf("%s:%s", c.Host, c.Port)
	c.writer = bufio.NewWriter(os.Stdout)
	c.reader = bufio.NewReader(os.Stdin)

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

// communicator manages TCP communication.
func (c Client) communicator() error {
	err := c.fileCheck()
	if err == nil {
		err = c.fileHandler()
		if err != nil {
			return fmt.Errorf("communicator: %w", err)
		}

		return nil
	}
	c.Log.Println(err)

	for {
		end, err := c.clReader()
		if err != nil {
			return fmt.Errorf("communicator: %w", err)
		}

		if end {
			return nil
		}

		err = c.serverReader()
		if err != nil {
			return fmt.Errorf("communicator: %w", err)
		}
	}
}

// fileHandler manages reading commands from file.
func (c Client) fileHandler() error {
	if err := c.fReader(); err != nil {
		return fmt.Errorf("file handler: %w", err)
	}

	return nil
}
