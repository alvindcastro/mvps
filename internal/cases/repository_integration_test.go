package cases_test

import (
	"context"
	"testing"
	"time"

	"github.com/alvin/mvps/internal/cases"
)

func TestCaseRepository_CreateTransferCreditCase_PersistsSubmittedCase(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()
	createdAt := time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)
	created := cases.Case{
		ID:          "case-0001",
		CaseNumber:  "TC-2026-0001",
		LearnerRef:  "STU-300900111",
		FormType:    cases.FormTypeTransferCredit,
		Status:      cases.StatusSubmitted,
		Route:       cases.RouteRegistrarTransferCredit,
		RouteReason: cases.RouteReasonTransferComplete,
		RequestFields: map[string]string{
			"prior_institution": "Example University",
		},
		SubmittedAt: createdAt,
		CreatedAt:   createdAt,
	}

	event := cases.TimelineEvent{
		ID:        "case-0001-event-1",
		CaseID:    "case-0001",
		EventType: cases.EventTypeSubmitted,
		Status:    cases.StatusSubmitted,
		Message:   cases.SubmittedTimelineEventMessage,
		CreatedAt: createdAt,
	}

	if err := repo.Create(context.Background(), created, []cases.TimelineEvent{event}); err != nil {
		t.Fatalf("create case: %v", err)
	}

	stored, err := repo.GetByID(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("get by id: %v", err)
	}

	if stored.CaseNumber != created.CaseNumber {
		t.Fatalf("expected case number %s, got %s", created.CaseNumber, stored.CaseNumber)
	}

	if stored.Route != cases.RouteRegistrarTransferCredit {
		t.Fatalf("expected registrar route, got %s", stored.Route)
	}
}

func TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()
	createdAt := time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)

	if err := repo.Create(context.Background(), cases.Case{
		ID:          "case-0001",
		CaseNumber:  "TC-2026-0001",
		LearnerRef:  "STU-300900111",
		FormType:    cases.FormTypeTransferCredit,
		Status:      cases.StatusSubmitted,
		Route:       cases.RouteRegistrarTransferCredit,
		RouteReason: cases.RouteReasonTransferComplete,
		SubmittedAt: createdAt,
		CreatedAt:   createdAt,
	}, []cases.TimelineEvent{{
		ID:        "case-0001-event-1",
		CaseID:    "case-0001",
		EventType: cases.EventTypeSubmitted,
		Status:    cases.StatusSubmitted,
		Message:   cases.SubmittedTimelineEventMessage,
		CreatedAt: createdAt,
	}}); err != nil {
		t.Fatalf("create case: %v", err)
	}

	timeline, err := repo.ListTimeline(context.Background(), "case-0001")
	if err != nil {
		t.Fatalf("list timeline: %v", err)
	}

	if len(timeline) != 1 {
		t.Fatalf("expected 1 timeline event, got %d", len(timeline))
	}

	if timeline[0].Message != cases.SubmittedTimelineEventMessage {
		t.Fatalf("expected submitted message, got %s", timeline[0].Message)
	}
}

func TestInMemoryRepository_GetByID_ReturnsNotFound(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()

	_, err := repo.GetByID(context.Background(), "missing")
	if err != cases.ErrCaseNotFound {
		t.Fatalf("expected ErrCaseNotFound, got %v", err)
	}
}

func TestInMemoryRepository_ListTimeline_ReturnsNotFound(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()

	_, err := repo.ListTimeline(context.Background(), "missing")
	if err != cases.ErrCaseNotFound {
		t.Fatalf("expected ErrCaseNotFound, got %v", err)
	}
}

func TestInMemoryRepository_ListSubmittedByRoute_FiltersSubmittedCases(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()
	createdAt := time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)

	mustCreateCase(t, repo, cases.Case{
		ID:          "case-0001",
		CaseNumber:  "TC-2026-0001",
		LearnerRef:  "STU-300900111",
		FormType:    cases.FormTypeTransferCredit,
		Status:      cases.StatusSubmitted,
		Route:       cases.RouteRegistrarTransferCredit,
		RouteReason: cases.RouteReasonTransferComplete,
		SubmittedAt: createdAt,
		CreatedAt:   createdAt,
	})
	mustCreateCase(t, repo, cases.Case{
		ID:          "case-0002",
		CaseNumber:  "TC-2026-0002",
		LearnerRef:  "STU-300900112",
		FormType:    cases.FormTypeTransferCredit,
		Status:      "in_review",
		Route:       cases.RouteRegistrarTransferCredit,
		RouteReason: cases.RouteReasonTransferComplete,
		SubmittedAt: createdAt.Add(time.Minute),
		CreatedAt:   createdAt.Add(time.Minute),
	})

	found, err := repo.ListSubmittedByRoute(context.Background(), cases.RouteRegistrarTransferCredit)
	if err != nil {
		t.Fatalf("list submitted by route: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected 1 submitted case, got %d", len(found))
	}

	if found[0].ID != "case-0001" {
		t.Fatalf("expected case-0001, got %s", found[0].ID)
	}
}

func mustCreateCase(t *testing.T, repo *cases.InMemoryRepository, created cases.Case) {
	t.Helper()

	err := repo.Create(context.Background(), created, []cases.TimelineEvent{{
		ID:        created.ID + "-event-1",
		CaseID:    created.ID,
		EventType: cases.EventTypeSubmitted,
		Status:    cases.StatusSubmitted,
		Message:   cases.SubmittedTimelineEventMessage,
		CreatedAt: created.CreatedAt,
	}})
	if err != nil {
		t.Fatalf("create case: %v", err)
	}
}
