package categories

import (
	"context"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/app/util"
	uc "profitti/internal/core/usecases/categories"

	"github.com/gin-gonic/gin"
)

type GetHandler interface {
	Get(c *gin.Context)
}

type ghandler struct {
	usecase uc.GUseCase
}

func NewG(usecase uc.GUseCase) GetHandler {
	return &ghandler{
		usecase: usecase,
	}
}

func (g *ghandler) Get(c *gin.Context) {
	ctx := context.Background()

	id, err := util.GetId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.HttpError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	res, err := g.usecase.GetUserCategories(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	cats := []dto.Category{}

	for _, cat := range res {
		dtcat := dto.Category{
			Id:   cat.Id,
			Name: cat.Name,
		}
		cats = append(cats, dtcat)
	}

	c.JSON(http.StatusOK, dto.CategoryRes{
		UserId:     id,
		Categories: cats,
	})
}
