package httpserver_test

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
	"github.com/alvin/mvps/internal/testsupport"
)

func TestPOSTCases_ValidTransferCredit_Returns201AndSubmittedCase(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	server := newTestServer()

	body, err := json.Marshal(scenario.CreateInput())
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/cases", bytes.NewReader(body))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", response.Code)
	}

	var payload struct {
		CaseID      string `json:"case_id"`
		CaseNumber  string `json:"case_number"`
		Status      string `json:"status"`
		Route       string `json:"route"`
		RouteReason string `json:"route_reason"`
		TimelineURL string `json:"timeline_url"`
	}

	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Route != scenario.ExpectedRoute {
		t.Fatalf("expected route %s, got %s", scenario.ExpectedRoute, payload.Route)
	}

	if payload.RouteReason != scenario.ExpectedRouteReason {
		t.Fatalf("expected route reason %s, got %s", scenario.ExpectedRouteReason, payload.RouteReason)
	}
}

func TestPOSTCases_MissingPriorInstitution_Returns400WithFieldError(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-missing-prior-institution.json")
	server := newTestServer()

	body, err := json.Marshal(scenario.CreateInput())
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/cases", bytes.NewReader(body))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", response.Code)
	}

	var payload struct {
		Error struct {
			Code   string `json:"code"`
			Fields []struct {
				Name string `json:"name"`
			} `json:"fields"`
		} `json:"error"`
	}

	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Error.Code != "validation_error" {
		t.Fatalf("expected validation_error, got %s", payload.Error.Code)
	}

	if len(payload.Error.Fields) != 1 || payload.Error.Fields[0].Name != "prior_institution" {
		t.Fatalf("expected prior_institution field error, got %+v", payload.Error.Fields)
	}
}

func TestGETTimeline_ForNewTransferCreditCase_ReturnsSubmittedEvent(t *testing.T) {
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

	server := httpserver.NewHandler(service, repo)
	request := httptest.NewRequest(http.MethodGet, "/cases/"+created.ID+"/timeline", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}

	var payload struct {
		CaseID string `json:"case_id"`
		Events []struct {
			EventType string `json:"event_type"`
			Status    string `json:"status"`
		} `json:"events"`
	}

	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.CaseID != created.ID {
		t.Fatalf("expected case id %s, got %s", created.ID, payload.CaseID)
	}

	if len(payload.Events) != 1 || payload.Events[0].EventType != cases.EventTypeSubmitted {
		t.Fatalf("expected submitted event, got %+v", payload.Events)
	}
}

func TestPOSTCases_InvalidJSON_Returns400(t *testing.T) {
	t.Parallel()

	server := newTestServer()
	request := httptest.NewRequest(http.MethodPost, "/cases", bytes.NewBufferString("{"))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", response.Code)
	}
}

func TestGETTimeline_ForUnknownCase_Returns404(t *testing.T) {
	t.Parallel()

	server := newTestServer()
	request := httptest.NewRequest(http.MethodGet, "/cases/missing/timeline", nil)
	request.Header.Set("X-Correlation-Id", "req-test")
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", response.Code)
	}

	var payload struct {
		Error struct {
			Code          string `json:"code"`
			CorrelationID string `json:"correlation_id"`
		} `json:"error"`
	}

	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Error.Code != "not_found" {
		t.Fatalf("expected not_found, got %s", payload.Error.Code)
	}

	if payload.Error.CorrelationID != "req-test" {
		t.Fatalf("expected req-test correlation id, got %s", payload.Error.CorrelationID)
	}
}

func TestHandler_UnsupportedMethod_Returns404(t *testing.T) {
	t.Parallel()

	server := newTestServer()
	request := httptest.NewRequest(http.MethodGet, "/cases", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", response.Code)
	}
}

func newTestServer() http.Handler {
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo, cases.WithClock(func() time.Time {
		return time.Date(2026, time.May, 5, 10, 0, 0, 0, time.UTC)
	}))

	return httpserver.NewHandler(service, repo)
}
