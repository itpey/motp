// Copyright (c) 2024 itpey
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/itpey/motp"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "motp-cli",
		Usage:   "Generate one-time passwords (OTP) using Mobile-OTP (mOTP)",
		Version: "0.1.0",
		Description: `motp-cli is a command-line tool for generating one-time passwords (OTPs)
using the Mobile-OTP (mOTP) algorithm. It allows you to specify a secret key and a PIN
to generate a unique OTP based on the current time period.`,
		Copyright: "MIT license\nFor more information, visit the GitHub repository: https://github.com/itpey/motp",
		Authors:   []*cli.Author{{Name: "itpey", Email: "itpey@github.com"}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "secret",
				Aliases:  []string{"s"},
				Usage:    "mOTP secret value (often hex or alphanumeric digits)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pin",
				Aliases:  []string{"p"},
				Usage:    "mOTP PIN value (usually 4 digits)",
				Required: true,
			},
			&cli.UintFlag{
				Name:     "duration",
				Aliases:  []string{"d"},
				Usage:    "Duration of mOTP codes in seconds (default 10s)",
				Value:    10,
				Required: false,
			},
			&cli.UintFlag{
				Name:     "length",
				Aliases:  []string{"l"},
				Usage:    "Length of mOTP output (default 6 characters)",
				Value:    6,
				Required: false,
			},
		},

		Action: func(c *cli.Context) error {
			// Get secret and pin from command-line flags
			secret := c.String("secret")
			pin := c.String("pin")
			period := c.Uint("duration")
			length := c.Uint("length")

			motp, err := motp.New(secret, pin, motp.WithPeriod(period), motp.WithDigits(length))
			if err != nil {
				fmt.Printf("Error creating MOTP instance: %v\n", err)
				return err
			}

			// Generate mOTP code based on the current Unix timestamp
			otp, err := motp.GenerateCurrent()
			if err != nil {
				return err
			}

			// Output the generated OTP
			fmt.Print(otp)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
