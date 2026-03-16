package partnership

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/domain"
	"profitti/internal/core/usecases/partnership"

	"github.com/gin-gonic/gin"
)

type GetHandler interface {
	GetPartnerships(c *gin.Context)
}

type gethandler struct {
	useCase partnership.UseCase
}

func NewGet(useCase partnership.UseCase) GetHandler {
	return &gethandler{
		useCase: useCase,
	}
}

func (g *gethandler) GetPartnerships(c *gin.Context) {
	id := c.Param("id")

	res, err := g.useCase.GetByUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resMapper(res))
}

func resMapper(r []*domain.Partnership) *dto.PartnershipsRs {
	partnerships := []dto.Partnership{}
	for _, p := range r {
		partnership := dto.Partnership{
			Id:         p.Id,
			Users:      p.Users,
			CurrencyId: p.CurrencyId,
		}

		partnerships = append(partnerships, partnership)
	}

	return &dto.PartnershipsRs{
		Partnerships: partnerships,
	}
}
