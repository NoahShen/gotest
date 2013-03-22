package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyS1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write([]byte("Hello World!"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 128)
		n, err2 := s.Read(buf)
		if err2 != nil {
			log.Fatal("error:", err2)
		}
		fmt.Print(string(buf[:n]))
	}
}
