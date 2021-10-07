package v1

import (
	"bytes"
	"fmt"
)

// clReader gets commands from command line.
func (c Client) clReader() (bool, error) {
	buf := make([]byte, c.bufSize)

	_, err := c.writer.WriteString("\n>> ")
	if err != nil {
		err = fmt.Errorf("clReader: %w", err)

		return false, err
	}

	err = c.writer.Flush()
	if err != nil {
		err = fmt.Errorf("clReader: %w", err)

		return false, err
	}

	num, err := c.reader.Read(buf)
	if err != nil {
		err = fmt.Errorf("clReader: %w", err)

		return false, err
	}

	buf[num-1] = 0

	err = c.limitChecker(buf, num)
	if err != nil {
		err = fmt.Errorf("clReader: %w", err)

		return false, err
	}

	buf = bytes.Trim(buf, "\x00")

	err = c.serverWriter(buf)
	if err != nil {
		err = fmt.Errorf("clReader: %w", err)

		return false, err
	}

	c.Log.Println("sent:\t", string(buf))

	if c.shutdown(buf) {
		return true, nil
	}

	return false, err
}

// serverWriter sends requests to the server.
func (c Client) serverWriter(send []byte) error {
	if _, err := c.conn.Write(send); err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	return nil
}

// serverReader gets responses from the server.
func (c Client) serverReader() error {
	receive := make([]byte, c.bufSize)

	if _, err := c.conn.Read(receive); err != nil {
		err = fmt.Errorf("serverReader: %w", err)

		return err
	}

	receive = bytes.Trim(receive, "\x00")

	c.Log.Println("received:", string(receive))

	_, err := c.writer.WriteString(string(receive))
	if err != nil {
		err = fmt.Errorf("serverReader: %w", err)

		return err
	}

	err = c.writer.Flush()
	if err != nil {
		err = fmt.Errorf("serverReader: %w", err)

		return err
	}

	return nil
}
