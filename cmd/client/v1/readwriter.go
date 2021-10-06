package v1

import (
	"bufio"
	"fmt"
	"os"
)

func (c Client) clToServer() (bool, error) {
	buf := make([]byte, c.bufSize)

	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	_, err := writer.WriteString(">> ")
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return false, err
	}

	err = writer.Flush()
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return false, err
	}

	num, err := reader.Read(buf)
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return false, err
	}

	buf[num-1] = 0

	err = c.checker(buf, num, writer)
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return false, err
	}

	err = c.serverWriter(buf)
	if err != nil {
		err = fmt.Errorf("communicator: %w", err)

		return false, err
	}

	if c.shutdown(buf) {
		return true, nil
	}

	return false, err
}

func (c Client) serverWriter(send []byte) error {
	if _, err := c.conn.Write(send); err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	return nil
}

func (c Client) serverReader() error {
	bufSize := 256 // 1 header + 255 commands
	receive := make([]byte, bufSize)

	if _, err := c.conn.Read(receive); err != nil {
		err = fmt.Errorf("serverReader: %w", err)

		return err
	}

	c.Log.Printf("received: %v", string(receive))

	return nil
}

func (c Client) checker(buf []byte, num int, writer *bufio.Writer) error {
	if num >= c.bufSize {
		c.Log.Println("message exceeds limit in ", c.bufSize)

		messTooLong := fmt.Sprintf(
			"Your message is too long."+
				"Only first %d bytes (%s) will be sent.", c.bufSize, string(buf))

		_, err := writer.WriteString(messTooLong)
		if err != nil {
			err = fmt.Errorf("serverWriter: %w", err)

			return err
		}

		err = writer.Flush()
		if err != nil {
			err = fmt.Errorf("serverWriter: %w", err)

			return err
		}
	}

	return nil
}

func (c Client) shutdown(buf []byte) bool {
	shutdownSign := make([]byte, c.bufSize)

	shutString := "STOP"

	_ = copy(shutdownSign, shutString)

	if string(buf) == string(shutdownSign) {
		c.Log.Println("TCP client shutting down...")

		return true
	}

	return false
}
