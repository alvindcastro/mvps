# 10 - TDD Phase 1 Traceability Matrix

## Purpose

This file maps the Phase 1 scope in `01-tdd-product-scope-and-phases.md` to the exact tests, synthetic fixtures, and demo proof required for the transfer-credit vertical slice.

Use it to prevent Phase 1 drift into later work such as duplicate handling, low-confidence fallback logic, or broader workflow classification.

## Phase 1 Rule

Only the canonical transfer-credit slice is in Phase 1:

- Synthetic learner submits a complete transfer-credit request.
- Required fields are validated before case creation.
- The system creates a case and suggests `registrar_transfer_credit`.
- The case is visible in the reviewer queue query/read model.
- The learner sees a submitted timeline.
- No approval or denial decision is automated.

## Fixture Inventory

Use the fixture contract in `testdata/scenarios/README.md`.

| Fixture ID | Planned Path | Purpose |
|---|---|---|
| `phase1-transfer-credit-complete` | `testdata/scenarios/transfer-credit-complete.json` | Happy-path transfer-credit submission |
| `phase1-transfer-credit-missing-prior-institution` | `testdata/scenarios/transfer-credit-missing-prior-institution.json` | Required-field validation failure |
| `phase1-transfer-credit-adverse-decision-request` | `testdata/scenarios/transfer-credit-adverse-decision-request.json` | Guardrail proof for no auto-approval or denial |

## Outcome Traceability

| Phase 1 outcome | Fixture | Required tests | Demo proof |
|---|---|---|---|
| Required transfer-credit fields are validated | `phase1-transfer-credit-missing-prior-institution` | `TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError`, `TestPOSTCases_MissingPriorInstitution_Returns400WithFieldError` | Invalid submission fails before case creation |
| Complete input routes to `registrar_transfer_credit` | `phase1-transfer-credit-complete` | `TestTriageTransferCredit_WithCompleteInput_SuggestsRegistrarRoute` | Created case response shows `registrar_transfer_credit` |
| Complete input creates a submitted case | `phase1-transfer-credit-complete` | `TestCreateTransferCreditCase_WithCompleteInput_CreatesSubmittedCase`, `TestCaseRepository_CreateTransferCreditCase_PersistsSubmittedCase`, `TestPOSTCases_ValidTransferCredit_Returns201AndSubmittedCase` | Created case has stable ID, case number, and `submitted` status |
| Case creation appends the submitted timeline event | `phase1-transfer-credit-complete` | `TestStatusTimeline_OnCaseCreation_ContainsSubmittedEvent`, `TestGETTimeline_ForNewTransferCreditCase_ReturnsSubmittedEvent` | Timeline response includes one initial `submitted` event |
| Reviewer visibility exists for the submitted case | `phase1-transfer-credit-complete` | `TestReviewerQueue_NewTransferCreditCase_IsVisibleToStaff`, `TestE2ETransferCreditSubmission_AppearsInReviewerQueue` | Reviewer queue query/read model shows the case after learner submission |
| No approval or denial outcome is automated | `phase1-transfer-credit-adverse-decision-request` | `TestAdverseDecisionGuardrail_NeverAutoApproves`, `TestAdverseDecisionGuardrail_NeverAutoDenies`, `TestTransferCreditTriage_AdverseDecisionRequest_RequiresHumanReview` | Demo callout and test output show human review is required |

## Explicitly Deferred

These are not Phase 1 obligations unless `01-tdd-product-scope-and-phases.md` changes:

- Duplicate detection
- Idempotency keys
- Low-confidence fallback logic
- Unsupported institution mapping
- Multi-workflow classification
- Real adapter contracts
- Reviewer decision endpoints

## Documentation Sync

When Phase 1 changes, update:

- `README.md`
- `01-tdd-product-scope-and-phases.md`
- `02-tdd-go-architecture-and-implementation.md`
- `03-tdd-data-model-and-api-contract.md`
- `04-tdd-ai-triage-workflow-and-rules.md`
- `07-tdd-30-60-90-day-build-plan.md`
- `08-tdd-ci-makefile-and-test-skeletons.md`
- `09-tdd-phase-1-transfer-credit-vertical-slice.md`
- `testdata/scenarios/README.md`
