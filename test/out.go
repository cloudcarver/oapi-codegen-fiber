package main 

import "github.com/gofiber/fiber/v2"

type AuthFunc func(c *fiber.Ctx, rules ...string) error

func RegisterAuthFunc(app *fiber.App, f AuthFunc) {
	
	app.Get("/api/v1/test0", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		rules := []string{
			"admin:write", "admin:read", 
		}
		if err := f(c, rules...); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return c.Next()
	})
	app.Post("/api/v1/test0", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/test2", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		rules := []string{
			"admin:read", 
		}
		if err := f(c, rules...); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return c.Next()
	})
	app.Get("/api/v1/user/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		rules := []string{
			"user:read", 
		}
		if err := f(c, rules...); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return c.Next()
	})
}
