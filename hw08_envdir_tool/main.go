package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal(errors.New("at least 2 arguments must be there"))
	}

	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(errors.New("Environment reading error: " + err.Error()))
	}

	os.Exit(RunCmd(os.Args[2:], env))
}
