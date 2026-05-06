package cases

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alvin/mvps/internal/triage"
)

type Service struct {
	repo     Repository
	clock    func() time.Time
	mu       sync.Mutex
	sequence int
}

type ServiceOption func(*Service)

func WithClock(clock func() time.Time) ServiceOption {
	return func(service *Service) {
		service.clock = clock
	}
}

func WithStartingSequence(sequence int) ServiceOption {
	return func(service *Service) {
		service.sequence = sequence
	}
}

func NewService(repo Repository, options ...ServiceOption) *Service {
	service := &Service{
		repo:  repo,
		clock: time.Now,
	}

	for _, option := range options {
		option(service)
	}

	return service
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Case, error) {
	if err := ValidateCreateInput(input); err != nil {
		return Case{}, err
	}

	decision := triage.DecideTransferCreditRoute(triage.Input{
		LearnerRef: input.LearnerRef,
		Term:       input.Term,
		Fields:     cloneFields(input.Fields),
	})

	now := s.clock().UTC()
	caseID, caseNumber := s.nextIdentifiers(now)
	created := Case{
		ID:            caseID,
		CaseNumber:    caseNumber,
		LearnerRef:    input.LearnerRef,
		FormType:      input.FormType,
		Term:          input.Term,
		Status:        StatusSubmitted,
		Route:         decision.Route,
		RouteReason:   decision.RouteReason,
		RequestFields: cloneFields(input.Fields),
		SubmittedAt:   now,
		CreatedAt:     now,
	}

	events := []TimelineEvent{
		{
			ID:        fmt.Sprintf("%s-event-1", caseID),
			CaseID:    caseID,
			EventType: EventTypeSubmitted,
			Status:    StatusSubmitted,
			Message:   SubmittedTimelineEventMessage,
			CreatedAt: now,
		},
	}

	if err := s.repo.Create(ctx, created, events); err != nil {
		return Case{}, err
	}

	return created, nil
}

func (s *Service) Timeline(ctx context.Context, caseID string) ([]TimelineEvent, error) {
	return s.repo.ListTimeline(ctx, caseID)
}

func (s *Service) nextIdentifiers(now time.Time) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sequence++
	return fmt.Sprintf("case-%04d", s.sequence), fmt.Sprintf("TC-%d-%04d", now.Year(), s.sequence)
}
