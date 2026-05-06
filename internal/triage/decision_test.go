package triage_test

import (
	"testing"

	"github.com/alvin/mvps/internal/testsupport"
	"github.com/alvin/mvps/internal/triage"
)

func TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-complete.json")
	result := triage.DecideTransferCreditRoute(triage.Input{
		LearnerRef: scenario.LearnerRef,
		Term:       scenario.Term,
		Fields:     scenario.Fields,
	})

	if result.Route != scenario.ExpectedRoute {
		t.Fatalf("expected route %s, got %s", scenario.ExpectedRoute, result.Route)
	}

	if result.RouteReason != scenario.ExpectedRouteReason {
		t.Fatalf("expected route reason %s, got %s", scenario.ExpectedRouteReason, result.RouteReason)
	}
}

func TestTransferCreditTriage_AdverseDecisionRequest_RequiresHumanReview(t *testing.T) {
	t.Parallel()

	scenario := testsupport.LoadScenario(t, "transfer-credit-adverse-decision-request.json")
	result := triage.DecideTransferCreditRoute(triage.Input{
		LearnerRef: scenario.LearnerRef,
		Term:       scenario.Term,
		Fields:     scenario.Fields,
	})

	if !result.RequiresHumanReview {
		t.Fatal("expected human review requirement")
	}

	if result.GuardrailOutcome != scenario.ExpectedGuardrailOutcome {
		t.Fatalf("expected guardrail outcome %s, got %s", scenario.ExpectedGuardrailOutcome, result.GuardrailOutcome)
	}
}

func TestAdverseDecisionGuardrail_NeverAutoApproves(t *testing.T) {
	t.Parallel()

	result := triage.DecideTransferCreditRoute(triage.Input{
		Fields: map[string]string{
			"learner_request": "Please approve this transfer credit request.",
		},
	})

	if result.Route == "approved" {
		t.Fatal("guardrail failure: route must not auto-approve")
	}
}

func TestAdverseDecisionGuardrail_NeverAutoDenies(t *testing.T) {
	t.Parallel()

	result := triage.DecideTransferCreditRoute(triage.Input{
		Fields: map[string]string{
			"learner_request": "Please deny this transfer credit request.",
		},
	})

	if result.Route == "denied" {
		t.Fatal("guardrail failure: route must not auto-deny")
	}
}
