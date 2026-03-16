package routes

import (
	"profitti/internal/app/transport/http/handlers/expenses"
	"profitti/internal/app/transport/http/handlers/financials"
	"profitti/internal/app/transport/http/handlers/users"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	// Users
	RegisterHandler users.RegisterHandler
	LoginHandler    users.LoginHandler
	// Financials
	CreateFinancialHandler     financials.CreateHandler
	GetFinancialsByUserHandler financials.GetByUserHandler
	// Expenses
	CreateExpenseHandler     expenses.CreateHandler
	GetExpensesByUserHandler expenses.GetByUserHandler
}

func (h *Routes) Init(s *gin.Engine) {
	// Users
	userGroup := s.Group("/users")
	userGroup.POST("/register", h.RegisterHandler.Register)
	userGroup.POST("/login", h.LoginHandler.Login)
	// Financials
	financialGroup := s.Group("/financials")
	financialGroup.POST("/create", h.CreateFinancialHandler.Create)
	financialGroup.GET("/user/:id", h.GetFinancialsByUserHandler.GetByUser)
	// Expenses
	expenseGroup := s.Group("/expenses")
	expenseGroup.POST("/create", h.CreateExpenseHandler.Create)
	expenseGroup.GET("/user/:id", h.GetExpensesByUserHandler.GetByUser)
}
