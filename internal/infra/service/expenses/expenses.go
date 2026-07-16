package expenses

import (
	"context"
	"profitti/internal/core/domain"
	expense "profitti/internal/infra/database/repository/expenses"
	"profitti/internal/infra/database/repository/financial"
	"profitti/internal/infra/database/repository/partnership"
	"sync"
)

type ExpenseService interface {
	Create(context.Context, *domain.Expense) (string, error)
	GetUserExpenses(context.Context, string) ([]*domain.Expense, error)
}

type service struct {
	repo        expense.ExpenseRepo
	finrepo     financial.FinancialRepo
	partnership partnership.Partnership
}

func New(repo expense.ExpenseRepo, finrepo financial.FinancialRepo, partnership partnership.Partnership) ExpenseService {
	return &service{
		repo:        repo,
		finrepo:     finrepo,
		partnership: partnership,
	}
}

func (s *service) Create(ctx context.Context, e *domain.Expense) (string, error) {
	res, err := s.repo.InsertOne(ctx, e)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *service) GetUserExpenses(ctx context.Context, id string) ([]*domain.Expense, error) {
	response := []*domain.Expense{}
	idschan := make(chan string)
	errchan := make(chan error, 1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		res, err := s.finrepo.SelectUserFinancials(ctx, id)
		if err != nil {
			errchan <- err
			return
		}

		for _, f := range res {
			idschan <- f.Id
		}
	}()

	go func() {
		defer wg.Done()
		res, err := s.partnership.Select(ctx, id)
		if err != nil {
			errchan <- err
			return
		}

		for _, p := range res {
			idschan <- p.Id
		}
	}()

	go func() {
		wg.Wait()
		close(idschan)
		close(errchan)
	}()

	for fpid := range idschan {
		res, err := s.repo.SelectUserExpenses(ctx, fpid)
		if err != nil {
			return nil, err
		}

		response = append(response, res...)
	}

	if err := <-errchan; err != nil {
		return nil, err
	}

	return response, nil
}
