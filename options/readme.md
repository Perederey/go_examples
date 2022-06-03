# Option functions pattern in Go

Option functions pattern is a straightforward habit that allows us to overwrite default values..

Envision you are the creator of the new server. Your server has default values: host, maximum idle connections, and 10+ others.

```go
// Server is a simple example of a struct
type Server struct {
	Name string
	Host string
  // other values...
	MaxIdleConnections   uint
	MaxSessionConnection time.Duration
}
```

## Option #1. Create a new structure every time.

Possible, but thank you.

You can produce every time a new structure. But this method is so monotonous, and you have to be very careful with many settings.

```go
s := Server{
		Name:                 "my name",
		Host:                 "host",
    // don't forget about 10+ options.
		MaxIdleConnections:   20,
		MaxSessionConnection: 5 * time.Minute,
	}
```

## Option #2.  Use Option functions pattern

The Best option for options.

All process has only three simple steps:

- Define in `New` method a default struct
- Add a function `type Option func(s *Server)`
- Use it, for example `WithName(name) Option`

Let’s implement it.

Full example below:

```go
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
```

## Output

```bash
{default http://default-eu 20 5m0s}
{default https://another-host.eu 20 5m0s}
{default https://eu.ru 50 5m0s}
{new Name https://another-host 20 50µs}
```