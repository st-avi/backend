package utility

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenTotp(email string) (keySecret, base64QR string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Stavi",
		AccountName: email,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		return "", "", err
	}

	// image -> png bytes -> base64
	img, err := key.Image(200, 200)
	if err != nil {
		return "", "", err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", "", err
	}
	base64QR = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	return key.Secret(), base64QR, nil
}

func ValidateTotp(code, secret string) bool {
	return totp.Validate(code, secret)
}
