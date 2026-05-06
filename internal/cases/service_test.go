package cases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alvin/mvps/internal/cases"
	"github.com/alvin/mvps/internal/testsupport"
)

func TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-missing-prior-institution.json")
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(fixedClock))

	_, err := service.Create(context.Background(), scenario.CreateInput())
	if err == nil {
		t.Fatal("expected validation error")
	}

	var validationErr cases.ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected cases.ValidationError, got %T", err)
	}

	if len(validationErr.Fields) != 1 {
		t.Fatalf("expected 1 field error, got %d", len(validationErr.Fields))
	}

	if validationErr.Fields[0].Name != "prior_institution" {
		t.Fatalf("expected prior_institution field error, got %s", validationErr.Fields[0].Name)
	}
}

func TestCreateTransferCreditCase_WithCompleteInput_CreatesSubmittedCase(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(fixedClock))

	created, err := service.Create(context.Background(), scenario.CreateInput())
	if err != nil {
		t.Fatalf("create case: %v", err)
	}

	if created.ID != "case-0001" {
		t.Fatalf("expected case-0001, got %s", created.ID)
	}

	if created.CaseNumber != "TC-2026-0001" {
		t.Fatalf("expected TC-2026-0001, got %s", created.CaseNumber)
	}

	if created.Status != cases.StatusSubmitted {
		t.Fatalf("expected status submitted, got %s", created.Status)
	}

	if created.Route != scenario.ExpectedRoute {
		t.Fatalf("expected route %s, got %s", scenario.ExpectedRoute, created.Route)
	}

	if created.RouteReason != scenario.ExpectedRouteReason {
		t.Fatalf("expected route reason %s, got %s", scenario.ExpectedRouteReason, created.RouteReason)
	}

	events, err := repo.ListTimeline(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("list timeline: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 timeline event, got %d", len(events))
	}

	if events[0].EventType != cases.EventTypeSubmitted {
		t.Fatalf("expected submitted event, got %s", events[0].EventType)
	}
}

func TestCreateTransferCreditCase_WithUnsupportedFormType_ReturnsValidationError(t *testing.T) {
	t.Parallel()

	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(fixedClock))

	_, err := service.Create(context.Background(), cases.CreateInput{
		FormType:   "refund_withdrawal",
		LearnerRef: "STU-300900111",
		Term:       "Fall 2026",
		Fields: map[string]string{
			"prior_institution": "Example University",
			"prior_course_code": "MGMT 101",
			"target_program":    "Business Administration Diploma",
		},
	})
	if err == nil {
		t.Fatal("expected validation error")
	}

	var validationErr cases.ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected validation error, got %T", err)
	}
}

func TestService_Timeline_ReturnsStoredEvents(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(fixedClock), cases.WithStartingSequence(5))

	created, err := service.Create(context.Background(), scenario.CreateInput())
	if err != nil {
		t.Fatalf("create case: %v", err)
	}

	events, err := service.Timeline(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("timeline: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 timeline event, got %d", len(events))
	}

	if created.ID != "case-0006" {
		t.Fatalf("expected case-0006, got %s", created.ID)
	}
}

func fixedClock() time.Time {
	return time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)
}
