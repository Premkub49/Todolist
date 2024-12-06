package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = os.Getenv("JWT_SECRET")

func registerData(c *fiber.Ctx) error {
	user := new(User)
	if err = c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user.Password = string(hashPassword)
	err = createUser(user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func loginData(c *fiber.Ctx) error {
	user := new(User)
	if err = c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	selectUser, err := getUser(user)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "WrongPassword",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}

func checkCookie(c *fiber.Ctx) error {
	type Body struct {
		Token string `json:"token"`
	}
	body := new(Body)
	if err := c.BodyParser(body); err != nil {
		log.Println("Error cookie Send ", err)
		log.Println(body.Token)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	_, err := jwt.ParseWithClaims(body.Token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		log.Println("key error ", err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.SendStatus(fiber.StatusOK)
}
