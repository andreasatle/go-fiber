package jwt

import (
	"errors"
	"time"

	"github.com/andreasatle/go-fiber/src/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	ID uint
}

func EncodeJWTToken(c *fiber.Ctx, user database.User) error {
	token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["sub"] = user.Username
	claims["user_id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour).Unix()
	secretKey := []byte("your_secret_key")
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return err
    }
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	})

	return nil
}

func DecodeJWTToken(c *fiber.Ctx) (*User, error) {
	tokenString := c.Cookies("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You can validate the signing method and return the secret key here
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		username := claims["sub"].(string)
		user_id := claims["user_id"].(float64)
		return &User{Username: username, ID: uint(user_id)}, nil
	}
	return nil, errors.New("token is not valid")

}
