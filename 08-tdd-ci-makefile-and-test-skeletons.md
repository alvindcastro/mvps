# 08 - TDD CI, Makefile, and Go Test Skeletons

## Makefile Targets

```makefile
.PHONY: test test-unit test-integration test-contract test-e2e coverage eval race lint ci

test:
	go test ./...

test-unit:
	go test ./internal/...

test-integration:
	go test -tags=integration ./...

test-contract:
	go test -tags=contract ./...

test-e2e:
	go test -tags=e2e ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

race:
	go test -race ./...

eval:
	go test -tags=evaluation ./internal/evaluation/...

lint:
	golangci-lint run

ci: lint test test-integration test-contract coverage eval
```

## Coverage Gate Script

```bash
#!/usr/bin/env bash
set -euo pipefail

go test ./... -coverprofile=coverage.out
coverage=$(go tool cover -func=coverage.out | awk '/total:/ {print substr($3, 1, length($3)-1)}')

required=85

awk -v coverage="$coverage" -v required="$required" 'BEGIN {
  if (coverage < required) {
    printf("coverage %.2f is below required %.2f\n", coverage, required)
    exit 1
  }
  printf("coverage %.2f meets required %.2f\n", coverage, required)
}'
```

## GitHub Actions Example

```yaml
name: ci

on:
  pull_request:
  push:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest

    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_USER: app
          POSTGRES_PASSWORD: app
          POSTGRES_DB: student_forms_test
        ports:
          - 5432:5432
        options: >-
          --health-cmd "pg_isready -U app -d student_forms_test"
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - name: Download dependencies
        run: go mod download

      - name: Run unit tests
        run: make test-unit

      - name: Run integration tests
        run: make test-integration

      - name: Run contract tests
        run: make test-contract

      - name: Run evaluation gates
        run: make eval

      - name: Run coverage gate
        run: ./scripts/check-coverage.sh
```

---

# Go Test Skeletons

## Case Service Test

```go
package cases_test

import (
    "context"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/cases"
)

func TestCaseService_Create_WithValidTransferCreditInput_PersistsCase(t *testing.T) {
    t.Parallel()

    repo := cases.NewFakeRepository()
    svc := cases.NewService(repo)

    input := cases.CreateInput{
        LearnerRef: "STU-300900111",
        FormType:   "transfer_credit",
        Term:       "Fall 2026",
        Fields: map[string]string{
            "prior_institution": "Example University",
            "prior_course_code": "MGMT 101",
            "target_program":    "Business Administration Diploma",
        },
    }

    got, err := svc.Create(context.Background(), input)
    if err != nil {
        t.Fatalf("Create returned error: %v", err)
    }

    if got.CaseID == "" {
        t.Fatal("expected case id")
    }

    if got.Status != cases.StatusSubmitted {
        t.Fatalf("expected submitted status, got %s", got.Status)
    }
}
```

## Rule Engine Test

```go
package rules_test

import (
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/rules"
)

func TestRuleEngine_ApprovalRequest_NeverAutoApproves(t *testing.T) {
    t.Parallel()

    engine := rules.NewEngine()

    result := engine.Evaluate(rules.Input{
        FormType:   "transfer_credit",
        Confidence: 0.99,
        Text:       "Can you approve my transfer credit now?",
        Fields: map[string]string{
            "prior_institution": "Example University",
            "prior_course_code": "MGMT 101",
        },
    })

    if !result.RequiresHumanReview {
        t.Fatal("expected human review for approval request")
    }

    if result.FinalDecisionMade {
        t.Fatal("system must not make final learner-affecting decision")
    }
}
```

## Handler Test

```go
package api_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/platform/httpserver"
)

func TestPOSTCases_WithValidInput_ReturnsCreated(t *testing.T) {
    t.Parallel()

    srv := httpserver.NewTestServer(t)

    body := bytes.NewBufferString(`{
      "form_type": "transfer_credit",
      "learner_ref": "STU-300900111",
      "term": "Fall 2026",
      "fields": {
        "prior_institution": "Example University",
        "prior_course_code": "MGMT 101"
      }
    }`)

    req := httptest.NewRequest(http.MethodPost, "/cases", body)
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    srv.ServeHTTP(rr, req)

    if rr.Code != http.StatusCreated {
        t.Fatalf("expected 201, got %d: %s", rr.Code, rr.Body.String())
    }
}
```

## Adapter Contract Test

```go
package crm_contract_test

import (
    "context"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/adapters/crm"
)

func TestCRMContract_CreateQueueItem_Success(t *testing.T) {
    t.Parallel()

    adapter := crm.NewMockAdapter()

    result, err := adapter.CreateQueueItem(context.Background(), crm.QueueItemInput{
        CaseID:         "case-123",
        IdempotencyKey: "idem-123",
        Route:          "registrar_transfer_credit",
        Summary:        "Synthetic transfer credit request ready for review.",
    })

    if err != nil {
        t.Fatalf("CreateQueueItem returned error: %v", err)
    }

    if result.QueueID == "" {
        t.Fatal("expected queue id")
    }
}
```

## Outbox Worker Test

```go
package outbox_test

import (
    "context"
    "testing"
    "time"

    "github.com/yourname/student-forms-orchestrator/internal/outbox"
)

func TestWorker_WhenAdapterFails_RetriesEvent(t *testing.T) {
    t.Parallel()

    repo := outbox.NewFakeRepository()
    clock := outbox.NewFakeClock(time.Date(2026, 5, 5, 10, 0, 0, 0, time.UTC))
    processor := outbox.NewFailingProcessor()

    worker := outbox.NewWorker(repo, processor, clock)

    eventID := repo.AddPending("case.routed", []byte(`{"case_id":"case-123"}`))

    err := worker.ProcessOne(context.Background())
    if err == nil {
        t.Fatal("expected processing error")
    }

    event := repo.MustGet(eventID)
    if event.AttemptCount != 1 {
        t.Fatalf("expected attempt count 1, got %d", event.AttemptCount)
    }

    if event.Status != outbox.StatusPending {
        t.Fatalf("expected event to be pending for retry, got %s", event.Status)
    }
}
```

## Evaluation Test

```go
package evaluation_test

import (
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/evaluation"
)

func TestEvaluation_FalseAutoRouteRate_BelowThreshold(t *testing.T) {
    t.Parallel()

    fixtures := evaluation.LoadFixtures(t, "testdata/evaluation")
    report := evaluation.Run(fixtures)

    if report.FalseAutoRouteRate > 0.02 {
        t.Fatalf("false auto-route rate %.2f exceeds threshold", report.FalseAutoRouteRate)
    }
}
```

---

# Strict PR Checklist

Every pull request must answer:

- [ ] What failing test was written first?
- [ ] What behaviour does the test prove?
- [ ] Did the test fail before implementation?
- [ ] What is the smallest code written to pass it?
- [ ] What was refactored after green?
- [ ] Are there unit tests?
- [ ] Are there integration tests if persistence/API changed?
- [ ] Are there contract tests if adapter changed?
- [ ] Are there E2E tests if user journey changed?
- [ ] Did privacy/security/accessibility guardrails pass?
- [ ] Did coverage gate pass?
- [ ] Was documentation updated?

## Mutation Testing Note

For important rule and guardrail packages, consider mutation testing later.

Suggested target areas:

- [ ] Confidence threshold comparisons
- [ ] No-auto-approval rule
- [ ] No-auto-denial rule
- [ ] Missing-field detection
- [ ] Duplicate detection
- [ ] Authorization checks
- [ ] Upload allowlist

Mutation testing is optional for the applicant MVP, but mentioning it in the roadmap shows quality maturity.
