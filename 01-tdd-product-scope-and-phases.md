# 01 - TDD Product Scope and Phases

## MVP Name

**Student Forms Triage and Status Orchestrator**

## Current Repo Focus

The repo is currently aligned around a single executable slice:

- Phase 1 builds transfer-credit only.
- Phase 1 is the default demo path for the repo.
- Prerequisite waiver and refund/withdrawal remain Phase 2 work.

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

## Phase 1 Scope Boundary

For Phase 1, only these outcomes must work end to end:

- Create a transfer-credit case from a synthetic learner submission
- Validate required transfer-credit fields
- Suggest `registrar_transfer_credit`
- Persist the case and initial timeline event
- Show the case in a reviewer queue
- Show the learner-facing submitted status timeline

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

## Canonical Phase 1 Scenario

```text
Given a complete transfer-credit request
When the learner submits the form
Then the system creates a case
And the case is routed to registrar_transfer_credit
And the reviewer queue shows the case
And the learner sees a submitted status timeline
And no approval or denial decision is made
```

## Explicit Deliverables

### Product Deliverables

- [x] One canonical transfer-credit scenario is documented and treated as the Phase 1 demo path.
- [x] Phase 1 scope boundary is documented and excludes other workflow types.
- [x] Acceptance criteria name the exact behaviours required to demo the slice.
- [x] README points to the same Phase 1 demo path.

### Test Deliverables

- [ ] Failing unit tests exist for validation and routing before implementation.
- [ ] Failing persistence and API tests exist before implementation.
- [ ] Failing E2E test exists for learner submission through reviewer visibility.
- [ ] Test fixtures use synthetic transfer-credit data only.
- [ ] Guardrail tests prove no auto-approval and no auto-denial on the slice.

### Slice Deliverables

- [ ] Minimal transfer-credit case model exists.
- [ ] Minimal validation for required fields exists.
- [ ] Minimal route suggestion exists for `registrar_transfer_credit`.
- [ ] Minimal case creation endpoint exists.
- [ ] Minimal learner timeline endpoint exists.
- [ ] Minimal reviewer queue view/query exists.

## Red Tests First

- [ ] Write failing unit test for transfer-credit validation.
- [ ] Write failing unit test for route selection.
- [ ] Write failing unit test for no auto-approval guardrail.
- [ ] Write failing unit test for no auto-denial guardrail.
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

## Phase 1 Exit Checklist

- [ ] A synthetic learner can submit a complete transfer-credit request.
- [ ] The case is stored once and has a stable identifier.
- [ ] The route suggestion is `registrar_transfer_credit`.
- [ ] The initial timeline contains a submitted event.
- [ ] The reviewer queue can surface the submitted case.
- [ ] Missing required fields produce a test-covered validation failure.
- [ ] The slice has no approval or denial automation.
- [ ] README and this file describe the same Phase 1 demo path.

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
