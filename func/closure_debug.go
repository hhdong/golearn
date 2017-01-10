package main

import (
	"log"
	"runtime"
)

func main() {

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
		return
	}
	where()

	var wh = log.Print

	wh()
}
