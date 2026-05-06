# 07 - Strict TDD 30/60/90 Day Build Plan

## Rule

Every milestone starts with tests. If a task cannot be tested, rewrite it until it can.

---

# First 30 Days - Tested Foundation

## Goal

Create one tested transfer-credit vertical slice first, then harden the platform around that slice.

## Week 1 - Transfer Credit Domain and Routing

### Red

- [ ] Write failing transfer-credit validation tests.
- [ ] Write failing transfer-credit route tests.
- [ ] Write failing complete-input case-creation test.
- [ ] Write failing no-auto-approval test.
- [ ] Write failing no-auto-denial test.

### Green

- [ ] Initialize the Go module and minimal package layout.
- [ ] Implement transfer-credit required-field validation.
- [ ] Implement deterministic route suggestion for `registrar_transfer_credit`.
- [ ] Implement guardrail rules that force human review for approval or denial outcomes.

### Refactor

- [ ] Convert validation and route tests to table-driven form.
- [ ] Add structured scenario fixtures.
- [ ] Extract route reason codes.
- [ ] Add initial Makefile targets for the slice.

### Done

- [ ] Transfer-credit validation and route tests pass.
- [ ] Guardrail tests pass.
- [ ] README documents the Phase 1 demo path.

---

## Week 2 - Persistence, API, and Timeline

### Red

- [ ] Write failing case repository tests.
- [ ] Write failing `POST /cases` API tests.
- [ ] Write failing `GET /cases/{id}/timeline` API test.
- [ ] Write failing timeline event transaction test.

### Green

- [ ] Implement case model.
- [ ] Implement case repository.
- [ ] Implement case API.
- [ ] Implement status timeline.
- [ ] Add `cases` and `status_events` migrations only.

### Refactor

- [ ] Extract case service.
- [ ] Add repository and handler test helpers.
- [ ] Add fixture builders for transfer-credit requests.

### Done

- [ ] Case creation works for complete transfer-credit input.
- [ ] Timeline is created transactionally with case creation.
- [ ] Repository integration tests pass.

---

## Week 3 - Reviewer Queue and E2E Slice

### Red

- [ ] Write failing reviewer queue visibility test.
- [ ] Write failing learner-to-reviewer E2E transfer-credit test.
- [ ] Write failing submitted-status timeline assertion in the E2E path.

### Green

- [ ] Implement the minimal reviewer queue read model.
- [ ] Expose the queue query or view needed for the demo harness.
- [ ] Add a minimal demo script or walkthrough data.

### Refactor

- [ ] Extract reviewer queue query helpers.
- [ ] Remove duplication between API and E2E fixtures.
- [ ] Tighten status naming across learner and reviewer views.

### Done

- [ ] The reviewer queue shows new transfer-credit cases.
- [ ] The end-to-end transfer-credit workflow passes.
- [ ] No automated approval or denial is present anywhere in the slice.

---

## Week 4 - Platform Hardening After the Slice

### Red

- [ ] Write failing health endpoint test.
- [ ] Write failing readiness endpoint test.
- [ ] Write failing config test.
- [ ] Write failing logger redaction test.

### Green

- [ ] Implement health endpoint.
- [ ] Implement readiness endpoint.
- [ ] Implement config loading.
- [ ] Implement logging redaction.

### Refactor

- [ ] Extract server setup.
- [ ] Add test cleanup helpers.
- [ ] Add CI job ordering that matches the Phase 1 slice.

### 30-Day Done

- [ ] One complete transfer-credit workflow works.
- [ ] Tests prove the workflow in unit, integration, and E2E layers.
- [ ] README has demo instructions.
- [ ] Docker Compose starts the system.

---

# Days 31 to 60 - Tested Intelligence and Adapters

## Goal

Add multi-workflow triage, mock enterprise adapters, and safe AI abstraction.

## Weeks 5-6 - Multi-Workflow Rules

### Red

- [ ] Write failing prerequisite waiver tests.
- [ ] Write failing refund tests.
- [ ] Write failing unknown request tests.
- [ ] Write failing low-confidence tests.
- [ ] Write failing duplicate submission tests.

### Green

- [ ] Add workflow-specific rules.
- [ ] Add deterministic classifier.
- [ ] Add confidence thresholds.
- [ ] Add duplicate detection.

### Refactor

- [ ] Convert workflow rules to config table.
- [ ] Add shared scenario fixtures.
- [ ] Add regression tests.

### Done

- [ ] All three workflows pass.
- [ ] Unknowns fallback safely.
- [ ] Duplicate detection works.

---

# Weeks 7-8 - Mock Adapters

### Red

- [ ] Write failing SIS contract tests.
- [ ] Write failing CRM contract tests.
- [ ] Write failing LMS notification contract tests.
- [ ] Write failing knowledge adapter contract tests.
- [ ] Write failing adapter logging redaction tests.

### Green

- [ ] Implement mock SIS.
- [ ] Implement mock CRM.
- [ ] Implement mock LMS.
- [ ] Implement mock knowledge adapter.
- [ ] Implement adapter call logging.

### Refactor

- [ ] Extract adapter interfaces.
- [ ] Add shared contract test suite.
- [ ] Add failure simulation.

### Done

- [ ] Contract tests prove replaceability.
- [ ] Adapter failures are testable.
- [ ] Domain code depends only on interfaces.

---

# Weeks 9 - Outbox and Worker

### Red

- [ ] Write failing outbox event creation test.
- [ ] Write failing worker success test.
- [ ] Write failing retry test.
- [ ] Write failing dead-letter test.
- [ ] Write failing idempotent processing test.

### Green

- [ ] Implement outbox.
- [ ] Implement worker.
- [ ] Implement retry policy.
- [ ] Implement dead-letter status.
- [ ] Implement idempotent event handling.

### Refactor

- [ ] Add fake clock.
- [ ] Separate polling and processing.
- [ ] Extract retry calculator.

### 60-Day Done

- [ ] Multi-workflow triage works.
- [ ] Mock integrations work.
- [ ] Outbox worker works.
- [ ] Contract tests pass.
- [ ] Integration tests pass.

---

# Days 61 to 90 - Tested Operations and Portfolio

## Goal

Make the project operational, measurable, and interview-ready.

## Weeks 10-11 - Evaluation and Metrics

### Red

- [ ] Write failing evaluation fixture loader test.
- [ ] Write failing intent accuracy gate test.
- [ ] Write failing false auto-route gate test.
- [ ] Write failing metrics summary test.
- [ ] Write failing p95 latency metric test.

### Green

- [ ] Add labelled evaluation fixtures.
- [ ] Add evaluation runner.
- [ ] Add metrics endpoint.
- [ ] Add metrics dashboard.
- [ ] Add Markdown evaluation report.

### Refactor

- [ ] Add confusion matrix.
- [ ] Add failing examples section.
- [ ] Add regression fixture workflow.

### Done

- [ ] `make eval` passes.
- [ ] Metrics are visible.
- [ ] Evaluation report is generated.

---

# Week 12 - Privacy, Security, Accessibility

### Red

- [ ] Write failing synthetic-data acknowledgement test.
- [ ] Write failing upload limit tests.
- [ ] Write failing role access tests.
- [ ] Write failing prompt redaction test.
- [ ] Write failing accessibility smoke tests.

### Green

- [ ] Add privacy notice.
- [ ] Add upload restrictions.
- [ ] Add demo auth roles.
- [ ] Add redaction.
- [ ] Add accessible labels and errors.

### Refactor

- [ ] Extract privacy policy.
- [ ] Extract upload policy.
- [ ] Extract authorization policy.
- [ ] Add accessibility checklist.

### Done

- [ ] Guardrail tests pass.
- [ ] Accessibility checklist is documented.
- [ ] Security limitations are documented.

---

# Week 13 - Portfolio Packaging

### Red

- [ ] Write failing docs check for required README sections.
- [ ] Write failing script check for demo fixture availability.
- [ ] Write failing smoke test for Docker startup.

### Green

- [ ] Finish README.
- [ ] Add screenshots.
- [ ] Add diagrams.
- [ ] Add demo script.
- [ ] Add CI badge.
- [ ] Add coverage report.
- [ ] Add evaluation report.

### Refactor

- [ ] Simplify demo flow.
- [ ] Remove dead docs.
- [ ] Polish naming.

### 90-Day Done

- [ ] All tests pass.
- [ ] Coverage gates pass.
- [ ] Evaluation gates pass.
- [ ] Docker startup works.
- [ ] Portfolio README is complete.
- [ ] Project is ready to share.
