package middleware

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/infra/service/auth"

	"github.com/gin-gonic/gin"
)

const tokenKey = "access_token"

type midd struct {
	jwt auth.JWT
}

func New(jwt auth.JWT) AuthMidd {
	return &midd{
		jwt: jwt,
	}
}

type AuthMidd interface {
	CheckToken() gin.HandlerFunc
}

func (m *midd) CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(tokenKey)

		if token == "" {
			c.Next()
			return
		}

		claims, err := m.jwt.Validate(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.HttpError{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		id, _ := (*claims)["id"]
		c.Set("id", id)
		c.Next()
	}
}
