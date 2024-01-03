// usb/usb.go
package usb

import (
	"os/exec"
	"strings"
)

// ListUSBDevices returns a list of connected USB devices
func ListUSBDevices() ([]string, error) {
	cmd := exec.Command("system_profiler", "SPUSBDataType", "-json")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse the output to extract USB device information
	// Modify this according to the actual output format
	// Example: "USB 3.0 Bus" is assumed here; adapt to your system's output
	devices := []string{}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "USB 3.0 Bus") {
			devices = append(devices, line)
		}
	}

	return devices, nil
}
