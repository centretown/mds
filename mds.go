package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/tarm/serial"
	// "github.com/centretown/mdsprocess"
)

func main() {
	flag.Parse()
	//if verbose {
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s=%v\n", f.Name, f.Value)
	})

	settings, err := MakeSettings()
	if err != nil {
		fmt.Printf("while making settings: %v\n", err)
		return
	}

	fmt.Printf("settings: %v\n", *settings)

	httpQuit := make(chan int)
	go listenAndServeHTTP(settings, httpQuit)
	go listenAndServeSerial(settings, httpQuit)
	<-httpQuit

}

func listenAndServeHTTP(settings *Settings, quit chan<- int) (err error) {
	fs := http.FileServer(http.Dir("./web/home"))
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(fs)
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler: r,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("http.ListenAndServe: %v\n", err)
	}
	quit <- 1
	return
}

func listenAndServeSerial(settings *Settings, quit <-chan int) (err error) {
	config := &serial.Config{
		Name: settings.serialPort,
		Baud: settings.serialBaud,
	}
	port, err := SerialOpen(config)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer port.Close()

	out := make(chan string)
	quitScan := make(chan int)
	quitPrint := make(chan int)

	go ScanLines(port, out, quitScan)
	go PrintLines("console", os.Stdout, out, quitPrint)

	<-quit
	quitScan <- 1
	quitPrint <- 1
	return
}
