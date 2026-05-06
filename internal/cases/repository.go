package cases

import (
	"context"
	"errors"
	"sort"
	"sync"
)

var ErrCaseNotFound = errors.New("case not found")

type Repository interface {
	Create(ctx context.Context, created Case, events []TimelineEvent) error
	GetByID(ctx context.Context, caseID string) (Case, error)
	ListTimeline(ctx context.Context, caseID string) ([]TimelineEvent, error)
	ListSubmittedByRoute(ctx context.Context, route string) ([]Case, error)
}

type InMemoryRepository struct {
	mu        sync.RWMutex
	cases     map[string]Case
	timelines map[string][]TimelineEvent
	order     []string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		cases:     make(map[string]Case),
		timelines: make(map[string][]TimelineEvent),
	}
}

func (r *InMemoryRepository) Create(_ context.Context, created Case, events []TimelineEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	storedCase := created
	storedCase.RequestFields = cloneFields(created.RequestFields)
	r.cases[created.ID] = storedCase

	storedEvents := make([]TimelineEvent, len(events))
	copy(storedEvents, events)
	r.timelines[created.ID] = storedEvents
	r.order = append(r.order, created.ID)

	return nil
}

func (r *InMemoryRepository) GetByID(_ context.Context, caseID string) (Case, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	created, ok := r.cases[caseID]
	if !ok {
		return Case{}, ErrCaseNotFound
	}

	created.RequestFields = cloneFields(created.RequestFields)
	return created, nil
}

func (r *InMemoryRepository) ListTimeline(_ context.Context, caseID string) ([]TimelineEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	events, ok := r.timelines[caseID]
	if !ok {
		if _, exists := r.cases[caseID]; !exists {
			return nil, ErrCaseNotFound
		}
		return []TimelineEvent{}, nil
	}

	cloned := make([]TimelineEvent, len(events))
	copy(cloned, events)
	return cloned, nil
}

func (r *InMemoryRepository) ListSubmittedByRoute(_ context.Context, route string) ([]Case, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]Case, 0, len(r.order))
	for _, caseID := range r.order {
		item := r.cases[caseID]
		if item.Status == StatusSubmitted && item.Route == route {
			item.RequestFields = cloneFields(item.RequestFields)
			items = append(items, item)
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].SubmittedAt.Before(items[j].SubmittedAt)
	})

	return items, nil
}
