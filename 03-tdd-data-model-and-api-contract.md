# 03 - TDD Data Model and API Contract

## Database Rule

Schema changes require failing repository or migration tests first.

## Core Tables

- `cases`
- `case_inputs`
- `documents`
- `extractions`
- `status_events`
- `reviews`
- `adapter_calls`
- `outbox_events`

## Migration TDD

Before creating a table, write a failing test that needs it.

## Phase 1 - Cases Schema

### Red Tests

- [ ] `TestCaseRepository_CreateAndGet_RoundTrip`
- [ ] `TestCaseRepository_ListByStatus_ReturnsMatchingCases`
- [ ] `TestCaseRepository_DuplicateIdempotencyKey_ReturnsExistingCase`
- [ ] `TestCaseRepository_UpdateStatus_AppendsTimelineEvent`

### Green Tasks

- [ ] Create `cases` migration.
- [ ] Create `status_events` migration.
- [ ] Implement repository methods.
- [ ] Add transaction boundary.

### Refactor Tasks

- [ ] Extract scan helpers.
- [ ] Add repository test fixtures.
- [ ] Add migration reset helper.

### Acceptance Criteria

- [ ] Repository tests pass against real PostgreSQL.
- [ ] Duplicate idempotency key is enforced by database constraint.
- [ ] Status updates are transactional.

---

# Phase 2 - Documents and Extractions Schema

## Red Tests

- [ ] `TestDocumentRepository_SaveDocumentMetadata_PersistsHash`
- [ ] `TestDocumentRepository_DuplicateHashForCase_FlagsDuplicate`
- [ ] `TestExtractionRepository_SaveExtraction_PersistsFieldConfidence`
- [ ] `TestExtractionRepository_ListByCase_ReturnsAllExtractions`

## Green Tasks

- [ ] Create `documents` table.
- [ ] Create `extractions` table.
- [ ] Add file hash field.
- [ ] Add document quality field.
- [ ] Add extraction confidence field.
- [ ] Implement repositories.

## Refactor Tasks

- [ ] Add fixture builders.
- [ ] Add typed confidence value.
- [ ] Add document validation helpers.

## Acceptance Criteria

- [ ] Duplicate documents are testable.
- [ ] Extraction results round-trip.
- [ ] Field confidence is preserved.

---

# Phase 3 - API Contract TDD

## API Contract Rule

OpenAPI changes must be accompanied by failing handler tests.

## Endpoints

```text
GET    /healthz
GET    /readyz
POST   /cases
GET    /cases
GET    /cases/{case_id}
GET    /cases/{case_id}/timeline
POST   /cases/{case_id}/documents
POST   /cases/{case_id}/triage
POST   /cases/{case_id}/reviews
GET    /metrics/summary
```

## Red API Tests

- [ ] `TestPOSTCases_ValidTransferCredit_Returns201AndCaseNumber`
- [ ] `TestPOSTCases_MissingLearnerRef_Returns400WithFieldError`
- [ ] `TestGETCases_WithStatusFilter_ReturnsMatchingCases`
- [ ] `TestGETTimeline_ForNewCase_ReturnsSubmittedEvent`
- [ ] `TestPOSTDocuments_WithUnsupportedType_Returns415`
- [ ] `TestPOSTTriage_ForUnknownCase_Returns404`
- [ ] `TestPOSTReviews_WithInvalidDecision_Returns400`
- [ ] `TestGETMetricsSummary_ReturnsCounts`

## Green Tasks

- [ ] Implement handlers.
- [ ] Implement request validation.
- [ ] Implement error response format.
- [ ] Implement response DTOs.
- [ ] Implement pagination.
- [ ] Implement metrics query.
- [ ] Add OpenAPI file.

## Refactor Tasks

- [ ] Generate types from OpenAPI or validate DTOs against spec.
- [ ] Add shared API test helper.
- [ ] Add golden response tests.
- [ ] Keep domain separate from transport DTOs.

## Standard Error Shape

```json
{
  "error": {
    "code": "validation_error",
    "message": "The request contains invalid fields.",
    "fields": [
      {
        "name": "learner_ref",
        "message": "learner_ref is required"
      }
    ],
    "correlation_id": "req-123"
  }
}
```

## Acceptance Criteria

- [ ] Every endpoint has success and failure tests.
- [ ] Every write endpoint creates audit or status event where appropriate.
- [ ] API spec matches test behaviour.
- [ ] API tests run in CI.

---

# Phase 4 - Adapter Contract TDD

## Contract Test Matrix

| Adapter | Contract Scenarios |
|---|---|
| SIS | found, not found, unavailable, invalid learner reference |
| CRM | create queue item, duplicate idempotency key, unavailable |
| LMS/Notification | send success, invalid target, unavailable |
| Knowledge | search found, no result, stale result, cited result |

## Red Tests

- [ ] Write adapter contract tests before mock adapters.
- [ ] Write adapter failure tests before retry code.
- [ ] Write redaction tests before adapter logging.
- [ ] Write latency metric tests before instrumentation.

## Green Tasks

- [ ] Implement mock adapters.
- [ ] Implement adapter call logging.
- [ ] Implement retryable error types.
- [ ] Implement non-retryable error types.

## Refactor Tasks

- [ ] Share adapter contract suite.
- [ ] Add mock fixtures.
- [ ] Add table-driven adapter tests.

## Acceptance Criteria

- [ ] Contract tests prove replaceability.
- [ ] Mock adapters can simulate failure.
- [ ] Domain logic depends only on interfaces.
- [ ] Adapter calls are redacted in logs.

---

# Phase 5 - Data Integrity TDD

## Red Tests

- [ ] `TestIdempotencyKey_SameNormalizedInput_SameKey`
- [ ] `TestIdempotencyKey_DifferentTerm_DifferentKey`
- [ ] `TestStatusEvents_AreAppendOnly`
- [ ] `TestAuditEvents_CannotBeUpdatedByRepository`
- [ ] `TestCaseStatus_InvalidTransition_ReturnsError`

## Green Tasks

- [ ] Implement normalized idempotency key.
- [ ] Add status transition validator.
- [ ] Add append-only repository methods.
- [ ] Add database constraints.

## Refactor Tasks

- [ ] Extract state machine.
- [ ] Add transition table.
- [ ] Add state machine tests.

## Acceptance Criteria

- [ ] Invalid transitions fail.
- [ ] Audit/status events are append-only.
- [ ] Idempotency is deterministic.
