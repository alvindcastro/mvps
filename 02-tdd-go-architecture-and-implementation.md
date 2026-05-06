# 02 - TDD Go Architecture and Implementation

## Architectural Rule

Architecture must emerge from tests. Phase 1 starts with one transfer-credit vertical slice and only the packages needed to make those failing tests pass.

## Phase 1 Boundary

- Build only the transfer-credit submission path.
- Persist only the case record and its submitted timeline event.
- Expose only the API surface needed by the Phase 1 tests.
- Represent the reviewer queue from case status and route instead of adding a separate queue subsystem.
- Defer duplicate detection, idempotency, and low-confidence routing until later phases.
- Defer health/readiness hardening, adapters, outbox, metrics, and multi-workflow logic until the transfer-credit slice is green.

## Recommended Go Structure for Phase 1

```text
student-forms-orchestrator/
├── cmd/
│   └── api/
├── internal/
│   ├── cases/
│   │   ├── case.go
│   │   ├── service.go
│   │   ├── validation.go
│   │   ├── repository.go
│   │   ├── service_test.go
│   │   └── repository_integration_test.go
│   ├── triage/
│   │   ├── decision.go
│   │   └── decision_test.go
│   ├── reviewerqueue/
│   │   ├── query.go
│   │   └── query_test.go
│   ├── platform/
│   │   ├── database/
│   │   └── httpserver/
│   │       ├── handler.go
│   │       └── handler_integration_test.go
│   └── testsupport/
├── migrations/
├── testdata/
│   └── scenarios/
├── e2e/
│   └── transfer_credit_submission_test.go
└── docs/
```

## Earliest Implementation Sequence

1. Write a failing unit test for transfer-credit validation.
2. Write a failing unit test for transfer-credit route selection.
3. Write failing guardrail tests proving the slice never auto-approves and never auto-denies.
4. Write a failing unit test that a complete transfer-credit input creates a submitted case.
5. Write a failing repository integration test that creates a submitted case and timeline event.
6. Write a failing handler integration test for `POST /cases`.
7. Write a failing handler integration test for `GET /cases/{id}/timeline`.
8. Write a failing reviewer-queue query test.
9. Write a failing E2E test proving a learner submission appears in the reviewer queue.
10. Only after those failures exist, add the smallest production code needed to pass them in order.

## Build Tags

Use build tags to keep slow tests explicit.

```go
//go:build integration
```

```go
//go:build e2e
```

## Phase 1 - Transfer Credit Vertical Slice

### Red Tests

- [ ] `TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError`
- [ ] `TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoApproves`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoDenies`
- [ ] `TestCreateTransferCreditCase_WithCompleteInput_CreatesSubmittedCase`
- [ ] `TestCaseRepository_CreateTransferCreditCase_PersistsSubmittedCase`
- [ ] `TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent`
- [ ] `TestPOSTCases_ValidTransferCredit_Returns201AndSubmittedCase`
- [ ] `TestGETTimeline_ForNewTransferCreditCase_ReturnsSubmittedEvent`
- [ ] `TestReviewerQueue_NewTransferCreditCase_IsVisibleToStaff`
- [ ] `TestE2ETransferCreditSubmission_AppearsInReviewerQueue`

### Green Implementation

- [ ] Create a minimal `Case` aggregate with `submitted` status.
- [ ] Create `CreateInput` for transfer-credit requests only.
- [ ] Implement validation for the required transfer-credit fields.
- [ ] Implement a deterministic route decision that returns `registrar_transfer_credit`.
- [ ] Implement guardrail rules that block approval or denial automation.
- [ ] Implement a repository that writes the case row and the submitted timeline event in one transaction.
- [ ] Implement an HTTP handler for `POST /cases`.
- [ ] Implement an HTTP handler for `GET /cases/{id}/timeline`.
- [ ] Implement the smallest reviewer-queue query/read model needed by the E2E test.

### Refactor Tasks

- [ ] Extract `CaseService` after the first green path passes.
- [ ] Extract validation helpers from the handler into the domain layer.
- [ ] Extract a route decision object instead of scattering string literals.
- [ ] Add `httptest` and database helper fixtures.
- [ ] Convert validation and routing coverage to table-driven tests.

### Acceptance Criteria

- [ ] The transfer-credit slice passes unit, integration, and E2E tests.
- [ ] The first reviewer-visible state is derived from `submitted` status plus `registrar_transfer_credit` route.
- [ ] API handlers remain thin and are testable without starting a real server.
- [ ] The slice proves no approval or denial path is automated.
- [ ] No production code was added before its corresponding test failed.
- [ ] Coverage for `internal/cases` is at least 90%.

---

## Phase 2 - Platform Hardening After the Slice

### Red Tests

- [ ] `TestHealthz_ReturnsOK`
- [ ] `TestReadyz_WhenDatabaseUnavailable_ReturnsServiceUnavailable`
- [ ] `TestConfig_LoadsRequiredEnvironment`
- [ ] `TestLogger_RedactsSensitiveFields`

### Green Implementation

- [ ] Add health and readiness endpoints after the core slice is stable.
- [ ] Add config loading.
- [ ] Add logger redaction.

## Phase 3 - Workflow Expansion

### Red Tests

- [ ] `TestCreatePrerequisiteWaiverCase_WithMissingCourseCode_ReturnsValidationError`
- [ ] `TestCreateRefundWithdrawalCase_WithMissingReason_ReturnsValidationError`
- [ ] `TestTriageUnknownRequest_RoutesToManualReview`

### Green Implementation

- [ ] Add multi-workflow validation and routing.
- [ ] Keep transfer-credit tests passing unchanged.

## Phase 4 - Adapters and Outbox

### Rule

Every adapter or worker component needs a failing contract or integration test before implementation.

### Red Tests

- [ ] `TestCRMContract_CreateQueueItem_Success`
- [ ] `TestOutbox_OnCaseCreated_StoresPendingEvent`
- [ ] `TestWorker_ProcessPendingRouteEvent_CreatesCRMQueueItem`

### Green Implementation

- [ ] Add mock CRM adapter.
- [ ] Add outbox persistence.
- [ ] Add worker processing.
