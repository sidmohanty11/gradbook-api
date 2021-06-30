package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func User(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostQuestion(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostAnswer(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostStory(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetQuestion(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAnswer(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetStory(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAQuestion(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAnAnswer(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAStory(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}
