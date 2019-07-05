package usecase

import (
	"context"

	"github.com/strwrd/rest-scrapper/model"
)

// GetJournalsByArchieveID
func (u *usecase) GetJournalsByArchieveID(ctx context.Context, ID string) ([]*model.Journal, error) {
	// Retrieve data from DB
	journals, err := u.repo.GetJournalsByArchieveID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return journals, nil
}
