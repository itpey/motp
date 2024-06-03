[//]: # "Title: motp"
[//]: # "Author: itpey"
[//]: # "Attendees: itpey"
[//]: # "Tags: #itpey #go #motp #golang #go-lang #cli #password #otp"

<h1 align="center">
 MOTP
</h1>

<p align="center">
motp is a Go package that provides a <a href="https://motp.sourceforge.net">Mobile-OTP (mOTP)</a> generator. This package allows you to generate time-based one-time passwords (OTP) using a secret key and a PIN.
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/itpey/motp">
    <img src="https://pkg.go.dev/badge/github.com/itpey/motp.svg" alt="itpey motp Go Reference">
  </a>
  <a href="https://github.com/itpey/motp/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/itpey/motp" alt="itpey motp license">
  </a>
</p>

# Features

- Configurable period for OTP validity
- Configurable number of digits in the OTP
- Generates OTP based on the current Unix timestamp

## Installation

To install the package, use the following command:

```bash
go get github.com/itpey/motp
```

## Usage

### Creating a New MOTP Instance

You can create a new MOTP instance by specifying a secret key and a PIN. Optionally, you can configure the period and the number of digits for the OTP.

```go
package main

import (
    "fmt"
    "log"
    "github.com/itpey/motp"
)

func main() {
    // Create a new MOTP instance with default settings
    motpInstance, err := motp.New("your_secret_key", "your_pin")
    if err != nil {
        log.Fatalf("Error creating MOTP instance: %v", err)
    }

    // Generate an OTP based on the current Unix timestamp
    otp, err := motpInstance.GenerateCurrent()
    if err != nil {
        log.Fatalf("Error generating OTP: %v", err)
    }

    fmt.Printf("Generated OTP: %s\n", otp)
}
```

### Configuring the Period and Number of Digits

You can customize the period and the number of digits in the OTP by using functional options:

```go
package main

import (
    "fmt"
    "log"
    "github.com/itpey/motp"
)

func main() {
    // Create a new MOTP instance with custom period and digits
    motpInstance, err := motp.New("your_secret_key", "your_pin", motp.WithPeriod(30), motp.WithDigits(8))
    if err != nil {
        log.Fatalf("Error creating MOTP instance: %v", err)
    }

    // Generate an OTP based on the current Unix timestamp
    otp, err := motpInstance.GenerateCurrent()
    if err != nil {
        log.Fatalf("Error generating OTP: %v", err)
    }

    fmt.Printf("Generated OTP: %s\n", otp)
}
```

# CLI Tool

The motp-cli is a command-line tool for generating one-time passwords (OTPs) using the <a href="https://motp.sourceforge.net">Mobile-OTP (mOTP)</a>algorithm. It allows you to specify a secret key and a PIN to generate a unique OTP based on the current time period.

## Installation

Make sure you have Go installed and configured on your system. Use go install to install motp-cli:

```bash
go install github.com/itpey/motp/cmd/motp@latest
```

Ensure that your `GOBIN` directory is in your `PATH` for the installed binary to be accessible globally.

## Usage

**$ motp** --secret `your_secret_key` --pin `your_pin` --duration `duration` --length `length`

Flags

- `--secret, -s`: mOTP secret value (often hex or alphanumeric digits) **[required]**
- `--pin, -p`: mOTP PIN value (usually 4 digits) **[required]**
- `--duration, -d`: Duration of mOTP codes in seconds (default 10s) **[optional]**
- `--length, -l`: Length of mOTP output (default 6 characters) **[optional]**
-

## Example

**$ motp** --secret `mysecret` --pin `1234` --duration `30` --length `8`

# Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/itpey/motp/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

# License

motp is open-source software released under the MIT License. You can find a copy of the license in the [LICENSE](https://github.com/itpey/motp/blob/main/LICENSE) file.

# Author

motp was created by [itpey](https://github.com/itpey)
