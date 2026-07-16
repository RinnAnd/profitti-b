package util

import (
	"profitti/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) (string, error) {
	id, ok := c.Get("id")
	if !ok {
		return "", domain.Token409
	}

	return id.(string), nil
}
