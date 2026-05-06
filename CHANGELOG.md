# Changelog

All notable implementation changes for this repository should be recorded here.

## [Unreleased]

### Phase 1 Transfer-Credit Bootstrap

- What: Added the first runnable Go implementation for the Phase 1 transfer-credit slice, including transfer-credit validation, deterministic routing, adverse-decision guardrails, case creation, submitted timeline events, reviewer-queue visibility, scenario fixtures, an OpenAPI stub, and test/coverage scripts.
- Who: Implemented by Codex in the local workspace, with parallel subagents used for Phase 1 scope review and fixture-contract preparation.
- Where: Added or updated `go.mod`, `Makefile`, `cmd/api/`, `internal/cases/`, `internal/triage/`, `internal/reviewerqueue/`, `internal/platform/httpserver/`, `internal/testsupport/`, `e2e/`, `docs/openapi-phase1.yaml`, `scripts/`, `testdata/scenarios/`, `README.md`, and `11-phase-1-bootstrap-implementation-status.md`.
- When: 2026-05-05.
- How: Built the narrowest executable slice from the Phase 1 markdown specs using a stdlib-first Go bootstrap with an in-memory repository, then verified it with unit tests, tagged integration tests, tagged E2E tests, and coverage gates.
- Why: The repo previously described the transfer-credit MVP only as planning artifacts. This change makes the first logical spec executable while preserving the broader PostgreSQL-backed roadmap for later hardening.

### Verification

- `go test ./internal/cases ./internal/triage ./internal/reviewerqueue ./internal/platform/httpserver`
- `go test -tags=integration ./internal/cases ./internal/platform/httpserver`
- `go test -tags=e2e ./e2e/...`
- `bash ./scripts/check-case-domain-coverage.sh`
- `bash ./scripts/check-phase1-coverage.sh`

### Suggested Commit

- `feat: add phase 1 transfer-credit bootstrap slice`
