.PHONY: test-unit test-integration test-e2e test-phase1 coverage-case-domain coverage-phase1 ci-phase1

test-unit:
	go test ./internal/cases ./internal/triage ./internal/reviewerqueue ./internal/platform/httpserver

test-integration:
	go test -tags=integration ./internal/cases ./internal/platform/httpserver

test-e2e:
	go test -tags=e2e ./e2e/...

test-phase1: test-unit test-integration test-e2e

coverage-case-domain:
	bash ./scripts/check-case-domain-coverage.sh

coverage-phase1:
	bash ./scripts/check-phase1-coverage.sh

ci-phase1: test-phase1 coverage-case-domain
