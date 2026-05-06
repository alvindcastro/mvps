# 05 - TDD Privacy, Security, and Accessibility

## Guardrail Rule

Privacy, security, and accessibility requirements must be enforced with tests, not just documented.

## Phase 1 - Synthetic Data Guardrails

### Red Tests

- [ ] `TestSubmission_RealLookingSIN_IsRejected`
- [ ] `TestSubmission_RealLookingEmail_ShowsWarningOrRejectsDependingMode`
- [ ] `TestSubmission_RequiresSyntheticDataAcknowledgement`
- [ ] `TestUploadedDocument_WithoutSyntheticWatermark_IsRejectedInDemoMode`

### Green Tasks

- [ ] Add synthetic-data acknowledgement.
- [ ] Add demo-mode validation.
- [ ] Add synthetic fixture generator.
- [ ] Add upload watermark check placeholder.
- [ ] Add warnings in UI/API response.

### Refactor Tasks

- [ ] Move demo-mode checks to policy package.
- [ ] Add configurable strictness.
- [ ] Add table-driven sensitive-pattern tests.

### Acceptance Criteria

- [ ] Demo strongly discourages real data.
- [ ] Sensitive-looking values are detected.
- [ ] Tests prove synthetic-only posture.

---

# Phase 2 - No-Auto-Decision Guardrails

## Red Tests

- [ ] `TestGuardrail_ApprovalLanguage_RequiresHumanReview`
- [ ] `TestGuardrail_DenialLanguage_RequiresHumanReview`
- [ ] `TestGuardrail_RefundOutcomeLanguage_RequiresHumanReview`
- [ ] `TestGuardrail_EligibilityLanguage_RequiresHumanReview`
- [ ] `TestGuardrail_AutoApproval_IsImpossibleThroughRouteEngine`
- [ ] `TestGuardrail_AutoDenial_IsImpossibleThroughRouteEngine`

## Green Tasks

- [ ] Add adverse-decision detector.
- [ ] Add protected route outcomes.
- [ ] Add human-review requirement.
- [ ] Remove any auto-final-decision path.
- [ ] Add learner-facing disclaimer.

## Refactor Tasks

- [ ] Centralize guardrail logic.
- [ ] Add reason codes.
- [ ] Add regression tests for every new decision-like phrase.

## Acceptance Criteria

- [ ] No route result can equal approved or denied.
- [ ] Learner-impacting requests always require human review.
- [ ] Tests block accidental future regression.

---

# Phase 3 - Logging and Redaction

## Red Tests

- [ ] `TestLogger_RedactsLearnerRef`
- [ ] `TestLogger_DoesNotLogUploadedFileContents`
- [ ] `TestAdapterCallLog_RedactsRequestPayload`
- [ ] `TestPromptBuilder_RedactsSensitiveFields`
- [ ] `TestAuditLog_StoresReasonCodesNotRawSensitiveText`

## Green Tasks

- [ ] Add redaction function.
- [ ] Add log wrapper.
- [ ] Add adapter log sanitizer.
- [ ] Add prompt redaction.
- [ ] Add audit reason codes.

## Refactor Tasks

- [ ] Make redaction policy table-driven.
- [ ] Add fuzz tests for redaction.
- [ ] Add tests for nested JSON redaction.

## Acceptance Criteria

- [ ] Sensitive fields do not appear in logs.
- [ ] Prompt payload is redacted.
- [ ] Audit events are useful but minimized.

---

# Phase 4 - Upload Security

## Red Tests

- [ ] `TestUpload_UnsupportedMimeType_Returns415`
- [ ] `TestUpload_TooLarge_Returns413`
- [ ] `TestUpload_EmptyFile_Returns400`
- [ ] `TestUpload_DuplicateHash_FlagsDuplicate`
- [ ] `TestUpload_FileNameTraversal_IsSanitized`

## Green Tasks

- [ ] Add file size limit.
- [ ] Add MIME type allowlist.
- [ ] Add server-generated file names.
- [ ] Add SHA-256 hashing.
- [ ] Add duplicate file detection.
- [ ] Add sanitized storage path.

## Refactor Tasks

- [ ] Extract upload policy.
- [ ] Add table-driven upload tests.
- [ ] Add fake storage implementation.

## Acceptance Criteria

- [ ] Unsafe uploads are rejected.
- [ ] Duplicate uploads are detected.
- [ ] File names cannot control storage paths.

---

# Phase 5 - Role and Access Tests

## Red Tests

- [ ] `TestReviewerEndpoint_WithoutReviewerRole_ReturnsForbidden`
- [ ] `TestLearnerEndpoint_WithCaseOwner_ReturnsCase`
- [ ] `TestLearnerEndpoint_WithDifferentLearner_ReturnsForbidden`
- [ ] `TestAdminMetrics_WithoutAdminRole_ReturnsForbidden`

## Green Tasks

- [ ] Add demo auth middleware.
- [ ] Add role concept.
- [ ] Add case ownership check.
- [ ] Add reviewer-only routes.
- [ ] Add admin-only metrics if needed.

## Refactor Tasks

- [ ] Extract authorization policy.
- [ ] Add table-driven authorization tests.
- [ ] Keep demo auth clearly marked as non-production.

## Acceptance Criteria

- [ ] Reviewer routes are protected.
- [ ] Learners cannot see other cases.
- [ ] Demo auth limitations are documented.

---

# Phase 6 - Accessibility TDD

## Accessibility Test Checklist

Some accessibility must be checked with automated tools and some manually.

## Automated/Scriptable Tests

- [ ] `TestUI_FormFields_HaveLabels`
- [ ] `TestUI_ErrorMessages_LinkToFields`
- [ ] `TestUI_StatusTimeline_HasAccessibleStructure`
- [ ] `TestUI_Buttons_HaveAccessibleNames`
- [ ] `TestUI_NoColorOnlyStatusIndicators`

## Manual Tests

- [ ] Keyboard-only learner form submission.
- [ ] Keyboard-only reviewer decision.
- [ ] Screen-reader friendly timeline.
- [ ] 200% zoom layout.
- [ ] Mobile layout.
- [ ] Visible focus states.
- [ ] Plain-language messages.

## Green Tasks

- [ ] Add semantic HTML.
- [ ] Add labels.
- [ ] Add field-level error messages.
- [ ] Add accessible timeline markup.
- [ ] Add focus states.
- [ ] Add non-color status text.
- [ ] Add skip link.

## Acceptance Criteria

- [ ] Accessibility checklist is part of PR template.
- [ ] Critical UI accessibility tests pass.
- [ ] Manual accessibility findings are documented.

---

# Phase 7 - Privacy/Security CI Gates

## Required CI Checks

- [ ] Unit tests.
- [ ] Integration tests.
- [ ] Guardrail tests.
- [ ] Redaction tests.
- [ ] Upload security tests.
- [ ] Authorization tests.
- [ ] Accessibility smoke tests.
- [ ] Secret scan.
- [ ] Dependency vulnerability scan.

## Definition of Done

- [ ] No feature merges without guardrail tests passing.
- [ ] No new endpoint without access test.
- [ ] No new log field without redaction review.
- [ ] No AI provider call without redaction test.
