package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/alvin/mvps/internal/cases"
)

type Handler struct {
	service *cases.Service
	repo    cases.Repository
	reqID   atomic.Int64
}

func NewHandler(service *cases.Service, repo cases.Repository) http.Handler {
	handler := &Handler{
		service: service,
		repo:    repo,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cases", handler.handleCases)
	mux.HandleFunc("/cases/", handler.handleCaseSubresources)
	return mux
}

func (h *Handler) handleCases(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.NotFound(writer, request)
		return
	}

	var payload cases.CreateInput
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		h.writeError(writer, request, http.StatusBadRequest, "invalid_json", "The request body must be valid JSON.", nil)
		return
	}

	created, err := h.service.Create(request.Context(), payload)
	if err != nil {
		var validationErr cases.ValidationError
		if errors.As(err, &validationErr) {
			h.writeError(writer, request, http.StatusBadRequest, "validation_error", "The request contains invalid fields.", validationErr.Fields)
			return
		}

		h.writeError(writer, request, http.StatusInternalServerError, "internal_error", "The request could not be completed.", nil)
		return
	}

	response := struct {
		CaseID      string `json:"case_id"`
		CaseNumber  string `json:"case_number"`
		Status      string `json:"status"`
		Route       string `json:"route"`
		RouteReason string `json:"route_reason"`
		TimelineURL string `json:"timeline_url"`
	}{
		CaseID:      created.ID,
		CaseNumber:  created.CaseNumber,
		Status:      created.Status,
		Route:       created.Route,
		RouteReason: created.RouteReason,
		TimelineURL: "/cases/" + created.ID + "/timeline",
	}

	h.writeJSON(writer, http.StatusCreated, response)
}

func (h *Handler) handleCaseSubresources(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.NotFound(writer, request)
		return
	}

	parts := strings.Split(strings.Trim(request.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "cases" || parts[2] != "timeline" {
		http.NotFound(writer, request)
		return
	}

	events, err := h.repo.ListTimeline(request.Context(), parts[1])
	if err != nil {
		if errors.Is(err, cases.ErrCaseNotFound) {
			h.writeError(writer, request, http.StatusNotFound, "not_found", "The requested case was not found.", nil)
			return
		}

		h.writeError(writer, request, http.StatusInternalServerError, "internal_error", "The request could not be completed.", nil)
		return
	}

	response := struct {
		CaseID string                `json:"case_id"`
		Events []cases.TimelineEvent `json:"events"`
	}{
		CaseID: parts[1],
		Events: events,
	}

	h.writeJSON(writer, http.StatusOK, response)
}

func (h *Handler) writeError(writer http.ResponseWriter, request *http.Request, statusCode int, code, message string, fields []cases.FieldError) {
	response := struct {
		Error struct {
			Code          string             `json:"code"`
			Message       string             `json:"message"`
			Fields        []cases.FieldError `json:"fields,omitempty"`
			CorrelationID string             `json:"correlation_id"`
		} `json:"error"`
	}{}

	response.Error.Code = code
	response.Error.Message = message
	response.Error.Fields = fields
	response.Error.CorrelationID = h.correlationID(request)

	h.writeJSON(writer, statusCode, response)
}

func (h *Handler) writeJSON(writer http.ResponseWriter, statusCode int, payload any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_ = json.NewEncoder(writer).Encode(payload)
}

func (h *Handler) correlationID(request *http.Request) string {
	if value := strings.TrimSpace(request.Header.Get("X-Correlation-Id")); value != "" {
		return value
	}

	return "req-" + strconv.FormatInt(h.reqID.Add(1), 10)
}
