package handler

import (
	"net/http"
	"urlShortnerer/cmd/model"

	"github.com/labstack/echo/v5"
)

type CreateLinkRequest struct {
}

func HealthHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "handler and router work!"})
}

func CreateLinkHandler(c *echo.Context) error {
	var Link model.Link

	if err := c.Bind(&Link); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Link created successfully!"})
}

func GetLinkHandler(c *echo.Context) error {
	// Implement the logic to get a link here
	return c.JSON(http.StatusOK, map[string]string{"message": "Link retrieved successfully!"})
}
