package expenses

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/usecases/expenses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateHandler interface {
	Create(c *gin.Context)
}

type createHandler struct {
	u expenses.CreateUseCase
}

func NewCreate(useCase expenses.CreateUseCase) CreateHandler {
	return &createHandler{
		u: useCase,
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
	_, err = h.u.CreateExpense(c, req.Domain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	createRes := dto.CreateExpenseRes{
		Msg: "expense created succesfully",
	}
	c.JSON(http.StatusCreated, createRes)
}

func decodeRequest(c *gin.Context) (*dto.Expense, error) {
	req := &dto.Expense{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	return req, nil
}
