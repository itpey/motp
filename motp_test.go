package motp

import (
	"testing"
)

func TestNewMOTPDefaultValues(t *testing.T) {
	secret := "testsecret"
	pin := "1234"

	motp, err := New(secret, pin)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if motp.secret != secret {
		t.Errorf("Expected secret: %s, got: %s", secret, motp.secret)
	}

	if motp.pin != pin {
		t.Errorf("Expected pin: %s, got: %s", pin, motp.pin)
	}

	if motp.period != 10 {
		t.Errorf("Expected period: %d, got: %d", 10, motp.period)
	}

	if motp.digits != 6 {
		t.Errorf("Expected digits: %d, got: %d", 6, motp.digits)
	}
}

func TestNewMOTPWithCustomValues(t *testing.T) {
	secret := "testsecret"
	pin := "1234"
	period := uint(30)
	digits := uint(8)

	motp, err := New(secret, pin, WithPeriod(period), WithDigits(digits))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if motp.period != period {
		t.Errorf("Expected period: %d, got: %d", period, motp.period)
	}

	if motp.digits != digits {
		t.Errorf("Expected digits: %d, got: %d", digits, motp.digits)
	}
}

func TestNewMOTPWithInvalidPeriod(t *testing.T) {
	secret := "testsecret"
	pin := "1234"
	period := uint(0)

	_, err := New(secret, pin, WithPeriod(period))
	if err == nil {
		t.Fatalf("Expected error for invalid period, got none")
	}
}

func TestNewMOTPWithInvalidDigits(t *testing.T) {
	secret := "testsecret"
	pin := "1234"
	digits := uint(0)

	_, err := New(secret, pin, WithDigits(digits))
	if err == nil {
		t.Fatalf("Expected error for invalid digits, got none")
	}
}

func TestGenerateOTP(t *testing.T) {
	secret := "testsecret"
	pin := "1234"
	period := uint(10)
	digits := uint(6)
	unixSeconds := int64(1625097600)

	motp, err := New(secret, pin, WithPeriod(period), WithDigits(digits))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	otp, err := motp.Generate(unixSeconds)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedLength := int(digits)
	if len(otp) != expectedLength {
		t.Errorf("Expected OTP length: %d, got: %d", expectedLength, len(otp))
	}
}

func TestGenerateCurrentOTP(t *testing.T) {
	secret := "testsecret"
	pin := "1234"

	motp, err := New(secret, pin)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	otp, err := motp.GenerateCurrent()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedLength := int(motp.digits)
	if len(otp) != expectedLength {
		t.Errorf("Expected OTP length: %d, got: %d", expectedLength, len(otp))
	}
}

func TestGenerateOTPWithNegativeUnixSeconds(t *testing.T) {
	secret := "testsecret"
	pin := "1234"
	unixSeconds := int64(-1)

	motp, err := New(secret, pin)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, err = motp.Generate(unixSeconds)
	if err == nil {
		t.Fatalf("Expected error for negative unixSeconds, got none")
	}
}
