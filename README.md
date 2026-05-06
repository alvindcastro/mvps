# Student Forms Triage and Status Orchestrator MVP

## Strict TDD Applicant Build Package

This package defines a Go-first, strict test-driven development plan for a Student Forms Triage and Status Orchestrator MVP.

The project is designed as an applicant portfolio build for an AI/Automation Solutions Developer role focused on learner digital experience, AI-assisted triage, workflow automation, and enterprise integration readiness.

## Current Build Focus

Phase 1 is the transfer-credit vertical slice.

Until Phase 1 is complete, the repo should optimize for one demo path:

```text
Learner submits transfer-credit request
-> system validates required fields
-> system creates a case
-> system suggests registrar_transfer_credit
-> reviewer queue query shows the case
-> learner sees a submitted status timeline
-> no approval or denial decision is made
```

## Non-Negotiable TDD Rule

No production code is written unless a failing test exists first.

Every feature must follow:

```text
RED   - Write the smallest meaningful failing test.
GREEN - Write the simplest code to pass the test.
REFACTOR - Improve design without changing behaviour.
VERIFY - Run full relevant test suite and coverage gate.
DOCUMENT - Update docs, examples, or decision records.
```

## MVP Goal

Build a privacy-aware student forms triage and status workflow that:

- Accepts synthetic student form requests
- Classifies request type
- Extracts required fields
- Detects missing information
- Routes cases to the right mock queue
- Shows a learner-facing status timeline
- Provides a reviewer console
- Keeps human review in control
- Simulates Banner, CRM, LMS, and knowledge-base integrations through mock adapters

Phase 1 delivers only the transfer-credit slice of that broader MVP. Other request types remain planned work for later phases.

## Preferred Language

Go is the primary backend and orchestration language.

## Test Layers

The full roadmap uses the layers below. In Phase 1, only the layers needed for the transfer-credit slice are mandatory.

| Layer | Purpose | Phase 1 Requirement |
|---|---|---|
| Unit tests | Domain logic, routing, validation, guardrails | Required |
| Integration tests | PostgreSQL repositories and API handlers | Required |
| Contract tests | Mock SIS/CRM/LMS/knowledge adapters | Deferred until adapters exist |
| End-to-end tests | Learner submission through reviewer visibility | Required |
| Evaluation tests | Classifier and routing accuracy on labelled synthetic data | Deferred until multi-workflow classification exists |
| Accessibility checks | Keyboard, labels, status messages | Deferred until the reviewer and learner UI surface exists |
| Security/privacy tests | Redaction, upload limits, no-auto-decision guardrails | Guardrail tests are required; broader privacy/security suites are deferred until the related surface exists |

## Coverage Gates

Minimum expectations for applicant MVP:

| Scope | Minimum |
|---|---:|
| Domain packages | 90% |
| Rule engine | 95% |
| Adapter contracts | 90% |
| API handlers | 80% |
| Overall backend | 85% |

Coverage is not a substitute for meaningful tests. Any critical guardrail must have explicit tests.

## Files

| File | Purpose |
|---|---|
| `01-tdd-product-scope-and-phases.md` | Product scope and the executable Phase 1 transfer-credit slice |
| `02-tdd-go-architecture-and-implementation.md` | Go architecture, test layout, and Red-Green-Refactor tasks |
| `03-tdd-data-model-and-api-contract.md` | Schema, API, and contract tests |
| `04-tdd-ai-triage-workflow-and-rules.md` | AI/rules workflow with evaluation-driven TDD |
| `05-tdd-privacy-security-accessibility.md` | Guardrail tests for privacy, security, and accessibility |
| `06-tdd-demo-plan-and-portfolio-packaging.md` | Demo packaging with proof of tests |
| `07-tdd-30-60-90-day-build-plan.md` | 30/60/90 plan where every milestone starts with tests |
| `08-tdd-ci-makefile-and-test-skeletons.md` | CI, Makefile targets, and Go test skeletons |
| `09-tdd-phase-1-transfer-credit-vertical-slice.md` | Canonical execution map for the Phase 1 transfer-credit slice |
| `10-tdd-phase-1-traceability-matrix.md` | Phase 1 traceability from scope to tests, fixtures, and demo proof |
| `11-phase-1-bootstrap-implementation-status.md` | Current runnable Phase 1 implementation status and known gaps |
| `testdata/scenarios/README.md` | Synthetic fixture contract for the Phase 1 transfer-credit slice |

## Current Implementation Status

The repo now includes a runnable Go bootstrap for the Phase 1 transfer-credit slice:

- `POST /cases`
- `GET /cases/{case_id}/timeline`
- deterministic transfer-credit routing
- reviewer queue visibility through an internal read model/query
- scenario fixtures, unit tests, tagged E2E coverage, and coverage scripts

The current implementation is intentionally narrow. It proves the transfer-credit flow with an in-memory repository first, while the PostgreSQL migrations and persistence hardening described in the roadmap remain follow-on work.

See `11-phase-1-bootstrap-implementation-status.md` for the current code surface, verified commands, and deferred items.

## Phase 1 Demo Path

- Use `01-tdd-product-scope-and-phases.md` as the source of truth for Phase 1 scope and exit criteria.
- Use `09-tdd-phase-1-transfer-credit-vertical-slice.md` as the ordered build sequence for Phase 1 execution.
- Use `10-tdd-phase-1-traceability-matrix.md` to keep the canonical scenario, tests, and demo proof aligned.
- Demo only the transfer-credit request path in Phase 1.
- Prove the path with failing-first unit, integration, API, and E2E tests.
- Show reviewer visibility through the minimal queue read model/query used by the demo harness.
- Keep reviewer action human-controlled; no approvals or denials are automated.
- Treat other workflows as deferred until Phase 2.

## Definition of Done

A feature is done only when:

- [ ] A failing test was written first.
- [ ] The test failure was observed.
- [ ] The simplest passing code was written.
- [ ] Refactor was completed safely.
- [ ] Unit tests pass.
- [ ] Integration tests pass where relevant.
- [ ] Contract tests pass where relevant.
- [ ] E2E test passes where relevant.
- [ ] Coverage gate passes.
- [ ] Privacy/security/accessibility guardrails still pass.
- [ ] Documentation was updated.
- [ ] If Phase 1 changed, the transfer-credit demo path in README still matches `01-tdd-product-scope-and-phases.md`.
