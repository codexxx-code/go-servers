package http

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Proxy(c *fiber.Ctx) error {
	link := c.Query("link")
	if link == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "link parameter is required",
		})
	}

	bodyBytes := c.Body()
	req, err := http.NewRequestWithContext(context.Background(), c.Method(), link, bytes.NewReader(bodyBytes))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Forward headers from the original request
	for key, values := range c.GetReqHeaders() {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Use io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(resp.StatusCode).Send(body)
}
