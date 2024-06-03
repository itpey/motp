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

package motp

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

// MOTP represents a Mobile-OTP (mOTP) generator.
type MOTP struct {
	secret string // Secret key used for OTP generation.
	pin    string // PIN used for OTP generation.
	period uint   // Period (in seconds) for OTP validity.
	digits uint   // Number of digits in the generated OTP.
}

// MOTPOption is a functional option for configuring MOTP instances.
type MOTPOption func(*MOTP) error

// WithPeriod configures the period (in seconds) for the MOTP generator.
func WithPeriod(period uint) MOTPOption {
	return func(m *MOTP) error {
		if period < 1 {
			return errors.New("period must be positive")
		}
		m.period = period
		return nil
	}
}

// WithDigits configures the number of digits in the MOTP code.
func WithDigits(digits uint) MOTPOption {
	return func(m *MOTP) error {
		if digits < 1 || digits > 32 {
			return errors.New("digits must be in the range 1-32")
		}
		m.digits = digits
		return nil
	}
}

// New creates a new MOTP instance with the specified parameters and options.
// Default values:
// - period: 10 (default period in seconds)
// - digits: 6 (default number of digits in OTP)
func New(secret, pin string, options ...MOTPOption) (*MOTP, error) {
	// Initialize MOTP instance with default values
	m := &MOTP{
		secret: secret,
		pin:    pin,
		period: 10,
		digits: 6,
	}

	// Apply provided options to customize MOTP instance
	for _, option := range options {
		if err := option(m); err != nil {
			return nil, err
		}
	}

	return m, nil
}

// NewMOTP creates a new MOTP instance with the specified parameters and options.
// Default values:
// - period: 10 (default period in seconds)
// - digits: 6 (default number of digits in OTP)
func (m *MOTP) Generate(unixSeconds int64) (string, error) {
	if unixSeconds < 0 {
		return "", errors.New("unixSeconds must be non-negative")
	}

	// Determine the period based on the configured period
	epochPeriod := unixSeconds / int64(m.period)

	// Combine epoch period, secret key, and PIN to form the input for OTP generation
	hashInput := fmt.Sprintf("%d%s%s", epochPeriod, m.secret, m.pin)

	// Compute MD5 hash of the input string
	hasher := md5.New()
	_, err := hasher.Write([]byte(hashInput))
	if err != nil {
		return "", err
	}
	hashBytes := hasher.Sum(nil)

	// Convert hashBytes to hexadecimal string and truncate to the desired number of digits
	hashStr := hex.EncodeToString(hashBytes)
	return hashStr[:m.digits], nil
}

// GenerateCurrent generates a Mobile-OTP (mOTP) code based on the current Unix timestamp.
func (m *MOTP) GenerateCurrent() (string, error) {
	// Get the current Unix timestamp
	unixSeconds := time.Now().Unix()

	// Delegate to Generate method using current Unix timestamp
	return m.Generate(unixSeconds)
}
