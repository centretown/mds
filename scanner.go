package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// ScanLines from a reader and output to each line
func ScanLines(reader io.Reader, out chan<- string, quit <-chan int) (err error) {
	bufReader := bufio.NewReaderSize(reader, 512)

	var (
		str string
	)

	for {
		select {
		case <-quit:
			fmt.Println("end")
			return
		default:
			str, err = bufReader.ReadString('\n')
			if err != nil {
				return
			}
			out <- str
			time.Sleep(time.Millisecond)
		}
	}
}

// PrintLines -
func PrintLines(id string, writer io.Writer, in <-chan string, quit <-chan int) {
	for {
		select {
		case msg := <-in:
			writer.Write([]byte(id))
			writer.Write([]byte(msg))

		case <-quit:
			fmt.Println("end")
			return

		default:
			time.Sleep(time.Millisecond)
		}
	}
}
