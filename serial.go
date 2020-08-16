package main

import (
	"github.com/tarm/serial"
)

// SerialOpen - opens a serial port
func SerialOpen(c *serial.Config) (s *serial.Port, err error) {
	s, err = serial.OpenPort(c)
	return
}
