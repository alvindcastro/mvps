package cases

import "time"

const (
	FormTypeTransferCredit        = "transfer_credit"
	StatusSubmitted               = "submitted"
	EventTypeSubmitted            = "submitted"
	RouteRegistrarTransferCredit  = "registrar_transfer_credit"
	RouteHumanReviewRequired      = "human_review_required"
	RouteReasonTransferComplete   = "transfer_credit_complete"
	RouteReasonAdverseDecision    = "adverse_decision_request"
	SubmittedTimelineEventMessage = "Transfer credit request submitted."
)

type Case struct {
	ID            string
	CaseNumber    string
	LearnerRef    string
	FormType      string
	Term          string
	Status        string
	Route         string
	RouteReason   string
	RequestFields map[string]string
	SubmittedAt   time.Time
	CreatedAt     time.Time
}

type TimelineEvent struct {
	ID        string    `json:"id,omitempty"`
	CaseID    string    `json:"case_id,omitempty"`
	EventType string    `json:"event_type"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateInput struct {
	FormType   string            `json:"form_type"`
	LearnerRef string            `json:"learner_ref"`
	Term       string            `json:"term"`
	Fields     map[string]string `json:"fields"`
}

func cloneFields(fields map[string]string) map[string]string {
	if len(fields) == 0 {
		return map[string]string{}
	}

	cloned := make(map[string]string, len(fields))
	for key, value := range fields {
		cloned[key] = value
	}

	return cloned
}
