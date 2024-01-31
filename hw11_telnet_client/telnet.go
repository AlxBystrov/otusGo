package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type customTelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (cl *customTelnetClient) Connect() error {
	var err error
	if cl.conn, err = net.DialTimeout("tcp", cl.address, cl.timeout); err != nil {
		return err
	}
	return nil
}

func (cl *customTelnetClient) Close() error {
	return cl.conn.Close()
}

func (cl *customTelnetClient) Send() error {
	_, err := io.Copy(cl.conn, cl.in)
	return err
}

func (cl *customTelnetClient) Receive() error {
	_, err := io.Copy(cl.out, cl.conn)
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &customTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
