package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// make cpu work
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}

	// block main goroutine
	select {}
}
