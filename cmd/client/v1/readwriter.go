package v1

import (
	"bufio"
	"bytes"
	"fmt"
)

// clReader gets commands from command line.
func (c Client) clReader() (bool, error) {
	buf := make([]byte, c.bufSize)

	_, err := c.writer.WriteString("\n>> ")
	if err != nil {
		return false, fmt.Errorf("clReader: %w", err)
	}

	err = c.writer.Flush()
	if err != nil {
		return false, fmt.Errorf("clReader: %w", err)
	}

	num, err := c.reader.Read(buf)
	if err != nil {
		return false, fmt.Errorf("clReader: %w", err)
	}

	buf[num-1] = 0

	err = c.limitChecker(buf, num)
	if err != nil {
		return false, fmt.Errorf("clReader: %w", err)
	}

	buf = bytes.Trim(buf, "\x00")

	if c.shutdown(buf) {
		return true, nil
	}

	err = c.serverWriter(buf)
	if err != nil {
		return false, fmt.Errorf("clReader: %w", err)
	}

	c.Log.Println("sent:\t", string(buf))

	return false, err
}

// fReader gets commands from file.
func (c Client) fReader() error {
	var fileInput []string

	c.scanner.Split(bufio.ScanLines)

	for c.scanner.Scan() {
		fileInput = append(fileInput, c.scanner.Text())
	}

	for _, val := range fileInput {
		buf := []byte(val)
		buf = bytes.Trim(buf, "\x00")

		if c.shutdown(buf) {
			return nil
		}

		err := c.serverWriter(buf)
		if err != nil {
			return fmt.Errorf("fReader: %w", err)
		}

		c.Log.Println("sent:\t", string(buf))

		message := "\nsent: " + string(buf) + "\n"

		_, err = c.writer.WriteString(message)
		if err != nil {
			return fmt.Errorf("fReader: %w", err)
		}

		err = c.writer.Flush()
		if err != nil {
			return fmt.Errorf("fReader: %w", err)
		}

		err = c.serverReader()
		if err != nil {
			return fmt.Errorf("fReader: %w", err)
		}
	}

	return nil
}

// serverWriter sends requests to the server.
func (c Client) serverWriter(send []byte) error {
	if _, err := c.conn.Write(send); err != nil {
		return fmt.Errorf("serverWriter: %w", err)
	}

	return nil
}

// serverReader gets responses from the server.
func (c Client) serverReader() error {
	receive := make([]byte, c.bufSize)

	if _, err := c.conn.Read(receive); err != nil {
		return fmt.Errorf("serverReader: %w", err)
	}

	receive = bytes.Trim(receive, "\x00")

	c.Log.Println("received:", string(receive))

	message := "\nserver: " + string(receive) + "\n"

	_, err := c.writer.WriteString(message)
	if err != nil {
		return fmt.Errorf("serverReader: %w", err)
	}

	err = c.writer.Flush()
	if err != nil {
		return fmt.Errorf("serverReader: %w", err)
	}

	return nil
}
