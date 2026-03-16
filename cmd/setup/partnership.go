package setup

import (
	"database/sql"
	handler "profitti/internal/app/transport/http/handlers/partnership"
	usecase "profitti/internal/core/usecases/partnership"
	"profitti/internal/infra/database/repository/partnership"
	service "profitti/internal/infra/service/partnership"
)

func Partnership(db *sql.DB, stp *Setup) {
	partnershipRepo := partnership.New(db)
	partnershipSrv := service.New(partnershipRepo)

	partnershipUsecase := usecase.New(partnershipSrv)
	partnershipHandler := handler.NewCreateHandler(partnershipUsecase)
	getPartnershipHandler := handler.NewGet(partnershipUsecase)

	stp.CreatePartnership = partnershipHandler
	stp.GetPartnerships = getPartnershipHandler
}
