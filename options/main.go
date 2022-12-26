// It is a function programming pattern that is used to provide optional arguments to a function

package main

import (
	"fmt"
	"time"
)

type Config struct {
	url     string
	port    string
	timeout time.Duration
}

type Option func(*Config)

func WithPort(port string) Option {
	return func(c *Config) {
		c.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.timeout = timeout
	}
}

func NewConfig(url string, opts ...Option) *Config {
	config := &Config{
		url: url,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func main() {
	appConfig := NewConfig("http://api.google.com")
	fmt.Println(appConfig)

	appConfig = NewConfig("http://api.google.com", WithPort("3000"))
	fmt.Println(appConfig)

	appConfig = NewConfig("http://api.google.com", WithPort("3000"), WithTimeout(30*time.Second))
	fmt.Println(appConfig)
}
