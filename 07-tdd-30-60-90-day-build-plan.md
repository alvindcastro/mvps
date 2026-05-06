# 07 - Strict TDD 30/60/90 Day Build Plan

## Rule

Every milestone starts with tests. If a task cannot be tested, rewrite it until it can.

---

# First 30 Days - Tested Foundation

## Goal

Create the backend foundation and one tested transfer-credit vertical slice.

## Week 1 - TDD Setup

### Red

- [ ] Write failing health endpoint test.
- [ ] Write failing readiness endpoint test.
- [ ] Write failing config test.
- [ ] Write failing logger redaction test.
- [ ] Write failing case creation test name only if needed.

### Green

- [ ] Initialize Go module.
- [ ] Implement health endpoint.
- [ ] Implement readiness endpoint.
- [ ] Implement config.
- [ ] Implement logging.

### Refactor

- [ ] Extract server setup.
- [ ] Add test helpers.
- [ ] Add Makefile test targets.
- [ ] Add CI test job.

### Done

- [ ] `make test` passes.
- [ ] Coverage report exists.
- [ ] README documents TDD rule.

---

# Week 2 - Case Creation

### Red

- [ ] Write failing case domain validation tests.
- [ ] Write failing case repository tests.
- [ ] Write failing `POST /cases` API tests.
- [ ] Write failing timeline event test.

### Green

- [ ] Implement case model.
- [ ] Implement case repository.
- [ ] Implement case API.
- [ ] Implement status timeline.

### Refactor

- [ ] Extract case service.
- [ ] Add table-driven validation tests.
- [ ] Add fixture builders.

### Done

- [ ] Case creation works.
- [ ] Timeline is created.
- [ ] Repository integration tests pass.

---

# Week 3 - Transfer Credit Rules

### Red

- [ ] Write failing complete transfer-credit route test.
- [ ] Write failing missing-field tests.
- [ ] Write failing missing-document test.
- [ ] Write failing no-auto-approval test.
- [ ] Write failing no-auto-denial test.

### Green

- [ ] Implement required fields.
- [ ] Implement route suggestion.
- [ ] Implement missing-information response.
- [ ] Implement guardrail rules.

### Refactor

- [ ] Extract rule engine.
- [ ] Add route reason codes.
- [ ] Add table-driven route tests.

### Done

- [ ] Rule package coverage >= 95%.
- [ ] No-auto-decision tests pass.

---

# Week 4 - Reviewer and Demo Slice

### Red

- [ ] Write failing reviewer queue test.
- [ ] Write failing reviewer decision test.
- [ ] Write failing learner timeline update test.
- [ ] Write failing E2E transfer-credit test.

### Green

- [ ] Implement reviewer list endpoint.
- [ ] Implement reviewer decision endpoint.
- [ ] Update timeline on review.
- [ ] Add minimal UI or API demo script.

### Refactor

- [ ] Extract review service.
- [ ] Add authorization placeholder tests.
- [ ] Add demo seed data.

### 30-Day Done

- [ ] One complete workflow works.
- [ ] Tests prove the workflow.
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
