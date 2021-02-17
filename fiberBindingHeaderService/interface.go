package fiberBindingHeaderService

import "github.com/gofiber/fiber/v2"

type FiberBindingHeaderInterface interface {
	BindFiberHeader(obj interface{}, c *fiber.Ctx) error
}
