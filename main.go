package main

import (
	"log"

	"github.com/ringo199/spider/ui"
)

func main() {
	p := ui.Initial()
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
