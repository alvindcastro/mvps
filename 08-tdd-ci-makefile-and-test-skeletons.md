# 08 - TDD CI, Makefile, and Go Test Skeletons

## Phase 1 CI Rule

CI for the first implementation slice must prove the transfer-credit workflow from red tests to green tests before broader suites are added. In Phase 1, prioritize unit, integration, and E2E support for the transfer-credit path only.

## Phase 1 Makefile Targets

```makefile
.PHONY: test-unit test-integration test-e2e test-phase1 coverage-case-domain coverage-phase1 ci-phase1 ci-full lint

test-unit:
	go test ./internal/cases ./internal/triage ./internal/reviewerqueue ./internal/platform/httpserver

test-integration:
	go test -tags=integration ./internal/cases ./internal/platform/httpserver

test-e2e:
	go test -tags=e2e ./e2e/...

test-phase1: test-unit test-integration test-e2e

coverage-case-domain:
	go test ./internal/cases -coverprofile=coverage-cases.out
	go tool cover -func=coverage-cases.out

coverage-phase1:
	go test ./internal/cases ./internal/triage ./internal/platform/httpserver -coverprofile=coverage.out
	go tool cover -func=coverage.out

lint:
	golangci-lint run

ci-phase1: test-phase1 coverage-case-domain

ci-full: lint ci-phase1 coverage-phase1
```

## Phase 1 Pipeline Order

1. Run unit tests for transfer-credit validation and route selection.
2. Run unit tests for the no-auto-approval and no-auto-denial guardrails.
3. Run integration tests for repository persistence and API handlers.
4. Run the E2E learner-submission-to-reviewer-queue test.
5. Run the `internal/cases` coverage gate at 90% or better.
6. Add lint and broader slice-coverage gates only after the canonical Phase 1 workflow is stable.

## Post-Phase-1 Slice Coverage Script

Use this broader slice gate only after the canonical Phase 1 workflow is stable and the case-domain 90% gate is already green.

```bash
#!/usr/bin/env bash
set -euo pipefail

go test ./internal/cases ./internal/triage ./internal/platform/httpserver -coverprofile=coverage.out
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

## Case Domain Coverage Gate

Use a separate package gate for the Phase 1 case domain.

```bash
#!/usr/bin/env bash
set -euo pipefail

go test ./internal/cases -coverprofile=coverage-cases.out
coverage=$(go tool cover -func=coverage-cases.out | awk '/total:/ {print substr($3, 1, length($3)-1)}')

required=90

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
  phase1-transfer-credit:
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

      - name: Run transfer-credit unit tests
        run: make test-unit

      - name: Run transfer-credit integration tests
        run: make test-integration

      - name: Run transfer-credit E2E tests
        run: make test-e2e

      - name: Run case domain coverage gate
        run: ./scripts/check-case-domain-coverage.sh
```

---

# Go Test Skeletons

## Validation Test

```go
package cases_test

import (
    "context"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/cases"
)

func TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError(t *testing.T) {
    t.Parallel()

    repo := cases.NewFakeRepository()
    svc := cases.NewService(repo)

    _, err := svc.Create(context.Background(), cases.CreateInput{
        LearnerRef: "STU-300900111",
        FormType:   "transfer_credit",
        Term:       "Fall 2026",
        Fields: map[string]string{
            "prior_course_code": "MGMT 101",
            "target_program":    "Business Administration Diploma",
        },
    })

    if err == nil {
        t.Fatal("expected validation error")
    }
}
```

## Route Decision Test

```go
package triage_test

import (
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/triage"
)

func TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute(t *testing.T) {
    t.Parallel()

    result := triage.DecideTransferCreditRoute(triage.Input{
        LearnerRef: "STU-300900111",
        Term:       "Fall 2026",
        Fields: map[string]string{
            "prior_institution": "Example University",
            "prior_course_code": "MGMT 101",
            "target_program":    "Business Administration Diploma",
        },
    })

    if result.Route != "registrar_transfer_credit" {
        t.Fatalf("expected registrar_transfer_credit, got %s", result.Route)
    }
}
```

## Guardrail Test

```go
package triage_test

import (
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/triage"
)

func TestAdverseDecisionGuardrail_NeverAutoApproves(t *testing.T) {
    t.Parallel()

    result := triage.DecideTransferCreditRoute(triage.Input{
        LearnerRef: "STU-300900111",
        Term:       "Fall 2026",
        Fields: map[string]string{
            "prior_institution": "Example University",
            "prior_course_code": "MGMT 101",
            "target_program":    "Business Administration Diploma",
            "requested_outcome": "approve this now",
        },
    })

    if result.Route == "approved" {
        t.Fatal("guardrail failure: route must not auto-approve")
    }
}
```

## Repository Integration Test

```go
//go:build integration

package cases_test

import (
    "context"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/cases"
)

func TestCaseRepository_CreateTransferCreditCase_PersistsSubmittedCase(t *testing.T) {
    t.Parallel()

    repo := cases.NewTestRepository(t)

    got, err := repo.Create(context.Background(), cases.CreateParams{
        LearnerRef: "STU-300900111",
        FormType:   "transfer_credit",
        Status:     "submitted",
        Route:      "registrar_transfer_credit",
    })

    if err != nil {
        t.Fatalf("Create returned error: %v", err)
    }

    if got.Status != "submitted" {
        t.Fatalf("expected submitted status, got %s", got.Status)
    }
}
```

## Handler Integration Test

```go
package api_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/platform/httpserver"
)

func TestPOSTCases_ValidTransferCredit_Returns201AndSubmittedCase(t *testing.T) {
    t.Parallel()

    srv := httpserver.NewTestServer(t)

    body := bytes.NewBufferString(`{
      "form_type": "transfer_credit",
      "learner_ref": "STU-300900111",
      "term": "Fall 2026",
      "fields": {
        "prior_institution": "Example University",
        "prior_course_code": "MGMT 101",
        "target_program": "Business Administration Diploma"
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

## Timeline Handler Test

```go
package api_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/yourname/student-forms-orchestrator/internal/platform/httpserver"
)

func TestGETTimeline_ForNewTransferCreditCase_ReturnsSubmittedEvent(t *testing.T) {
    t.Parallel()

    srv := httpserver.NewTestServerWithSubmittedCase(t)

    req := httptest.NewRequest(http.MethodGet, "/cases/case-123/timeline", nil)
    rr := httptest.NewRecorder()

    srv.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d: %s", rr.Code, rr.Body.String())
    }
}
```

## E2E Test

```go
//go:build e2e

package e2e_test

import "testing"

func TestE2ETransferCreditSubmission_AppearsInReviewerQueue(t *testing.T) {
    t.Parallel()

    t.Fatal("write the red E2E test before implementing the reviewer queue path")
}
```

---

## Strict PR Checklist for Phase 1

Every pull request for the first slice must answer:

- [ ] What failing transfer-credit test was written first?
- [ ] Did the failure happen before implementation?
- [ ] Which layer changed: validation, routing, repository, API, or E2E path?
- [ ] Was schema scope limited to `cases` and `status_events`?
- [ ] Was API scope limited to `POST /cases` and `GET /cases/{id}/timeline`, with reviewer visibility proven through the queue query/read model?
- [ ] Did the guardrail tests prove no auto-approval and no auto-denial?
- [ ] Did unit, integration, and E2E tests pass?
- [ ] Did the `internal/cases` 90% coverage gate pass?
- [ ] Were the Phase 1 docs updated?

## Deferred CI Expansion

Add these only after the transfer-credit slice is stable:

- `test-contract`
- adapter contract jobs
- evaluation jobs
- mutation testing
- broader multi-workflow smoke suites
