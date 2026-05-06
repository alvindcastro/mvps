# 11 - Phase 1 Bootstrap Implementation Status

## Purpose

This document records the current runnable implementation that now exists alongside the Phase 1 planning docs.

Use it as the source of truth for what is actually executable in the repo today, without rewriting the broader roadmap documents that still describe the intended PostgreSQL-backed build-out.

## Implemented Bootstrap Slice

The current codebase implements the smallest runnable version of the Phase 1 transfer-credit path:

- validate `transfer_credit` submissions for `learner_ref`, `term`, `prior_institution`, `prior_course_code`, and `target_program`
- deterministically route complete requests to `registrar_transfer_credit`
- require human review for approval or denial requests
- create a submitted case with a stable `case_id` and `case_number`
- append one submitted timeline event
- expose `POST /cases`
- expose `GET /cases/{case_id}/timeline`
- surface reviewer visibility through the internal reviewer-queue query

## Implemented Files

- `go.mod`
- `cmd/api/main.go`
- `internal/cases/`
- `internal/triage/`
- `internal/reviewerqueue/`
- `internal/platform/httpserver/`
- `internal/testsupport/`
- `e2e/transfer_credit_submission_test.go`
- `Makefile`
- `scripts/check-case-domain-coverage.sh`
- `scripts/check-phase1-coverage.sh`
- `docs/openapi-phase1.yaml`
- `testdata/scenarios/*.json`

## Verified Commands

The following commands were run successfully against the current bootstrap:

- `go test ./internal/cases ./internal/triage ./internal/reviewerqueue ./internal/platform/httpserver`
- `go test -tags=integration ./internal/cases ./internal/platform/httpserver`
- `go test -tags=e2e ./e2e/...`
- `bash ./scripts/check-case-domain-coverage.sh`
- `bash ./scripts/check-phase1-coverage.sh`

## Current Gaps Relative to the Roadmap

These items are still intentionally deferred:

- PostgreSQL-backed repository and real migrations
- transactional database coverage against a live datastore
- broader workflow support beyond transfer credit
- adapter contracts, outbox work, and reviewer decisioning

## Current Coverage Snapshot

- `internal/cases`: 91.8%
- broader Phase 1 slice gate: 89.2%

## Notes

This bootstrap keeps the Phase 1 product behavior narrow and testable while avoiding premature platform work. The next logical hardening step is replacing the in-memory repository with the PostgreSQL-backed implementation already described in `03-tdd-data-model-and-api-contract.md`.
