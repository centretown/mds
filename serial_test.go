package main

import (
	"os"
	"testing"
	"time"

	"github.com/tarm/serial"
)

const serialFileName = "/dev/ttyUSB0"

func TestScanner(t *testing.T) {
	settings, err := MakeSettings()
	config := &serial.Config{Name: settings.serialPort, Baud: settings.serialBaud}
	port, err := SerialOpen(config)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer port.Close()

	out := make(chan string)
	quitScan := make(chan int)
	quitPrint := make(chan int)
	go ScanLines(port, out, quitScan)
	go PrintLines("console-c", os.Stdout, out, quitPrint)
	go PrintLines("console-b", os.Stdout, out, quitPrint)
	time.Sleep(time.Second * 10)
	quitScan <- 1
	quitPrint <- 1
	time.Sleep(time.Second * 1)

}
