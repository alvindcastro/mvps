//go:build e2e

package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alvin/mvps/internal/cases"
	"github.com/alvin/mvps/internal/platform/httpserver"
	"github.com/alvin/mvps/internal/reviewerqueue"
	"github.com/alvin/mvps/internal/testsupport"
)

func TestE2ETransferCreditSubmission_AppearsInReviewerQueue(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(func() time.Time {
		return time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)
	}))
	server := httpserver.NewHandler(service, repo)

	body, err := json.Marshal(scenario.CreateInput())
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	postRequest := httptest.NewRequest(http.MethodPost, "/cases", bytes.NewReader(body))
	postResponse := httptest.NewRecorder()
	server.ServeHTTP(postResponse, postRequest)

	if postResponse.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", postResponse.Code)
	}

	var created struct {
		CaseID string `json:"case_id"`
	}
	if err := json.Unmarshal(postResponse.Body.Bytes(), &created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}

	timelineRequest := httptest.NewRequest(http.MethodGet, "/cases/"+created.CaseID+"/timeline", nil)
	timelineResponse := httptest.NewRecorder()
	server.ServeHTTP(timelineResponse, timelineRequest)

	if timelineResponse.Code != http.StatusOK {
		t.Fatalf("expected timeline 200, got %d", timelineResponse.Code)
	}

	items, err := reviewerqueue.ListSubmittedTransferCredit(context.Background(), repo)
	if err != nil {
		t.Fatalf("query reviewer queue: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 queue item, got %d", len(items))
	}

	if items[0].CaseID != created.CaseID {
		t.Fatalf("expected queue case id %s, got %s", created.CaseID, items[0].CaseID)
	}
}
