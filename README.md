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
-> reviewer sees the queued case
-> learner sees a submitted status timeline
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

| Layer | Purpose | Required Before Code |
|---|---|---|
| Unit tests | Domain logic, rules, classification, validation | Yes |
| Integration tests | PostgreSQL repositories, outbox, API handlers | Yes |
| Contract tests | Mock SIS/CRM/LMS/knowledge adapters | Yes |
| End-to-end tests | Learner and reviewer workflows | Yes |
| Evaluation tests | Classifier and routing accuracy on labelled synthetic data | Yes |
| Accessibility checks | Keyboard, labels, status messages | Yes |
| Security/privacy tests | Redaction, upload limits, no-auto-decision guardrails | Yes |

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

## Phase 1 Demo Path

- Use `01-tdd-product-scope-and-phases.md` as the source of truth for Phase 1 scope and exit criteria.
- Use `09-tdd-phase-1-transfer-credit-vertical-slice.md` as the ordered build sequence for Phase 1 execution.
- Demo only the transfer-credit request path in Phase 1.
- Prove the path with failing-first unit, integration, API, and E2E tests.
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
