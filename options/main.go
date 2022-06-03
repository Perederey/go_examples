package main

import (
	"fmt"
	"time"
)

// Server is a simple example of a struct
type Server struct {
	Name string
	Host string

	MaxIdleConnections   uint
	MaxSessionConnection time.Duration
}

// Option is a function for options
type Option func(s *Server)

// New server with options
func New(options ...Option) Server {
	// define a default server
	s := Server{
		Name:                 "default",
		Host:                 "http://default-eu",
		MaxIdleConnections:   20,
		MaxSessionConnection: 5 * time.Minute,
	}

	// apply options for the created server
	for _, fn := range options {
		fn(&s)
	}

	return s
}

// WithHost option overwrites a default value.
func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

// WithMaxIdleConnections option overwrites a default value.
func WithMaxIdleConnections(maxConnections uint) Option {
	return func(s *Server) {
		s.MaxIdleConnections = maxConnections
	}
}

// main simple example
func main() {
	// default server without options
	defaultServer := New()
	fmt.Println(defaultServer)

	// example, how we can overwrite one value
	serverWithHost := New(WithHost("https://another-host.eu"))
	fmt.Println(serverWithHost)

	// or even 2 values
	serverWithHostAndMaxIdleConnections := New(
		WithHost("https://eu.ru"),
		WithMaxIdleConnections(50),
	)

	fmt.Println(serverWithHostAndMaxIdleConnections)

	// or even create a custom function
	// but usually Option function isn't a public one.
	// the creators of packages allow us use pre-built options (with) functions.
	customServer := New(func(s *Server) {
		s.Name = "new Name"
		s.MaxSessionConnection = 50 * time.Microsecond
		s.Host = "https://another-host"
	})

	fmt.Println(customServer)
}
