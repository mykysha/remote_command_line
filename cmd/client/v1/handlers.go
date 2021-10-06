package v1

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func ClientStarter() {
	addr := fmt.Sprintf("%s:%s", Host, Port)
	conn, err := net.Dial(Type, addr)
	if err != nil {
		log.Println(err)

		return
	}

	err = communicator(conn)
	if err != nil {
		log.Println(err)
	}
}

func communicator(conn net.Conn) error {
	for {
		err := serverWriter(conn)
		if err != nil {
			err = fmt.Errorf("communicator: %w", err)

			return err
		}

		err = serverReader(conn)
		if err != nil {
			err = fmt.Errorf("communicator: %w", err)

			return err
		}
	}
}

func serverWriter(conn net.Conn) error {
	bufSize := 256 // 1 header + 255 commands
	send := make([]byte, bufSize)

	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	_, err := writer.WriteString(">> ")
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	err = writer.Flush()
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	_, err = reader.Read(send)
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	send = bytes.Trim(send, "\n")

	_, err = conn.Write(send)
	if err != nil {
		err = fmt.Errorf("serverWriter: %w", err)

		return err
	}

	return nil
}

func serverReader(conn net.Conn) error {
	bufSize := 256 // 1 header + 255 commands
	receive := make([]byte, bufSize)

	_, err := conn.Read(receive)
	if err != nil {
		err = fmt.Errorf("serverReader: %w", err)

		return err
	}

	log.Printf("received: %v", string(receive))

	return nil
}
