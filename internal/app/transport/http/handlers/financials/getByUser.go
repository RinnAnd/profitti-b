package financials

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/usecases/financials"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type getByUserHandler struct {
	usecase financials.GetUserFinancialsUseCase
}

type GetByUserHandler interface {
	GetByUser(c *gin.Context)
}

func NewGetByUser(uc financials.GetUserFinancialsUseCase) GetByUserHandler {
	return &getByUserHandler{
		usecase: uc,
	}
}

func (h *getByUserHandler) GetByUser(c *gin.Context) {
	id := c.Param("id")
	if !isValidUUID(id) {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: "Invalid user id",
		})
	}
	res, err := h.usecase.GetFinancialsByUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusCreated, res)
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
