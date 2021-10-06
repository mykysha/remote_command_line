package config

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/nndergunov/AOCKM_Lab1/api"
)

type Server struct {
	Host     string
	Port     string
	Type     string
	listener net.Listener
}

func (srv Server) Init() {
	addr := fmt.Sprintf("%s:%s", srv.Host, srv.Port)

	l, err := net.Listen(srv.Type, addr)
	if err != nil {
		log.Println(err)

		return
	}

	srv.listener = l

	defer func(l net.Listener) {
		err = l.Close()
		if err != nil {
			log.Println(err)
		}
	}(srv.listener)

	log.Printf("listening on %s", addr)

	logger := log.New(os.Stdout, "server ", log.LstdFlags)
	a := &api.API{
		Log:      logger,
		Listener: srv.listener,
	}

	a.Init()
}
