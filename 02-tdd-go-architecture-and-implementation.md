# 02 - TDD Go Architecture and Implementation

## Architectural Rule

Architecture must emerge from tests. Start with a modular monolith and extract boundaries only when tests show a need.

## Recommended Go Structure

```text
student-forms-orchestrator/
├── cmd/
│   ├── api/
│   └── worker/
├── internal/
│   ├── cases/
│   ├── rules/
│   ├── triage/
│   ├── review/
│   ├── audit/
│   ├── outbox/
│   ├── metrics/
│   ├── adapters/
│   │   ├── sis/
│   │   ├── crm/
│   │   ├── lms/
│   │   └── knowledge/
│   └── platform/
│       ├── database/
│       ├── httpserver/
│       └── config/
├── migrations/
├── testdata/
│   ├── scenarios/
│   ├── adapters/
│   └── evaluation/
├── openapi/
├── web/
└── docs/
```

## Test Structure

```text
internal/cases/
├── case.go
├── service.go
├── repository.go
├── service_test.go
├── repository_integration_test.go
└── handler_integration_test.go

internal/rules/
├── engine.go
├── engine_test.go
└── guardrails_test.go

internal/adapters/crm/
├── adapter.go
├── mock.go
└── contract_test.go
```

## Build Tags

Use build tags to separate slow tests.

```go
//go:build integration
```

```go
//go:build e2e
```

## Phase 1 - Go Test Harness First

### Red Tasks

- [ ] Create failing `TestHealthz_ReturnsOK`.
- [ ] Create failing `TestReadyz_WhenDatabaseUnavailable_ReturnsServiceUnavailable`.
- [ ] Create failing `TestConfig_LoadsRequiredEnvironment`.
- [ ] Create failing `TestLogger_RedactsSensitiveFields`.

### Green Tasks

- [ ] Implement HTTP server.
- [ ] Implement health endpoint.
- [ ] Implement readiness endpoint.
- [ ] Implement config loader.
- [ ] Implement logger redaction.

### Refactor Tasks

- [ ] Extract server setup into testable function.
- [ ] Add `httptest` helpers.
- [ ] Add fake config source.
- [ ] Add test cleanup helpers.

### Acceptance Criteria

- [ ] API can be tested without starting external server.
- [ ] Database readiness is testable.
- [ ] Logs redact learner references.
- [ ] Health tests pass before any feature code.

---

# Phase 2 - Case Service TDD

## Red Tests

- [ ] `TestCaseService_Create_WithValidTransferCreditInput_PersistsCase`
- [ ] `TestCaseService_Create_WithMissingLearnerRef_ReturnsValidationError`
- [ ] `TestCaseService_Create_CreatesSubmittedTimelineEvent`
- [ ] `TestCaseService_Create_WithDuplicateIdempotencyKey_ReturnsExistingCase`
- [ ] `TestCaseRepository_CreateAndGet_RoundTripsCase`

## Green Implementation

- [ ] Create `Case` struct.
- [ ] Create `CaseService`.
- [ ] Create `CaseRepository` interface.
- [ ] Create PostgreSQL repository.
- [ ] Create validation errors.
- [ ] Create timeline event writer.
- [ ] Create idempotency key generator.

## Refactor

- [ ] Move validation to domain package.
- [ ] Add table-driven validation tests.
- [ ] Add explicit error types.
- [ ] Add fake repository for service tests.

## Coverage Gate

- [ ] `internal/cases` coverage >= 90%.

---

# Phase 3 - Rule Engine TDD

## Red Tests

- [ ] `TestRuleEngine_TransferCreditComplete_ReturnsRouteReady`
- [ ] `TestRuleEngine_TransferCreditMissingTranscript_ReturnsNeedsMoreInfo`
- [ ] `TestRuleEngine_PrereqWaiverMissingCourseCode_ReturnsNeedsMoreInfo`
- [ ] `TestRuleEngine_RefundOutcomeQuestion_RequiresHumanReview`
- [ ] `TestRuleEngine_LowConfidence_RequiresHumanReview`
- [ ] `TestRuleEngine_ApprovalRequest_NeverAutoApproves`
- [ ] `TestRuleEngine_DenialRequest_NeverAutoDenies`

## Green Implementation

- [ ] Create rule input type.
- [ ] Create rule result type.
- [ ] Add required field rules.
- [ ] Add confidence threshold rule.
- [ ] Add adverse decision guardrail.
- [ ] Add route map.
- [ ] Add missing information result.

## Refactor

- [ ] Convert rules to composable functions.
- [ ] Add rule names and reason codes.
- [ ] Add shared test fixture builder.
- [ ] Keep rule engine deterministic.

## Coverage Gate

- [ ] `internal/rules` coverage >= 95%.

---

# Phase 4 - Adapter TDD

## Rule

Every adapter must have a contract test before implementation.

## Red Tests

- [ ] `TestSISContract_GetStudent_Found`
- [ ] `TestSISContract_GetStudent_NotFound`
- [ ] `TestSISContract_GetStudent_Unavailable`
- [ ] `TestCRMContract_CreateQueueItem_Success`
- [ ] `TestCRMContract_CreateQueueItem_DuplicateIdempotencyKey`
- [ ] `TestLMSContract_SendStatusUpdate_Success`
- [ ] `TestKnowledgeContract_Search_ReturnsCitedSnippets`

## Green Implementation

- [ ] Define interfaces.
- [ ] Implement mock SIS.
- [ ] Implement mock CRM.
- [ ] Implement mock LMS.
- [ ] Implement mock knowledge adapter.
- [ ] Implement adapter call logging.
- [ ] Implement error types.

## Refactor

- [ ] Extract shared adapter test suite.
- [ ] Add fake latency support.
- [ ] Add fake failure support.
- [ ] Keep domain independent of adapter implementations.

## Coverage Gate

- [ ] Adapter contract tests pass.
- [ ] Adapter package coverage >= 90%.

---

# Phase 5 - API Handler TDD

## Red Tests

- [ ] `TestPOSTCases_WithValidInput_ReturnsCreated`
- [ ] `TestPOSTCases_WithInvalidInput_ReturnsBadRequest`
- [ ] `TestGETCase_WithExistingCase_ReturnsCase`
- [ ] `TestGETTimeline_WithExistingCase_ReturnsEvents`
- [ ] `TestPOSTReview_WithReviewerDecision_UpdatesCaseStatus`
- [ ] `TestPOSTTriage_WithExistingCase_ReturnsTriageResult`

## Green Implementation

- [ ] Add routes.
- [ ] Add request DTOs.
- [ ] Add response DTOs.
- [ ] Add validation.
- [ ] Add error mapping.
- [ ] Add correlation ID.
- [ ] Add JSON response helpers.

## Refactor

- [ ] Extract handler dependencies.
- [ ] Add handler test helpers.
- [ ] Add OpenAPI consistency checks.
- [ ] Keep handlers thin.

## Coverage Gate

- [ ] Handler coverage >= 80%.
- [ ] All API tests pass with race detector where possible.

---

# Phase 6 - Worker and Outbox TDD

## Red Tests

- [ ] `TestOutbox_OnCaseCreated_StoresPendingEvent`
- [ ] `TestWorker_ProcessPendingRouteEvent_CreatesCRMQueueItem`
- [ ] `TestWorker_WhenAdapterFails_RetriesEvent`
- [ ] `TestWorker_AfterMaxRetries_MarksDeadLetter`
- [ ] `TestWorker_ReprocessingSameEvent_IsIdempotent`

## Green Implementation

- [ ] Add outbox repository.
- [ ] Add event publisher.
- [ ] Add worker loop.
- [ ] Add retry policy.
- [ ] Add dead-letter state.
- [ ] Add idempotency check.

## Refactor

- [ ] Add fake clock.
- [ ] Add deterministic worker tests.
- [ ] Separate polling from processing.
- [ ] Extract retry calculator.

## Coverage Gate

- [ ] Outbox package coverage >= 90%.
- [ ] Integration tests pass.
