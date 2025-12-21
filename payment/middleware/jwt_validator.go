package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTValidator(c *fiber.Ctx) error {
	secret := []byte(os.Getenv("JWT_SECRET"))

	if header := c.Get("Authorization"); strings.HasPrefix(header, "Bearer") {
		tokenString := strings.Split(header, " ")[1]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return secret, nil
		})
	
		if err != nil {
			c.Status(fiber.StatusUnauthorized).SendString("Not authorized, token failed!")
			return nil
		}
	
		err = c.Next()
	
		return err
	}

	c.Status(fiber.StatusUnauthorized).SendString("Not authorized, missing authorization header!")
	
	return nil
}
