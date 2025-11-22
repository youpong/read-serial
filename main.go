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
	// macOS の micro:bit のポート例: /dev/tty.usbmodem1101
	// Windows: COM3 など
	c := &serial.Config{
		Name: "/dev/tty.usbmodem1101", // ←実際の micro:bit のポート名に変更
		Baud: 115200,                  // micro:bit の標準 baud
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

		// MakeCode の serial.writeValue() の形式: "A0:123"
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
