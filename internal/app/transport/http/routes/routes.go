package routes

import (
	"profitti/internal/app/transport/http/handlers/categories"
	"profitti/internal/app/transport/http/handlers/expenses"
	"profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/app/transport/http/handlers/partnership"
	"profitti/internal/app/transport/http/handlers/users"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	RegisterHandler            users.RegisterHandler
	LoginHandler               users.LoginHandler
	CreateFinancialHandler     financials.CreateHandler
	GetFinancialsByUserHandler financials.GetByUserHandler
	CreateExpenseHandler       expenses.CreateHandler
	GetExpensesByUserHandler   expenses.GetByUserHandler
	CreatePartnership          partnership.CreateHandler
	GetPartnerships            partnership.GetHandler
	CreateCategory             categories.Handler
	GetCategories              categories.GetHandler
}

// ADD FRIENDSHIPS, ADD CHARGE NOTIFICATIONS, MAYBE BUDGETS OR INCOMES

func (h *Routes) Init(s *gin.Engine) {
	s.POST("/users/register", h.RegisterHandler.Register)
	s.POST("/users/login", h.LoginHandler.Login)
	s.POST("/financials/create", h.CreateFinancialHandler.Create)
	s.GET("/financials/user", h.GetFinancialsByUserHandler.GetByUser)
	s.POST("/expenses/create", h.CreateExpenseHandler.Create)
	s.GET("/expenses/user", h.GetExpensesByUserHandler.GetByUser)
	s.POST("/partnership/create", h.CreatePartnership.Create)
	s.GET("/partnership", h.GetPartnerships.GetPartnerships)
	s.POST("/category", h.CreateCategory.Post)
	s.GET("/category", h.GetCategories.Get)
}
