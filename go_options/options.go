package main

import "fmt"

type Client struct {
	httpPort, grpcPort int
	host               string
}
type Options func(*Client)

func WithHttpPort(port int) Options {
	return func(c *Client) { c.httpPort = port }
}
func WithGrpcPort(port int) Options {
	return func(c *Client) { c.grpcPort = port }
}

func WithHost(host string) Options {
	return func(c *Client) { c.host = host }
}

func NewClient(host string, options ...Options) *Client {
	c := &Client{host: host}
	for _, option := range options {
		option(c)
	}
	return c
}

func main() {
	httpPortOption := WithHttpPort(8080)
	grpcPortOption := WithGrpcPort(1997)
	c := NewClient("localhost", httpPortOption, grpcPortOption)
	fmt.Printf("httpPort : %d grpcPort : %d host : %s", c.httpPort, c.grpcPort, c.host)
}
