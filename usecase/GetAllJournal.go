package usecase

import (
	"context"

	"github.com/strwrd/rest-scrapper/model"
)

// GetAllJournal
func (u *usecase) GetAllJournal(ctx context.Context) ([]*model.Journal, error) {
	// Retrieve data from DB
	journals, err := u.repo.GetAllJournal(ctx)
	if err != nil {
		return nil, err
	}

	return journals, nil
}
