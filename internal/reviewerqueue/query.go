package reviewerqueue

import (
	"context"
	"time"

	"github.com/alvin/mvps/internal/cases"
)

type Item struct {
	CaseID      string    `json:"case_id"`
	CaseNumber  string    `json:"case_number"`
	LearnerRef  string    `json:"learner_ref"`
	Status      string    `json:"status"`
	Route       string    `json:"route"`
	SubmittedAt time.Time `json:"submitted_at"`
}

type Source interface {
	ListSubmittedByRoute(ctx context.Context, route string) ([]cases.Case, error)
}

func ListSubmittedTransferCredit(ctx context.Context, source Source) ([]Item, error) {
	found, err := source.ListSubmittedByRoute(ctx, cases.RouteRegistrarTransferCredit)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0, len(found))
	for _, created := range found {
		items = append(items, Item{
			CaseID:      created.ID,
			CaseNumber:  created.CaseNumber,
			LearnerRef:  created.LearnerRef,
			Status:      created.Status,
			Route:       created.Route,
			SubmittedAt: created.SubmittedAt,
		})
	}

	return items, nil
}
