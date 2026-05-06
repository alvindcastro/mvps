# 03 - TDD Data Model and API Contract

## Database Rule

Schema changes require failing repository or migration tests first. In Phase 1, only create the tables and fields needed for the transfer-credit vertical slice.

## Phase 1 Schema Boundary

- Keep Phase 1 to `cases` and `status_events`.
- Do not add `documents`, `extractions`, `reviews`, `adapter_calls`, or `outbox_events` until later phases need them.
- Do not add list endpoints, metrics endpoints, or generalized workflow tables during the first slice.

## Phase 1 Minimum Tables

### `cases`

- `id` UUID primary key
- `case_number` text unique not null
- `learner_ref` text not null
- `form_type` text not null
- `status` text not null
- `route` text not null
- `route_reason` text not null
- `request_fields_json` JSONB not null
- `idempotency_key` text unique not null
- `submitted_at` timestamptz not null
- `created_at` timestamptz not null

### `status_events`

- `id` UUID primary key
- `case_id` UUID not null references `cases(id)`
- `event_type` text not null
- `status` text not null
- `message` text not null
- `created_at` timestamptz not null

## Migration TDD

Before creating either table, write a failing repository test that needs the table and its constraints.

## Phase 1 - Repository and Timeline TDD

### Red Tests

- [ ] `TestCaseRepository_CreateTransferCreditCase_PersistsSubmittedCase`
- [ ] `TestCaseRepository_DuplicateIdempotencyKey_ReturnsExistingCase`
- [ ] `TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent`
- [ ] `TestCaseRepository_CreateAndTimeline_AreTransactional`

### Green Tasks

- [ ] Create the `cases` migration with unique constraints on `case_number` and `idempotency_key`.
- [ ] Create the `status_events` migration with a foreign key to `cases`.
- [ ] Implement repository create and get-by-id methods.
- [ ] Implement timeline read by `case_id`.
- [ ] Wrap case creation and submitted-event insertion in one transaction.

### Refactor Tasks

- [ ] Extract row-scan helpers after the first round-trip test passes.
- [ ] Add repository fixture builders for transfer-credit requests.
- [ ] Add a migration reset helper for integration tests.

### Acceptance Criteria

- [ ] Repository tests pass against real PostgreSQL.
- [ ] Duplicate submission protection is enforced by the database.
- [ ] A newly created case always has one submitted timeline event.
- [ ] Transaction rollback prevents orphaned case or timeline records.

---

## API Contract Rule

OpenAPI and handler changes must follow failing API tests. Phase 1 exposes only the transfer-credit endpoints needed by the first slice.

## Phase 1 Endpoints

```text
POST /cases
GET  /cases/{case_id}/timeline
```

## Phase 1 Request Contract

### `POST /cases`

```json
{
  "form_type": "transfer_credit",
  "learner_ref": "STU-300900111",
  "term": "Fall 2026",
  "fields": {
    "prior_institution": "Example University",
    "prior_course_code": "MGMT 101",
    "target_program": "Business Administration Diploma"
  }
}
```

## Phase 1 Success Contract

### `POST /cases` returns `201 Created`

```json
{
  "case_id": "case-123",
  "case_number": "TC-2026-0001",
  "status": "submitted",
  "route": "registrar_transfer_credit",
  "route_reason": "transfer_credit_complete",
  "timeline_url": "/cases/case-123/timeline"
}
```

### `GET /cases/{case_id}/timeline` returns `200 OK`

```json
{
  "case_id": "case-123",
  "events": [
    {
      "event_type": "submitted",
      "status": "submitted",
      "message": "Transfer credit request submitted.",
      "created_at": "2026-05-05T10:00:00Z"
    }
  ]
}
```

## Phase 1 Red API Tests

- [ ] `TestPOSTCases_ValidTransferCredit_Returns201AndSubmittedCase`
- [ ] `TestPOSTCases_MissingPriorInstitution_Returns400WithFieldError`
- [ ] `TestPOSTCases_DuplicateSubmission_ReturnsExistingCase`
- [ ] `TestGETTimeline_ForNewTransferCreditCase_ReturnsSubmittedEvent`

## Phase 1 Green Tasks

- [ ] Implement `POST /cases` with transfer-credit validation only.
- [ ] Persist the case and submitted timeline event before returning `201`.
- [ ] Return the deterministic route and route reason used by the reviewer queue.
- [ ] Implement `GET /cases/{case_id}/timeline`.
- [ ] Add a minimal OpenAPI document for these two endpoints only.

## Phase 1 Refactor Tasks

- [ ] Extract shared request decoding and validation helpers.
- [ ] Keep transport DTOs separate from the domain model.
- [ ] Add golden JSON tests for success and validation failure responses.

## Standard Error Shape

```json
{
  "error": {
    "code": "validation_error",
    "message": "The request contains invalid fields.",
    "fields": [
      {
        "name": "prior_institution",
        "message": "prior_institution is required"
      }
    ],
    "correlation_id": "req-123"
  }
}
```

## Phase 1 Acceptance Criteria

- [ ] The only required write path is `POST /cases`.
- [ ] The only required read path is `GET /cases/{case_id}/timeline`.
- [ ] API behaviour matches the repository transaction and timeline semantics.
- [ ] API tests run in CI before any broader endpoint surface is added.

---

## Deferred Until After Phase 1

### Later Tables

- `documents`
- `extractions`
- `reviews`
- `adapter_calls`
- `outbox_events`

### Later Endpoints

```text
GET  /healthz
GET  /readyz
GET  /cases
GET  /cases/{case_id}
POST /cases/{case_id}/documents
POST /cases/{case_id}/triage
POST /cases/{case_id}/reviews
GET  /metrics/summary
```
