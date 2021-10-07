package v1

import (
	"bytes"
	"fmt"
)

// limitChecker finds if byte slice exceeds maximum buffer.
func (c Client) limitChecker(buf []byte, num int) error {
	if num >= c.bufSize {
		c.Log.Println("message exceeds limit in ", c.bufSize)

		messTooLong := fmt.Sprintf(
			"Your message is too long."+
				"Only first %d bytes (%s) will be sent.", c.bufSize, string(buf))

		_, err := c.writer.WriteString(messTooLong)
		if err != nil {
			err = fmt.Errorf("checker: %w", err)

			return err
		}

		err = c.writer.Flush()
		if err != nil {
			err = fmt.Errorf("checker: %w", err)

			return err
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
