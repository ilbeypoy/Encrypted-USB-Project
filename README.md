# Encrypted USB Project

This project demonstrates a simple system for handling encrypted data on USB drives using a 32-bit encryption system. It consists of a Go application that encrypts and decrypts data using the Advanced Encryption Standard (AES) with a 32-bit key size.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Overview

The project comprises three main components:

1. **Internal Package (`internal/crypto/crypto.go`):** Contains the implementation of the encryption and decryption functions using AES with a 32-bit key.

2. **USB Package (`internal/usb/usb.go`):** Provides a simple USB device listing functionality.

3. **Visible Package (`visible/visible.go`):** Illustrates the usage of the encryption and USB listing components.

## Prerequisites

- Go (Golang) installed on your machine.
- Operating system compatible with the provided USB device listing code (currently assumes macOS).

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ilbeypoy/Encrypted-USB-Project.git
    ```

2. Change into the project directory:

    ```bash
    cd encrypted-usb-project
    ```

## Usage

1. **Encryption and Decryption:**

    - Ensure Go is installed on your machine.
    
    - Run the following command to encrypt and decrypt data:

        ```bash
        go run visible/visible.go
        ```

2. **USB Device Listing:**

    - The USB device listing functionality currently supports macOS. Modify the code in `internal/usb/usb.go` to adapt it to your operating system if needed.

    - Run the following command to list connected USB devices:

        ```bash
        go run cmd/main.go
        ```

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.

## License

This project is licensed under the [Apache License](LICENSE).
