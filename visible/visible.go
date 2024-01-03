// visible/visible.go
package main

import (
	"007/internal/crypto"
	"007/usb"
	"fmt"
	"time"
)

func main() {
	// Example usage of encryption and decryption
	key := []byte("32-byte-secret-key-1234567890123456")

	// Data to be encrypted
	data := []byte("super_secret_password")

	// Encrypt data
	encryptedData, err := crypto.Encrypt(data, key)
	if err != nil {
		panic(err)
	}

	// Decrypt data
	decryptedData, err := crypto.Decrypt(encryptedData, key)
	if err != nil {
		panic(err)
	}

	// Use decrypted data (e.g., print it)
	fmt.Println("Decrypted data:", string(decryptedData))

	// Monitor USB devices
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
