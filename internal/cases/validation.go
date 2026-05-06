package cases

import (
	"strings"
)

type FieldError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ValidationError struct {
	Fields []FieldError
}

func (e ValidationError) Error() string {
	return "the request contains invalid fields"
}

func ValidateCreateInput(input CreateInput) error {
	var fieldErrors []FieldError

	if strings.TrimSpace(input.FormType) == "" {
		fieldErrors = append(fieldErrors, FieldError{
			Name:    "form_type",
			Message: "form_type is required",
		})
	} else if input.FormType != FormTypeTransferCredit {
		fieldErrors = append(fieldErrors, FieldError{
			Name:    "form_type",
			Message: "form_type must be transfer_credit",
		})
	}

	if strings.TrimSpace(input.LearnerRef) == "" {
		fieldErrors = append(fieldErrors, FieldError{
			Name:    "learner_ref",
			Message: "learner_ref is required",
		})
	}

	if strings.TrimSpace(input.Term) == "" {
		fieldErrors = append(fieldErrors, FieldError{
			Name:    "term",
			Message: "term is required",
		})
	}

	requiredFields := []string{
		"prior_institution",
		"prior_course_code",
		"target_program",
	}

	for _, fieldName := range requiredFields {
		if strings.TrimSpace(input.Fields[fieldName]) == "" {
			fieldErrors = append(fieldErrors, FieldError{
				Name:    fieldName,
				Message: fieldName + " is required",
			})
		}
	}

	if len(fieldErrors) == 0 {
		return nil
	}

	return ValidationError{Fields: fieldErrors}
}
