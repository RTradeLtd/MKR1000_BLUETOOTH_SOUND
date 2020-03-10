package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/term"
)

func main() {
	s, err := term.Open("/dev/ttyACM0", term.Speed(9600))
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
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
		} else {
			continue
		}
		out := string(data[:read])
		out = strings.TrimSpace(out)
		if out == "" {
			continue
		}
		u, err := strconv.ParseUint(out, 10, 32)
		if err != nil {
			log.Println("error: ", err)
		}
		fmt.Println(u)
	}
}
