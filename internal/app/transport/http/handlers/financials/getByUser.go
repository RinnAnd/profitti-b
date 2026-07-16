package financials

import (
	"database/sql"
	"errors"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/app/util"
	"profitti/internal/core/domain"
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
	id, err := util.GetId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.HttpError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	res, err := h.usecase.GetFinancialsByUser(c, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusOK, dto.FinancialRes{
				User:       id,
				Financials: []dto.Financial{},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.FinancialRes{
		User:       id,
		Financials: domainToDTO(res),
	})
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func domainToDTO(res []*domain.Financial) []dto.Financial {
	result := []dto.Financial{}
	for _, fin := range res {
		fn := dto.Financial{
			Id:         fin.Id,
			UserId:     fin.UserId,
			CurrencyId: fin.CurrencyId,
		}

		result = append(result, fn)
	}

	return result
}
