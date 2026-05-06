package reviewerqueue_test

import (
	"context"
	"testing"
	"time"

	"github.com/alvin/mvps/internal/cases"
	"github.com/alvin/mvps/internal/reviewerqueue"
	"github.com/alvin/mvps/internal/testsupport"
)

func TestReviewerQueue_NewTransferCreditCase_IsVisibleToStaff(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(func() time.Time {
		return time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)
	}))

	created, err := service.Create(context.Background(), scenario.CreateInput())
	if err != nil {
		t.Fatalf("create case: %v", err)
	}

	items, err := reviewerqueue.ListSubmittedTransferCredit(context.Background(), repo)
	if err != nil {
		t.Fatalf("list queue: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 queue item, got %d", len(items))
	}

	if items[0].CaseID != created.ID {
		t.Fatalf("expected case id %s, got %s", created.ID, items[0].CaseID)
	}
}
