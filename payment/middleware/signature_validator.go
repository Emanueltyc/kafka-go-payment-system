package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SignatureValidator(c *fiber.Ctx) error {
	secret := os.Getenv("WEBHOOK_SECRET")

	if header := c.Get("X-Mock-Signature"); header != "" {
		err :=  validateSignature(c.Body(), header, secret)

		if err != nil {
			c.Status(fiber.StatusUnauthorized).SendString(err.Error())
			return nil
		}

		err = c.Next()

		return err
	}

	c.Status(fiber.StatusUnauthorized).SendString("Not authorized, missing Signature header!")

	return nil
}

func validateSignature(payload []byte, header string, secret string) error {
	parts := strings.Split(header, ",")
	var timestamp, signature string

	for _, p := range parts {
		if after, ok := strings.CutPrefix(p, "t="); ok  {
			timestamp = after
		}
		if after, ok :=strings.CutPrefix(p, "v1="); ok  {
			signature = after
		}
	}

	if timestamp == "" || signature == "" {
		return errors.New("invalid signature header")
	}

	signedPayload := fmt.Sprintf("%s.%s", timestamp, payload)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature), []byte(expected)) {
		return errors.New("signature mismatch")
	}

	return nil
}
