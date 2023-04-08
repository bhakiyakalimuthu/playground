package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//server

type Message struct {
	addrs, msg string
}

type Server struct {
	addrs  string
	server net.Listener
	quitCh chan struct{}
	msgCh  chan Message
}

func NewServer(addrs string) *Server {
	return &Server{
		addrs: addrs,
		msgCh: make(chan Message, 10),
	}
}

// create connection

func (s *Server) Start(ctx context.Context) error {
	server, err := net.Listen("tcp", s.addrs)
	if err != nil {
		return err
	}
	s.server = server
	defer server.Close()
	go s.AcceptConnection(ctx)
	<-s.quitCh
	close(s.msgCh)
	return nil
}

// accept connection

func (s *Server) AcceptConnection(ctx context.Context) {
	for {
		conn, err := s.server.Accept()
		if err != nil {
			fmt.Printf("failed to accept connction %v", err)
			continue
		}
		fmt.Printf("Remote connection addrs %s", conn.RemoteAddr())
		go s.ReadLoop(ctx, conn)
	}
}

// read data

func (s *Server) ReadLoop(ctx context.Context, conn net.Conn) {
	msg := new(bytes.Buffer)
	for {
		n, err := io.CopyN(msg, conn, 4000)
		if err != nil {
			fmt.Printf("failed to read message %v", err)
			continue
		}
		fmt.Printf("bytes received :%v", msg.Bytes())
		fmt.Printf("number of bytes received over connection :%d", n)
	}
}

func (s *Server) SendFiles(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	//n, err := conn.Write(file)
	if err != nil {
		return err
	}
	fmt.Printf("bytes sent over connection: %d", n)
	return nil
}
func main() {
	s := NewServer("localhost:3000")
	go func() {
		<-time.After(time.Second * 3)
		s.SendFiles(200000)
	}()
	log.Fatal(s.Start(context.Background()))
}
