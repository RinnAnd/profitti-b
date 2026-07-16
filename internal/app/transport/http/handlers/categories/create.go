package categories

import (
	"context"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/app/util"
	uc "profitti/internal/core/usecases/categories"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Post(c *gin.Context)
}

type handler struct {
	usecase uc.CUseCase
}

func New(usecase uc.CUseCase) Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) Post(c *gin.Context) {
	ctx := context.Background()

	cat, err := decode(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	id, err := util.GetId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.HttpError{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	dom := cat.Domain()
	dom.UserId = id

	if err = h.usecase.CreateCategory(ctx, dom); err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: "category created successfully",
	})
}

func decode(c *gin.Context) (*dto.Category, error) {
	cat := &dto.Category{}

	if err := c.ShouldBindBodyWithJSON(cat); err != nil {
		return nil, err
	}

	return cat, nil
}
