package main

import (
	"log"
	"github.com/galaco/vtf"
)

func main() {
	_,err := vtf.ReadFromFile("samples/read/test3.vtf")
	if err != nil {
		log.Fatal(err)
	}
}