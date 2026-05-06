# Phase 1 Scenario Fixtures

## Purpose

This folder holds synthetic scenario fixtures for the Phase 1 transfer-credit vertical slice.

Do not place real student data here.

## Phase 1 Rule

Only transfer-credit fixtures belong in this folder for the first slice.

Defer duplicate, low-confidence, and non-transfer-credit scenarios until later phases unless `01-tdd-product-scope-and-phases.md` expands Phase 1.

## Planned Phase 1 Fixtures

| Fixture File | Fixture ID | Purpose |
|---|---|---|
| `transfer-credit-complete.json` | `phase1-transfer-credit-complete` | Canonical happy-path submission |
| `transfer-credit-missing-prior-institution.json` | `phase1-transfer-credit-missing-prior-institution` | Required-field validation failure |
| `transfer-credit-adverse-decision-request.json` | `phase1-transfer-credit-adverse-decision-request` | Guardrail proof for no auto-approval or denial |

## Required Fixture Shape

Each fixture should provide:

- `learner_ref`
- `term`
- `form_type`
- `fields`
- `expected_route`
- `expected_route_reason`
- `expected_missing_fields`
- `expected_timeline_events`
- `expected_guardrail_outcome`

Use `null` for `expected_route` and `expected_route_reason` when validation or a guardrail stops routing before a suggested route exists.

Use `expected_guardrail_outcome: "human_review_required"` for the adverse-decision fixture and `null` when no guardrail applies.

## Example Shape

```json
{
  "learner_ref": "STU-300900111",
  "term": "Fall 2026",
  "form_type": "transfer_credit",
  "fields": {
    "prior_institution": "Example University",
    "prior_course_code": "MGMT 101",
    "target_program": "Business Administration Diploma"
  },
  "expected_route": "registrar_transfer_credit",
  "expected_route_reason": "transfer_credit_complete",
  "expected_missing_fields": [],
  "expected_timeline_events": [
    "submitted"
  ],
  "expected_guardrail_outcome": null
}
```

## Linked Docs

- `09-tdd-phase-1-transfer-credit-vertical-slice.md`
- `10-tdd-phase-1-traceability-matrix.md`
