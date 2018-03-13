package main

import (
	"github.com/hoffx/EduRM/controller"
	"log"
	"sync"
)

func main() {
	c, err := controller.NewController("interpreter/test.txt", 16)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	go c.Process()

	for {
		select {
		case ctx := <- c.ContextChan:
			...
		}
	}
}
