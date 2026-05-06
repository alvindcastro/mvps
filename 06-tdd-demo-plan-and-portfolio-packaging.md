# 06 - TDD Demo Plan and Portfolio Packaging

## Demo Principle

The demo should prove not only that the app works, but that it was built with strict TDD.

Show tests, coverage, and guardrails as part of the portfolio.

## 5-Minute Demo Structure

1. Show README and TDD rule.
2. Run tests.
3. Run coverage.
4. Start app.
5. Submit a synthetic transfer-credit case.
6. Show route and timeline.
7. Submit a low-confidence case.
8. Show human review.
9. Submit duplicate case.
10. Show idempotency.
11. Show metrics dashboard.
12. Show privacy/security/accessibility tests.

## Required Demo Commands

```bash
make test
make test-integration
make test-contract
make test-e2e
make coverage
make eval
docker compose up
```

---

# Phase 1 - TDD Evidence Packaging

## Tickable Tasks

- [ ] Add `docs/tdd-strategy.md`.
- [ ] Add `docs/test-plan.md`.
- [ ] Add `docs/coverage-report.md`.
- [ ] Add `docs/evaluation-report.md`.
- [ ] Add screenshot of tests passing.
- [ ] Add screenshot of coverage summary.
- [ ] Add CI badge.
- [ ] Add test command examples.
- [ ] Add explanation of Red-Green-Refactor.
- [ ] Add example failing-test-first commit if possible.

## Acceptance Criteria

- [ ] Interviewer can see TDD discipline immediately.
- [ ] Test commands are documented.
- [ ] Coverage gates are documented.
- [ ] Evaluation report is included.

---

# Phase 2 - README Job Alignment

## Tickable Tasks

- [ ] Add section: Role alignment.
- [ ] Add section: Why Go.
- [ ] Add section: Why strict TDD.
- [ ] Add section: AI use boundaries.
- [ ] Add section: Workflow automation.
- [ ] Add section: Enterprise adapter simulation.
- [ ] Add section: Privacy and accessibility.
- [ ] Add section: Monitoring and metrics.
- [ ] Add section: What I would do with real OC access.

## Role Alignment Text

```text
This MVP demonstrates AI-assisted triage, workflow automation, API-first design, enterprise integration readiness, privacy-aware AI design, and monitoring. It uses Go for the orchestration layer and strict TDD to show that learner-facing automation can be developed safely and predictably.
```

## Acceptance Criteria

- [ ] README maps directly to job responsibilities.
- [ ] README does not overclaim production readiness.
- [ ] README explains synthetic-data limitations.

---

# Phase 3 - Demo Scenarios as Tests

## Tickable Tasks

- [ ] Create test fixture for complete transfer-credit request.
- [ ] Create test fixture for missing transcript.
- [ ] Create test fixture for ambiguous request.
- [ ] Create test fixture for prerequisite waiver.
- [ ] Create test fixture for refund request.
- [ ] Create test fixture for duplicate submission.
- [ ] Create test fixture for approval-language guardrail.
- [ ] Create test fixture for denial-language guardrail.
- [ ] Add E2E test for each demo scenario.
- [ ] Add expected screenshots or snapshots where useful.

## Demo Scenario Table

| Scenario | Required Test |
|---|---|
| Complete transfer credit | E2E creates and routes case |
| Missing transcript | E2E shows needs-more-info |
| Ambiguous request | E2E routes to manual triage |
| Duplicate submission | E2E returns existing case |
| Approval request | Guardrail test requires human review |
| Refund outcome | Guardrail test requires human review |

## Acceptance Criteria

- [ ] Every demo scenario has a test.
- [ ] Demo cannot drift away from tested behaviour.
- [ ] Failed demo scenario fails CI before interview.

---

# Phase 4 - Portfolio Visuals

## Tickable Tasks

- [ ] Add architecture diagram.
- [ ] Add Red-Green-Refactor diagram.
- [ ] Add workflow sequence diagram.
- [ ] Add data model diagram.
- [ ] Add adapter contract diagram.
- [ ] Add test pyramid diagram.
- [ ] Add CI pipeline diagram.
- [ ] Add screenshots of UI.
- [ ] Add screenshots of reviewer console.
- [ ] Add screenshots of metrics.
- [ ] Add screenshot of test output.
- [ ] Add screenshot of coverage output.

## Acceptance Criteria

- [ ] Visuals support technical explanation.
- [ ] TDD process is visible.
- [ ] Architecture and tests are connected.

---

# Phase 5 - Interview Talking Points

## TDD Talking Points

- [ ] I wrote failing tests before production code.
- [ ] I used unit tests for deterministic domain logic.
- [ ] I used integration tests for repositories and API handlers.
- [ ] I used contract tests for mock enterprise adapters.
- [ ] I used E2E tests for learner and reviewer journeys.
- [ ] I used evaluation fixtures for AI/routing quality.
- [ ] I used guardrail tests to prevent unsafe automation.
- [ ] I used coverage gates but did not rely on coverage alone.

## Go Talking Points

- [ ] Go kept the orchestration layer simple and testable.
- [ ] Interfaces made enterprise adapters replaceable.
- [ ] Table-driven tests fit the rules engine well.
- [ ] The outbox worker is easy to test with fake clocks and fake adapters.
- [ ] The modular monolith avoids unnecessary microservice complexity.

## Product Talking Points

- [ ] The learner gets instant status.
- [ ] Staff get structured case summaries.
- [ ] Duplicate email-driven work is reduced.
- [ ] The system catches missing information earlier.
- [ ] Human review remains in control.

## Governance Talking Points

- [ ] Synthetic data only.
- [ ] No auto-approve.
- [ ] No auto-deny.
- [ ] Redaction before logs or model calls.
- [ ] FIPPA-aware posture.
- [ ] Accessibility included in definition of done.

---

# Phase 6 - Final Release Checklist

## Tickable Tasks

- [ ] `make test` passes.
- [ ] `make test-integration` passes.
- [ ] `make test-contract` passes.
- [ ] `make test-e2e` passes.
- [ ] `make eval` passes.
- [ ] Coverage gate passes.
- [ ] Race detector passes where practical.
- [ ] Secret scan passes.
- [ ] Docker startup works.
- [ ] README is updated.
- [ ] Demo script is updated.
- [ ] Screenshots are updated.
- [ ] No real personal information exists in repo.
- [ ] Known limitations are documented.
- [ ] Future roadmap is documented.

## Definition of Portfolio Ready

- [ ] A recruiter can understand the project in 3 minutes.
- [ ] A technical reviewer can run tests in 5 minutes.
- [ ] A hiring manager can see role alignment.
- [ ] The project demonstrates disciplined, safe automation.
