package main

import (
	"log"
	"time"
)

func track(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func trackFunc(name string, callback func() string) {
	start := time.Now()
	output := callback()
	log.Printf("%s function took %s. Output => %s", name, time.Since(start), output)
}