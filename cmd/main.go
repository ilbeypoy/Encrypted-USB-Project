// cmd/main.go
package main

import (
	"007/usb"
	"fmt"
	"time"
)

func main() {
	for {
		usbDevices, err := usb.ListUSBDevices()
		if err != nil {
			fmt.Println("Error listing USB devices:", err)
		} else {
			fmt.Println("Connected USB devices:", usbDevices)
		}

		// Add logic to check if the required flash drives are present and take appropriate actions

		time.Sleep(5 * time.Second)
	}
}
