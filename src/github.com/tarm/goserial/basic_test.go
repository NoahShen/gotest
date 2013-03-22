package serial

import (
	"fmt"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	if testing.Short() {
		return
	}
	c0 := &Config{Name: "/dev/ttyUSB0", Baud: 9600}

	/*
		c1 := new(Config)
		c1.Name = "COM5"
		c1.Baud = 115200
	*/

	s, err := OpenPort(c0)
	if err != nil {
		t.Fatal(err)
	}

	//_, err = s.Write([]byte("Hello World!"))
	//if err != nil {
	//	t.Fatal(err)
	//}

	for {
		buf := make([]byte, 128)
		c, err2 := s.Read(buf)
		fmt.Println(c)
		if err2 != nil {
			t.Fatal(err2)
		}
		fmt.Println(string(buf[0:c]))
		time.Sleep(1000)
	}
}
