# Repository Guidelines

## Project Structure & Module Organization

This repository is currently a documentation-first MVP package. The root contains the canonical numbered design and TDD specs such as [README.md](/mnt/c/Users/alvin/GolandProjects/mvps/README.md), `01-tdd-product-scope-and-phases.md`, `02-tdd-go-architecture-and-implementation.md`, and `08-tdd-ci-makefile-and-test-skeletons.md`.

Planned implementation should follow the Go layout defined in `02-tdd-go-architecture-and-implementation.md`: `cmd/api/`, `internal/cases/`, `internal/triage/`, `internal/platform/`, `e2e/`, `migrations/`, and `testdata/scenarios/`. Keep Phase 1 limited to the transfer-credit vertical slice.

## Build, Test, and Development Commands

No executable build system is checked in yet. Until code lands, use the docs as the source of truth and keep changes scoped to the MVP slice.

When implementation begins, the expected commands are:

- `make test-unit` for domain and handler unit tests
- `make test-integration` for PostgreSQL-backed integration tests
- `make test-e2e` for the learner-to-reviewer journey
- `make ci` for lint, tests, and coverage gates
- `golangci-lint run` for linting

## Coding Style & Naming Conventions

Use Go as the primary implementation language. Follow standard Go formatting with tabs via `gofmt`. Keep package names lowercase (`cases`, `triage`, `reviewerqueue`) and prefer explicit, behavior-driven test names such as `TestCreateTransferCreditCase_WithMissingPriorInstitution_ReturnsValidationError`.

For docs, keep filenames descriptive and stable. Continue the existing numbered spec pattern for new top-level planning documents only when they extend the same sequence.

## Testing Guidelines

This repo follows a strict failing-test-first rule: no production code before a failing test exists. Phase 1 requires unit, integration, and E2E coverage for the transfer-credit workflow. Use Go build tags for slower suites, for example `//go:build integration` and `//go:build e2e`.

Honor the documented coverage gates: 90% for `internal/cases` and 85% overall backend coverage.

## Commit & Pull Request Guidelines

Recent history uses short conventional commits such as `docs: ...` and `chore: ...`; keep that format. PRs should describe the Phase 1 scenario affected, list tests run, and note any required doc sync in `README.md` and the numbered TDD specs. If behavior or flow changes, include the updated demo path and impacted files.
