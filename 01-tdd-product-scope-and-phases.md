# 01 - TDD Product Scope and Phases

## MVP Name

**Student Forms Triage and Status Orchestrator**

## Strict TDD Product Rule

Every product behaviour must be described as an executable test before implementation.

No vague task such as "build triage" is acceptable. It must become testable behaviour:

```text
Given a complete transfer-credit request
When the learner submits the form
Then the system creates a case
And the case is routed to registrar_transfer_credit
And the learner sees a status timeline
And no approval decision is made
```

## In Scope

- Transfer credit request triage
- Prerequisite/corequisite waiver triage
- Refund/withdrawal request triage
- Synthetic learner records
- Mock Banner/SIS adapter
- Mock CRM/case queue adapter
- Mock LMS/notification adapter
- Mock knowledge adapter
- Learner-facing status timeline
- Staff reviewer console
- Audit event log
- Confidence-based routing
- Missing-information checklist
- Duplicate detection
- Go API
- PostgreSQL
- Docker Compose
- CI test gates

## Out of Scope

- Real student data
- Real production integrations
- Automatic approvals
- Automatic denials
- Payment processing
- Production identity management
- Care/wellbeing automation
- Final institutional decisions

## TDD Acceptance Model

Each feature must define:

- [ ] User story
- [ ] Behaviour scenarios
- [ ] Failing unit test
- [ ] Failing integration test if persistence/API is involved
- [ ] Failing contract test if adapters are involved
- [ ] Failing E2E test if user journey is involved
- [ ] Coverage expectation
- [ ] Refactor notes
- [ ] Documentation update

---

# Phase 0 - Testable Product Framing

## Goal

Convert the MVP idea into behaviours that can be tested.

## Tickable TDD Tasks

- [ ] Write behaviour scenarios for transfer credit.
- [ ] Write behaviour scenarios for prerequisite waiver.
- [ ] Write behaviour scenarios for refund/withdrawal.
- [ ] Write behaviour scenarios for duplicate submission.
- [ ] Write behaviour scenarios for low-confidence triage.
- [ ] Write behaviour scenarios for missing information.
- [ ] Write behaviour scenarios for human review.
- [ ] Write behaviour scenarios for no-auto-approval.
- [ ] Write behaviour scenarios for no-auto-denial.
- [ ] Write behaviour scenarios for status timeline.
- [ ] Convert behaviours into initial test names.
- [ ] Add test plan to README before coding.
- [ ] Create `testdata/scenarios/` folder.
- [ ] Add synthetic scenario fixtures.
- [ ] Commit product tests and scenarios first.

## Required Test Names

- [ ] `TestCreateTransferCreditCase_WithCompleteInput_CreatesSubmittedCase`
- [ ] `TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError`
- [ ] `TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute`
- [ ] `TestTriageUnknownRequest_RoutesToManualReview`
- [ ] `TestDuplicateSubmission_ReturnsExistingCase`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoApproves`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoDenies`
- [ ] `TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent`

## Definition of Done

- [ ] Product behaviours are written as Given/When/Then.
- [ ] Test names exist before implementation.
- [ ] Scenario fixtures exist before implementation.
- [ ] Initial tests fail because implementation does not exist yet.

---

# Phase 1 - Transfer Credit Vertical Slice

## Goal

Build one end-to-end workflow strictly through failing tests.

## Red Tests First

- [ ] Write failing unit test for transfer-credit validation.
- [ ] Write failing unit test for route selection.
- [ ] Write failing repository test for case creation.
- [ ] Write failing API test for `POST /cases`.
- [ ] Write failing API test for `GET /cases/{id}/timeline`.
- [ ] Write failing E2E test for learner submission to reviewer queue.

## Green Implementation Tasks

- [ ] Implement minimum case model.
- [ ] Implement minimum validation.
- [ ] Implement minimum repository.
- [ ] Implement minimum API handler.
- [ ] Implement minimum timeline event.
- [ ] Implement minimum reviewer queue.

## Refactor Tasks

- [ ] Extract case service.
- [ ] Extract validation rules.
- [ ] Extract route decision object.
- [ ] Remove duplication in test fixtures.
- [ ] Add table-driven tests.

## Phase 1 Acceptance Criteria

- [ ] All transfer-credit tests pass.
- [ ] Coverage for case domain is at least 90%.
- [ ] E2E learner-to-reviewer test passes.
- [ ] No production code was added without a prior failing test.
- [ ] README includes the transfer-credit demo path.

---

# Phase 2 - Multi-Workflow TDD

## Goal

Add prerequisite waiver and refund workflows without breaking transfer credit.

## Red Tests First

- [ ] Write failing test for prerequisite waiver required fields.
- [ ] Write failing test for prerequisite waiver route.
- [ ] Write failing test for refund required fields.
- [ ] Write failing test for refund route.
- [ ] Write failing test for ambiguous refund/withdrawal request.
- [ ] Write failing regression test proving transfer credit still passes.

## Green Implementation Tasks

- [ ] Add form type enum.
- [ ] Add workflow-specific required fields.
- [ ] Add workflow-specific route map.
- [ ] Add missing information messages.
- [ ] Add route reason codes.

## Refactor Tasks

- [ ] Move workflow rules to table/config.
- [ ] Use table-driven tests for all workflow types.
- [ ] Simplify route decision code.
- [ ] Add shared test helpers.

## Acceptance Criteria

- [ ] All three workflows pass validation tests.
- [ ] All three workflows pass routing tests.
- [ ] Unknown workflow routes to manual triage.
- [ ] Transfer-credit regression tests still pass.

---

# Phase 3 - Human Review and Guardrails

## Goal

Guarantee learner-affecting decisions are never automated.

## Red Tests First

- [ ] Write failing test that approval requests require human review.
- [ ] Write failing test that denial requests require human review.
- [ ] Write failing test that refund outcome requests require human review.
- [ ] Write failing test that low confidence requires human review.
- [ ] Write failing test that conflicting extracted fields require human review.
- [ ] Write failing test that unreadable documents require human review.

## Green Implementation Tasks

- [ ] Add `requires_human_review`.
- [ ] Add adverse decision detector.
- [ ] Add confidence thresholds.
- [ ] Add conflict detector.
- [ ] Add document quality trigger.
- [ ] Add reviewer decision endpoint.

## Refactor Tasks

- [ ] Centralize guardrail rules.
- [ ] Add guardrail reason codes.
- [ ] Add test matrix for guardrail combinations.

## Acceptance Criteria

- [ ] No auto-approval test passes.
- [ ] No auto-denial test passes.
- [ ] Learner-impacting decisions require review.
- [ ] Guardrail package coverage is at least 95%.

---

# Phase 4 - Operational TDD

## Goal

Make the project measurable and reliable.

## Red Tests First

- [ ] Write failing test for duplicate detection.
- [ ] Write failing test for idempotency key generation.
- [ ] Write failing test for outbox event creation.
- [ ] Write failing test for outbox retry.
- [ ] Write failing test for dead-letter handling.
- [ ] Write failing test for metrics summary.
- [ ] Write failing test for audit event creation.

## Green Implementation Tasks

- [ ] Add idempotency key.
- [ ] Add outbox table.
- [ ] Add worker.
- [ ] Add retry policy.
- [ ] Add dead-letter state.
- [ ] Add metrics endpoint.
- [ ] Add audit log.

## Refactor Tasks

- [ ] Extract outbox service.
- [ ] Extract metrics service.
- [ ] Add fake clock for deterministic tests.
- [ ] Add fake worker dependencies.

## Acceptance Criteria

- [ ] Duplicate submission does not create duplicate queue item.
- [ ] Outbox retries safely.
- [ ] Metrics are generated from real stored data.
- [ ] Audit events are append-only.
