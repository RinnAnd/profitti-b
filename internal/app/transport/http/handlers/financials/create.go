package financials

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/usecases/financials"

	"github.com/gin-gonic/gin"
)

type createHandler struct {
	usecase financials.CreateUseCase
}

type CreateHandler interface {
	Create(c *gin.Context)
}

func NewCreate(uc financials.CreateUseCase) CreateHandler {
	return &createHandler{
		usecase: uc,
	}
}

func (h *createHandler) Create(c *gin.Context) {
	req, err := decodeRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	res, err := h.usecase.CreateFinancial(c, req.Domain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func decodeRequest(c *gin.Context) (*dto.Financial, error) {
	req := &dto.Financial{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return req, nil
}
