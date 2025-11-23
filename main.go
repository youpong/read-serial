package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

func main() {
	// Example micro:bit serial port
	//   Linux:   /dev/ttyACM0"
	//   macOS:   /dev/tty.usbmodem1101
	//   Windows: COM3
	c := &serial.Config{
		Name: "/dev/tty.usbmodem1101",
		Baud: 115200, // micro:bit standard baud
	}

	if len(os.Args) == 2 {
		fmt.Println(os.Args[1])
		c.Name = os.Args[1]
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	reader := bufio.NewScanner(s)

	for reader.Scan() {
		line := reader.Text()
		fmt.Println("RAW:", line)

		// The format for serial.writeValue() in MakeCode: ‘A0:123’
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				label := parts[0]
				value := parts[1]
				fmt.Printf("Label=%s  Value=%s\n", label, value)
			}
		}
	}
}
