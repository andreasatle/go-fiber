package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/andreasatle/go-fiber/src/crypt"
	"github.com/andreasatle/go-fiber/src/database"
	"github.com/andreasatle/go-fiber/src/templates"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
    Username string `binding:"required"`
    Password string `binding:"required"`
}

type BaseType struct {
	Title string
	IsAuthenticated bool
	Username string
	ID uint
	ErrorMessage string
}

type User struct {
	Username string
	ID uint
}
func GetPublicHome(c *fiber.Ctx) error {
	data := GetData(c, "Home")
	c.Set("Content-Type", "text/html")
	return templates.PublicHome.Execute(c.Response().BodyWriter(), data)
}

func GetPublicResume(c *fiber.Ctx) error {
	data := GetData(c, "Resume")
	c.Set("Content-Type", "text/html")
	return templates.PublicResume.Execute(c.Response().BodyWriter(), data)
}

func GetPublicContact(c *fiber.Ctx) error {
	data := GetData(c, "Contact")
	c.Set("Content-Type", "text/html")
	return templates.PublicContact.Execute(c.Response().BodyWriter(), data)
}

func GetPrivateHome(c *fiber.Ctx) error {
	data := GetData(c, "Private Home")

	c.Set("Content-Type", "text/html")
	return templates.PrivateHome.Execute(c.Response().BodyWriter(), data)
}

func GetAuthLogin(c *fiber.Ctx) error {
	errorMessage := c.Query("error")
	username := c.Query("username")
	c.Set("Content-Type", "text/html")
	return templates.AuthLogin.Execute(c.Response().BodyWriter(), BaseType {
		Title: "Login", 
		IsAuthenticated: false, 
		ErrorMessage: errorMessage, 
		Username: username,
	})
}

func GetAuthLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	c.Redirect("/public/home", fiber.StatusFound)
	return nil
}

func PostAuthLogin(c *fiber.Ctx) error {
	username := c.FormValue("Username")
	password := c.FormValue("Password")

	var user database.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.Redirect("/auth/login?error=Username not found", fiber.StatusFound)
		return nil
	}
	if err := crypt.ComparePasswords(user.Password, password); err != nil {
		login := fmt.Sprintf("/auth/login?error=Invalid password&username=%v", username)
		c.Redirect(login, fiber.StatusFound)
		return nil
	}

	if err := EncodeJWTToken(c, user); err != nil {
		c.Redirect("/auth/login?error=Token error", fiber.StatusFound)
	} else {
		c.Redirect("/private/home", fiber.StatusFound)
	}

	return nil
}

func Redirect(route string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Redirect(route, fiber.StatusFound)
	}
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

func GetData(c *fiber.Ctx, title string) BaseType {
	user, err := DecodeJWTToken(c)
	if err != nil {
		return BaseType{Title: title, IsAuthenticated: false}
	}
	return BaseType{Title: title, IsAuthenticated: true, Username: user.Username, ID: user.ID}
}