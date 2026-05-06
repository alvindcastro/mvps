package testsupport

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alvin/mvps/internal/cases"
)

type Scenario struct {
	LearnerRef                  string            `json:"learner_ref"`
	Term                        string            `json:"term"`
	FormType                    string            `json:"form_type"`
	Fields                      map[string]string `json:"fields"`
	ExpectedRoute               string            `json:"expected_route"`
	ExpectedRouteReason         string            `json:"expected_route_reason"`
	ExpectedMissingFields       []string          `json:"expected_missing_fields"`
	ExpectedTimelineEvents      []string          `json:"expected_timeline_events"`
	ExpectedRequiresHumanReview bool              `json:"expected_requires_human_review"`
	ExpectedGuardrailOutcome    string            `json:"expected_guardrail_outcome"`
	ExpectedReviewerExplanation string            `json:"expected_reviewer_explanation"`
}

func LoadScenario(t *testing.T, filename string) Scenario {
	t.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve scenario path")
	}

	path := filepath.Join(filepath.Dir(currentFile), "..", "..", "testdata", "scenarios", filename)
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read scenario %s: %v", filename, err)
	}

	var scenario Scenario
	if err := json.Unmarshal(content, &scenario); err != nil {
		t.Fatalf("decode scenario %s: %v", filename, err)
	}

	return scenario
}

func (s Scenario) CreateInput() cases.CreateInput {
	return cases.CreateInput{
		FormType:   s.FormType,
		LearnerRef: s.LearnerRef,
		Term:       s.Term,
		Fields:     cloneFields(s.Fields),
	}
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
