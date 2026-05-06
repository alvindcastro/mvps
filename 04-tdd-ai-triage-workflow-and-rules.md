# 04 - TDD AI Triage Workflow and Rules

## AI TDD Rule

AI behaviour must be tested through deterministic interfaces and labelled fixtures.

Do not test by hoping a live model responds correctly. Live model calls are optional smoke tests only.

## AI Responsibilities

Allowed:

- Classify request intent
- Extract structured fields
- Suggest route
- Generate plain-language status text
- Summarize missing information

Not allowed:

- Approve requests
- Deny requests
- Promise refunds
- Decide eligibility
- Process real student data

## Phase 1 - Transfer Credit Deterministic Routing Slice

### Scope Rule

Phase 1 starts from the transfer-credit submission path, not from open-ended classification.

The first runnable slice must prove deterministic routing, required-field validation, reviewer visibility, and no-decision guardrails from structured transfer-credit input plus labelled fixtures.

Defer duplicate handling, extraction conflicts, unsupported mappings, and low-confidence fallbacks until later phases unless `01-tdd-product-scope-and-phases.md` explicitly expands the slice.

### Route Outcomes Allowed in Phase 1

| Condition | Expected Action |
|---|---|
| Complete transfer-credit request | Route to `registrar_transfer_credit` |
| Missing required field | Return a validation error before case creation |
| Learner asks for approval or denial outcome | Return `human_review_required` |
| Any automated approval path | Disallowed |
| Any automated denial path | Disallowed |

### Red Tests

- [ ] `TestTransferCreditTriage_CompleteStructuredInput_SuggestsRegistrarTransferCredit`
- [ ] `TestTransferCreditTriage_MissingPriorInstitution_ReturnsValidationError`
- [ ] `TestTransferCreditTriage_AdverseDecisionRequest_RequiresHumanReview`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoApproves`
- [ ] `TestAdverseDecisionGuardrail_NeverAutoDenies`

### Green Tasks

- [ ] Define deterministic `TransferCreditTriageInput`.
- [ ] Implement required-field checks for transfer credit.
- [ ] Implement deterministic route decision rules.
- [ ] Implement guardrail rules for no auto-approval and no auto-denial.
- [ ] Return route reason codes and missing-field list.
- [ ] Return a reviewer-safe explanation for every guardrail-triggered result.

### Refactor Tasks

- [ ] Convert transfer-credit rules to table-driven fixtures.
- [ ] Extract guardrail policy from route policy.
- [ ] Add reusable structured-input fixture builders.
- [ ] Add route matrix tests for all Phase 1 outcomes.

### Acceptance Criteria

- [ ] A complete transfer-credit request deterministically routes to `registrar_transfer_credit`.
- [ ] Missing required information never reaches the happy-path queue.
- [ ] Learner-impacting approval or denial outcomes always require human review.
- [ ] No approval or denial outcome is automated anywhere in the slice.
- [ ] No live model is required for CI or the first demo slice.

---

## Phase 2 - Workflow Classification Expansion

### Red Tests

- [ ] `TestClassifier_TextMentionsTransferCredit_ReturnsTransferCredit`
- [ ] `TestClassifier_TextMentionsPrerequisiteWaiver_ReturnsPrereqWaiver`
- [ ] `TestClassifier_TextMentionsRefund_ReturnsRefundRequest`
- [ ] `TestClassifier_AmbiguousText_ReturnsUnknownLowConfidence`
- [ ] `TestClassifier_EmptyText_ReturnsValidationError`

### Green Tasks

- [ ] Implement classifier interface.
- [ ] Implement deterministic keyword classifier.
- [ ] Return confidence score.
- [ ] Return reason codes.
- [ ] Return a workflow suggestion without bypassing Phase 1 guardrails.

### Refactor Tasks

- [ ] Convert keyword rules to table-driven config.
- [ ] Add multilingual or synonym fixtures later.
- [ ] Keep classifier deterministic for baseline tests.

### Acceptance Criteria

- [ ] Baseline classifier passes all tests.
- [ ] Unknown requests safely route to manual review.
- [ ] Confidence is explainable.
- [ ] Transfer-credit routing continues to pass unchanged.

---

## Phase 3 - Extraction TDD

## Red Tests

- [ ] `TestExtractor_StructuredTransferCreditInput_ExtractsPriorInstitution`
- [ ] `TestExtractor_StructuredTransferCreditInput_ExtractsPriorCourseCode`
- [ ] `TestExtractor_MissingCourseCode_ReturnsMissingField`
- [ ] `TestExtractor_ConflictingEnteredAndExtractedValue_FlagsConflict`
- [ ] `TestExtractor_LowConfidenceField_RequiresHumanReview`

## Green Tasks

- [ ] Implement extraction interface.
- [ ] Implement structured input extractor.
- [ ] Implement mock document extractor.
- [ ] Add field confidence.
- [ ] Add conflict detection.
- [ ] Add missing-field detection.

## Refactor Tasks

- [ ] Extract field schema by workflow.
- [ ] Add reusable fixture builder.
- [ ] Add table-driven extraction tests.

## Acceptance Criteria

- [ ] Extraction is deterministic in tests.
- [ ] Missing fields are explicit.
- [ ] Conflicts trigger human review.
- [ ] Phase 1 transfer-credit routing still remains deterministic.

---

## Phase 4 - LLM Provider Behind Tests

## Red Tests

- [ ] `TestLLMClassifier_WithValidJSON_ReturnsClassification`
- [ ] `TestLLMClassifier_WithInvalidJSON_FallsBackSafely`
- [ ] `TestLLMClassifier_WithUnsupportedFormType_ReturnsManualReview`
- [ ] `TestLLMClassifier_RedactsLearnerRefBeforePrompt`
- [ ] `TestLLMClassifier_StoresPromptVersion`

## Green Tasks

- [ ] Define `LLMProvider` interface.
- [ ] Implement fake LLM provider.
- [ ] Implement JSON schema validation.
- [ ] Implement fallback path.
- [ ] Implement redaction before prompt.
- [ ] Store prompt version and model version.

## Refactor Tasks

- [ ] Separate prompt builder from classifier.
- [ ] Add golden prompt tests.
- [ ] Add prompt version constants.
- [ ] Keep live provider optional.

## Acceptance Criteria

- [ ] Invalid model output fails safely.
- [ ] Redaction happens before provider call.
- [ ] Prompt version is logged.
- [ ] Live model is not required for CI.
- [ ] The deterministic Phase 1 path remains the default fallback.

---

## Phase 5 - Cross-Workflow Routing Rules TDD

### Routing Thresholds After Phase 1

| Condition | Expected Action |
|---|---|
| Confidence >= 0.90 and complete | Auto-route to queue |
| Confidence 0.70 to 0.89 | Reviewer confirmation |
| Confidence < 0.70 | Manual triage |
| Missing fields | Needs more information |
| Conflicting values | Human review |
| Learner-impacting outcome | Human review |
| Duplicate submission | Return existing case |

## Red Tests

- [ ] `TestRouting_HighConfidenceCompleteCase_AutoRoutes`
- [ ] `TestRouting_MediumConfidenceCase_RequiresReviewerConfirmation`
- [ ] `TestRouting_LowConfidenceCase_RoutesManualTriage`
- [ ] `TestRouting_MissingFields_NeedsMoreInfo`
- [ ] `TestRouting_ConflictingValues_RequiresHumanReview`
- [ ] `TestRouting_LearnerImpactingOutcome_RequiresHumanReview`
- [ ] `TestRouting_DuplicateSubmission_ReturnsExistingCase`

## Green Tasks

- [ ] Implement route decision engine across workflows.
- [ ] Implement confidence thresholds.
- [ ] Implement completeness rules.
- [ ] Implement conflict rules.
- [ ] Implement adverse decision rules.
- [ ] Implement duplicate routing response.

## Refactor Tasks

- [ ] Extract route policy.
- [ ] Extract guardrail policy.
- [ ] Add route reason code constants.
- [ ] Add route matrix tests.

### Coverage Gate

- [ ] Routing package coverage >= 95%.

---

## Phase 6 - Evaluation-Driven TDD

## Evaluation Dataset

Create labelled synthetic examples before tuning logic.

## Required Fixture Fields

```json
{
  "id": "eval-001",
  "input_text": "I submitted my transcript and want transfer credit reviewed.",
  "expected_form_type": "transfer_credit",
  "expected_route": "registrar_transfer_credit",
  "expected_requires_human_review": false,
  "expected_missing_fields": []
}
```

## Red Tests

- [ ] `TestEvaluationDataset_LoadsAllFixtures`
- [ ] `TestEvaluation_BaselineIntentAccuracy_AboveThreshold`
- [ ] `TestEvaluation_FalseAutoRouteRate_BelowThreshold`
- [ ] `TestEvaluation_UnsupportedRequests_FallbackToManualReview`
- [ ] `TestEvaluation_NoAutoApprovalOrDenial`

## Green Tasks

- [ ] Create evaluation fixture loader.
- [ ] Create evaluation runner.
- [ ] Compute intent accuracy.
- [ ] Compute route accuracy.
- [ ] Compute false auto-route rate.
- [ ] Compute missing-field accuracy.
- [ ] Generate Markdown report.

## Refactor Tasks

- [ ] Add confusion matrix.
- [ ] Add failing examples section.
- [ ] Add regression fixtures for bugs.
- [ ] Add CI evaluation target.

## Minimum Evaluation Gates

| Metric | Gate |
|---|---:|
| Intent accuracy | >= 0.90 |
| Missing-field accuracy | >= 0.95 |
| Unsupported safe fallback | 1.00 |
| False auto-route rate | <= 0.02 |
| Auto approval/denial | 0 cases |

## Acceptance Criteria

- [ ] Evaluation runs in CI.
- [ ] Report is generated.
- [ ] Any failed gate blocks merge.
- [ ] Bugs become new fixtures before fix.
