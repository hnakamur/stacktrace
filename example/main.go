package main

import (
	"errors"
	"log"

	"github.com/hnakamur/stacktrace"
)

func logErrorWithStackTrace(msg interface{}) {
	log.Printf("error: %s\nstacktrace: %s\n", msg, stacktrace.StackWithSkip(2))
}

func b() {
	err := errors.New("some error")
	logErrorWithStackTrace(err)
}

func a() {
	b()
}

func main() {
	a()
}
