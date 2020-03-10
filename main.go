package main

import (
	"fmt"
	"log"

	"github.com/pkg/term"
)

func main() {
	s, err := term.Open("/dev/ttyACM0", term.Speed(9600))
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, err := s.Buffered()
		if err != nil {
			log.Println("error: ", err)
		}
		var (
			data = make([]byte, n)
			read int
		)
		if n > 0 {
			read, err = s.Read(data)
		}
		if err != nil {
			log.Println("error: ", err)
		}
		fmt.Println(string(data[:read]))
	}
}
