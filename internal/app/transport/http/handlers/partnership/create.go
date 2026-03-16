package partnership

import (
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/usecases/partnership"

	"github.com/gin-gonic/gin"
)

type CreateHandler interface {
	Create(c *gin.Context)
}

type createhandler struct {
	useCase partnership.UseCase
}

func NewCreateHandler(useCase partnership.UseCase) CreateHandler {
	return &createhandler{
		useCase: useCase,
	}
}

func (ch *createhandler) Create(c *gin.Context) {
	req, err := decodeRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := ch.useCase.Create(c, req.Domain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.PartnershipRs{
		Id:      res,
		Message: "partnership created",
	})
}

func decodeRequest(c *gin.Context) (*dto.PartnershipRq, error) {
	req := &dto.PartnershipRq{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}

	return req, nil
}
