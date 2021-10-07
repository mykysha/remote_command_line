package v1

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

var errUnknownAction = errors.New("using command line")

// fileCheck handles reading from file.
func (c *Client) fileCheck() error {
	answer := make([]byte, c.bufSize)

	_, err := c.writer.WriteString("if you want to read from file, type 'y' and file direction/name\n")
	if err != nil {
		return fmt.Errorf("file check: %w", err)
	}

	err = c.writer.Flush()
	if err != nil {
		return fmt.Errorf("file check: %w", err)
	}

	_, err = c.reader.Read(answer)
	if err != nil {
		return fmt.Errorf("file check: %w", err)
	}

	answer = bytes.Trim(answer, "\x00")

	ans := strings.Split(string(answer), " ")

	if flag := ans[0]; flag == "y" && len(ans) == 2 {
		addr := strings.Trim(ans[1], "\n")

		if _, err = os.Stat(addr); err == nil {
			file, err := os.Open(addr)
			if err != nil {
				return fmt.Errorf("file check: %w", err)
			}

			c.scanner = bufio.NewScanner(file)

			c.Log.Println("reading from: ", addr)

			return nil
		}

		return fmt.Errorf("file check: %w", err)
	}

	return fmt.Errorf("file check: %w", errUnknownAction)
}

// limitChecker finds if byte slice exceeds maximum buffer.
func (c Client) limitChecker(buf []byte, num int) error {
	if num >= c.bufSize {
		c.Log.Println("message exceeds limit in ", c.bufSize)

		messTooLong := fmt.Sprintf(
			"Your message is too long."+
				"Only first %d bytes (%s) will be sent.", c.bufSize, string(buf))

		_, err := c.writer.WriteString(messTooLong)
		if err != nil {
			return fmt.Errorf("checker: %w", err)
		}

		err = c.writer.Flush()
		if err != nil {
			return fmt.Errorf("checker: %w", err)
		}
	}

	return nil
}

// shutdown handles STOP request.
func (c Client) shutdown(buf []byte) bool {
	if bytes.Equal(buf, []byte("STOP")) {
		c.Log.Println("TCP client shutting down...")

		return true
	}

	return false
}
