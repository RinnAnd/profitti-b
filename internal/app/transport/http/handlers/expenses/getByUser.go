package expenses

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/usecases/expenses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetByUserHandler interface {
	GetByUser(c *gin.Context)
}

type getByUserUseCase struct {
	u expenses.GetByUserUseCase
}

func NewGetByUser(u expenses.GetByUserUseCase) GetByUserHandler {
	return &getByUserUseCase{
		u: u,
	}
}

func (h *getByUserUseCase) GetByUser(c *gin.Context) {
	id := c.Param("id")
	if !isValidUUID(id) {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: "Invalid user id",
		})
	}
	res, err := h.u.GetExpensesByUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, dto.GetExpensesByUserRes{
		Expenses: res,
	})
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
