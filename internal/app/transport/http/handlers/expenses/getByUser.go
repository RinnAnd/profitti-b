package expenses

import (
	"errors"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/app/util"
	"profitti/internal/core/domain"
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
	id, err := util.GetId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.HttpError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	res, err := h.u.GetExpensesByUser(c, id)
	if err != nil {
		if errors.Is(err, domain.User404) {
			c.JSON(http.StatusNotFound, dto.HttpError{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.GetExpensesByUserRes{
		Expenses: func() []*dto.Expense {
			expenses := []*dto.Expense{}
			for _, exp := range res {
				expense := &dto.Expense{
					Id:                 exp.Id,
					FinancialId:        exp.FinancialId,
					PartnershipId:      exp.PartnershipId,
					Name:               exp.Name,
					Description:        exp.Description,
					Amount:             exp.Amount,
					Category:           exp.Category,
					Expense_recurrence: exp.Expense_recurrence,
					Expiration_date:    exp.Expiration_date,
					Currency: dto.Currency{
						Id:       exp.CurrencyId,
						Currency: "COP",
					},
					CreatedAt: exp.CreatedAt,
				}

				expenses = append(expenses, expense)
			}
			return expenses
		}(),
	})
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
