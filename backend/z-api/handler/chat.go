package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetChats(c *fiber.Ctx) error {
	userInstance := c.Query("userInstance")
	userToken := c.Query("userToken")
	accountToken := c.Query("accountToken")
	chats, err := h.usecase.GetChats(userInstance, userToken, accountToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(chats)
}
