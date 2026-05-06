# 09 - TDD Phase 1 Transfer Credit Vertical Slice

## Purpose

This document is the execution map for Phase 1 in `01-tdd-product-scope-and-phases.md`.

Use it to keep the first implementation slice narrow:

- One workflow only: transfer credit
- One learner journey: submit form, create case, view timeline
- One staff journey: see the case in the reviewer queue
- No production code before the related failing test exists

## Phase 1 Boundaries

### Included

- Transfer-credit intake validation
- Transfer-credit route suggestion
- Case creation in PostgreSQL
- Submitted timeline event
- `POST /cases`
- `GET /cases/{id}/timeline`
- Reviewer queue visibility for newly created transfer-credit cases
- Synthetic fixtures and demo walkthrough

### Excluded Until Phase 2+

- Prerequisite waiver
- Refund or withdrawal
- Real document ingestion
- Live AI provider calls
- Real SIS, CRM, LMS, or identity integrations
- Reviewer decisioning beyond queue visibility

## Canonical Test Order

Write and observe failures in this order:

1. `TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError`
2. `TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute`
3. `TestCaseRepository_CreateAndGet_RoundTrip`
4. `TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent`
5. `TestPOSTCases_ValidTransferCredit_Returns201AndCaseNumber`
6. `TestGETTimeline_ForNewCase_ReturnsSubmittedEvent`
7. `TestReviewerQueue_NewTransferCreditCase_IsVisibleToStaff`
8. `TestE2ETransferCreditSubmission_AppearsInReviewerQueue`

Do not pull later tests forward unless an earlier red test proves the next production code cannot be reached.

## Red-Green-Refactor Sequence

### Step 1 - Validation Red

- Create the smallest failing domain test for required transfer-credit fields.
- Add only the minimum validation needed to reject incomplete input.
- Refactor into table-driven validation once the first pass is green.

### Step 2 - Routing Red

- Create a failing rule test for complete transfer-credit routing.
- Return a suggested route and reason code only.
- Do not allow approval, denial, or eligibility outcomes.

### Step 3 - Persistence Red

- Create the failing repository round-trip test against PostgreSQL.
- Add only `cases` and `status_events` support.
- Enforce duplicate protection with an idempotency key before expanding schema.

### Step 4 - API Red

- Create the failing `POST /cases` test.
- Map domain validation errors to the standard API error shape.
- Return the created case identifier and initial submitted status.

### Step 5 - Timeline Red

- Create the failing `GET /cases/{id}/timeline` test.
- Ensure case creation appends the submitted event transactionally.
- Keep the timeline append-only.

### Step 6 - Reviewer Queue Red

- Create the failing reviewer queue test before adding queue listing behaviour.
- Expose only what proves the handoff from learner submission to staff review.
- Keep reviewer actions out of scope unless required for the E2E test harness.

### Step 7 - E2E Red

- Create the learner-to-reviewer failing journey test.
- Prove the synthetic learner submission becomes a visible reviewer work item.
- Use mocks and fixtures only.

## Minimum Fixture Set

Create these fixtures before implementation expands:

- `transfer-credit-complete.json`
- `transfer-credit-missing-prior-institution.json`
- `transfer-credit-duplicate.json`
- `transfer-credit-low-confidence.json`

Each fixture should carry:

- learner reference
- academic term
- transfer-credit fields
- expected route
- expected missing fields
- expected timeline events

## Demo Path

The README demo path for Phase 1 should stay limited to:

1. Submit a complete synthetic transfer-credit form.
2. Show the created case response.
3. Fetch the case timeline and show the submitted event.
4. Open the reviewer queue and show the new case.
5. Call out that no automated approval or denial occurred.

## Evidence Required Before Phase 1 Is Done

- All listed Phase 1 unit, integration, API, and E2E tests are green.
- `internal/cases` domain coverage is at least 90%.
- The route suggestion is deterministic and explainable.
- Timeline writes are append-only and transactional.
- The reviewer queue shows the case without making a learner-affecting decision.
- README demo instructions match the implemented flow.

## Documentation Sync Checklist

Update these docs whenever the slice changes:

- `README.md`
- `01-tdd-product-scope-and-phases.md`
- `02-tdd-go-architecture-and-implementation.md`
- `03-tdd-data-model-and-api-contract.md`
- `04-tdd-ai-triage-workflow-and-rules.md`
- `07-tdd-30-60-90-day-build-plan.md`
- `08-tdd-ci-makefile-and-test-skeletons.md`
