package triage

import "strings"

const (
	RouteRegistrarTransferCredit = "registrar_transfer_credit"
	RouteHumanReviewRequired     = "human_review_required"
	RouteReasonTransferComplete  = "transfer_credit_complete"
	RouteReasonAdverseDecision   = "adverse_decision_request"
)

type Input struct {
	LearnerRef string
	Term       string
	Fields     map[string]string
}

type Result struct {
	Route               string
	RouteReason         string
	RequiresHumanReview bool
	ReviewerExplanation string
	GuardrailOutcome    string
}

func DecideTransferCreditRoute(input Input) Result {
	requestedOutcome := strings.ToLower(strings.TrimSpace(input.Fields["requested_outcome"]))
	if requestedOutcome == "" {
		requestedOutcome = strings.ToLower(strings.TrimSpace(input.Fields["learner_request"]))
	}

	if strings.Contains(requestedOutcome, "approve") ||
		strings.Contains(requestedOutcome, "deny") ||
		strings.Contains(requestedOutcome, "denial") ||
		strings.Contains(requestedOutcome, "reject") {
		return Result{
			Route:               RouteHumanReviewRequired,
			RouteReason:         RouteReasonAdverseDecision,
			RequiresHumanReview: true,
			ReviewerExplanation: "Human review required because approval or denial cannot be automated.",
			GuardrailOutcome:    RouteHumanReviewRequired,
		}
	}

	return Result{
		Route:               RouteRegistrarTransferCredit,
		RouteReason:         RouteReasonTransferComplete,
		RequiresHumanReview: false,
		ReviewerExplanation: "Transfer-credit request is complete and ready for reviewer triage.",
		GuardrailOutcome:    "",
	}
}
