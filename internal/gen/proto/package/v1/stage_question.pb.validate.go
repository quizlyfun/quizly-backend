// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: package/v1/stage_question.proto

package packagev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on StageQuestion with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StageQuestion) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StageQuestionMultiError, or
// nil if none found.
func (m *StageQuestion) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for StageId

	// no validation rules for TopicId

	if all {
		switch v := interface{}(m.GetQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Question",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestionValidationError{
				field:  "Question",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for QuestionType

	// no validation rules for QuestionCost

	if all {
		switch v := interface{}(m.GetAnswer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestionValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnswer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestionValidationError{
				field:  "Answer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AnswerTime

	// no validation rules for HostComment

	// no validation rules for SecretTopic

	// no validation rules for SecretCost

	// no validation rules for TransferType

	// no validation rules for IsKeepable

	if len(errors) > 0 {
		return StageQuestionMultiError(errors)
	}

	return nil
}

// StageQuestionMultiError is an error wrapping multiple validation errors
// returned by StageQuestion.ValidateAll() if the designated constraints
// aren't met.
type StageQuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestionMultiError) AllErrors() []error { return m }

// StageQuestionValidationError is the validation error returned by
// StageQuestion.Validate if the designated constraints aren't met.
type StageQuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestionValidationError) ErrorName() string { return "StageQuestionValidationError" }

// Error satisfies the builtin error interface
func (e StageQuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestionValidationError{}

// Validate checks the field values on CreateStageQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateStageQuestionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateStageQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateStageQuestionRequestMultiError, or nil if none found.
func (m *CreateStageQuestionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateStageQuestionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuestionId

	// no validation rules for TopicId

	// no validation rules for StageId

	if _, ok := _CreateStageQuestionRequest_QuestionType_InLookup[m.GetQuestionType()]; !ok {
		err := CreateStageQuestionRequestValidationError{
			field:  "QuestionType",
			reason: "value must be in list [STANDARD SAFE SECRET SUPER_SECRET AUCTION]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetQuestionCost() < 1 {
		err := CreateStageQuestionRequestValidationError{
			field:  "QuestionCost",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetAnswerTime(); val < 5 || val > 60 {
		err := CreateStageQuestionRequestValidationError{
			field:  "AnswerTime",
			reason: "value must be inside range [5, 60]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for HostComment

	// no validation rules for SecretTopic

	// no validation rules for SecretCost

	// no validation rules for IsKeepable

	if _, ok := _CreateStageQuestionRequest_TransferType_InLookup[m.GetTransferType()]; !ok {
		err := CreateStageQuestionRequestValidationError{
			field:  "TransferType",
			reason: "value must be in list [BEFORE AFTER NEVER]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateStageQuestionRequestMultiError(errors)
	}

	return nil
}

// CreateStageQuestionRequestMultiError is an error wrapping multiple
// validation errors returned by CreateStageQuestionRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateStageQuestionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateStageQuestionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateStageQuestionRequestMultiError) AllErrors() []error { return m }

// CreateStageQuestionRequestValidationError is the validation error returned
// by CreateStageQuestionRequest.Validate if the designated constraints aren't met.
type CreateStageQuestionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStageQuestionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStageQuestionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStageQuestionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStageQuestionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStageQuestionRequestValidationError) ErrorName() string {
	return "CreateStageQuestionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateStageQuestionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStageQuestionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStageQuestionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStageQuestionRequestValidationError{}

var _CreateStageQuestionRequest_QuestionType_InLookup = map[StageQuestionType]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

var _CreateStageQuestionRequest_TransferType_InLookup = map[TransferType]struct{}{
	1: {},
	2: {},
	3: {},
}

// Validate checks the field values on CreateStageQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateStageQuestionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateStageQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateStageQuestionResponseMultiError, or nil if none found.
func (m *CreateStageQuestionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateStageQuestionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStageQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateStageQuestionResponseValidationError{
					field:  "StageQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateStageQuestionResponseValidationError{
					field:  "StageQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStageQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateStageQuestionResponseValidationError{
				field:  "StageQuestion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateStageQuestionResponseMultiError(errors)
	}

	return nil
}

// CreateStageQuestionResponseMultiError is an error wrapping multiple
// validation errors returned by CreateStageQuestionResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateStageQuestionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateStageQuestionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateStageQuestionResponseMultiError) AllErrors() []error { return m }

// CreateStageQuestionResponseValidationError is the validation error returned
// by CreateStageQuestionResponse.Validate if the designated constraints
// aren't met.
type CreateStageQuestionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStageQuestionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStageQuestionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStageQuestionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStageQuestionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStageQuestionResponseValidationError) ErrorName() string {
	return "CreateStageQuestionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateStageQuestionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStageQuestionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStageQuestionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStageQuestionResponseValidationError{}

// Validate checks the field values on GetStageQuestionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStageQuestionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStageQuestionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStageQuestionRequestMultiError, or nil if none found.
func (m *GetStageQuestionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStageQuestionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StageQuestionId

	if len(errors) > 0 {
		return GetStageQuestionRequestMultiError(errors)
	}

	return nil
}

// GetStageQuestionRequestMultiError is an error wrapping multiple validation
// errors returned by GetStageQuestionRequest.ValidateAll() if the designated
// constraints aren't met.
type GetStageQuestionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStageQuestionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStageQuestionRequestMultiError) AllErrors() []error { return m }

// GetStageQuestionRequestValidationError is the validation error returned by
// GetStageQuestionRequest.Validate if the designated constraints aren't met.
type GetStageQuestionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStageQuestionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStageQuestionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStageQuestionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStageQuestionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStageQuestionRequestValidationError) ErrorName() string {
	return "GetStageQuestionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetStageQuestionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStageQuestionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStageQuestionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStageQuestionRequestValidationError{}

// Validate checks the field values on GetStageQuestionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetStageQuestionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStageQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStageQuestionResponseMultiError, or nil if none found.
func (m *GetStageQuestionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStageQuestionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStageQuestion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetStageQuestionResponseValidationError{
					field:  "StageQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetStageQuestionResponseValidationError{
					field:  "StageQuestion",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStageQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetStageQuestionResponseValidationError{
				field:  "StageQuestion",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetStageQuestionResponseMultiError(errors)
	}

	return nil
}

// GetStageQuestionResponseMultiError is an error wrapping multiple validation
// errors returned by GetStageQuestionResponse.ValidateAll() if the designated
// constraints aren't met.
type GetStageQuestionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStageQuestionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStageQuestionResponseMultiError) AllErrors() []error { return m }

// GetStageQuestionResponseValidationError is the validation error returned by
// GetStageQuestionResponse.Validate if the designated constraints aren't met.
type GetStageQuestionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStageQuestionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStageQuestionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStageQuestionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStageQuestionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStageQuestionResponseValidationError) ErrorName() string {
	return "GetStageQuestionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStageQuestionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStageQuestionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStageQuestionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStageQuestionResponseValidationError{}

// Validate checks the field values on UpdateStageQuestionsCostRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateStageQuestionsCostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateStageQuestionsCostRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateStageQuestionsCostRequestMultiError, or nil if none found.
func (m *UpdateStageQuestionsCostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStageQuestionsCostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StageId

	if l := len(m.GetQuestions()); l < 1 || l > 10 {
		err := UpdateStageQuestionsCostRequestValidationError{
			field:  "Questions",
			reason: "value must contain between 1 and 10 items, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetQuestions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateStageQuestionsCostRequestValidationError{
						field:  fmt.Sprintf("Questions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateStageQuestionsCostRequestValidationError{
						field:  fmt.Sprintf("Questions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateStageQuestionsCostRequestValidationError{
					field:  fmt.Sprintf("Questions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpdateStageQuestionsCostRequestMultiError(errors)
	}

	return nil
}

// UpdateStageQuestionsCostRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateStageQuestionsCostRequest.ValidateAll()
// if the designated constraints aren't met.
type UpdateStageQuestionsCostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStageQuestionsCostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStageQuestionsCostRequestMultiError) AllErrors() []error { return m }

// UpdateStageQuestionsCostRequestValidationError is the validation error
// returned by UpdateStageQuestionsCostRequest.Validate if the designated
// constraints aren't met.
type UpdateStageQuestionsCostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStageQuestionsCostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStageQuestionsCostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStageQuestionsCostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStageQuestionsCostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStageQuestionsCostRequestValidationError) ErrorName() string {
	return "UpdateStageQuestionsCostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStageQuestionsCostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStageQuestionsCostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStageQuestionsCostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStageQuestionsCostRequestValidationError{}

// Validate checks the field values on StageQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StageQuestion_Question) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion_Question with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StageQuestion_QuestionMultiError, or nil if none found.
func (m *StageQuestion_Question) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion_Question) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestion_QuestionValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestion_QuestionValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestion_QuestionValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StageQuestion_QuestionMultiError(errors)
	}

	return nil
}

// StageQuestion_QuestionMultiError is an error wrapping multiple validation
// errors returned by StageQuestion_Question.ValidateAll() if the designated
// constraints aren't met.
type StageQuestion_QuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestion_QuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestion_QuestionMultiError) AllErrors() []error { return m }

// StageQuestion_QuestionValidationError is the validation error returned by
// StageQuestion_Question.Validate if the designated constraints aren't met.
type StageQuestion_QuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestion_QuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestion_QuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestion_QuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestion_QuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestion_QuestionValidationError) ErrorName() string {
	return "StageQuestion_QuestionValidationError"
}

// Error satisfies the builtin error interface
func (e StageQuestion_QuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion_Question.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestion_QuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestion_QuestionValidationError{}

// Validate checks the field values on StageQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StageQuestion_Answer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StageQuestion_Answer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StageQuestion_AnswerMultiError, or nil if none found.
func (m *StageQuestion_Answer) ValidateAll() error {
	return m.validate(true)
}

func (m *StageQuestion_Answer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Text

	// no validation rules for Author

	// no validation rules for MediaUrl

	if all {
		switch v := interface{}(m.GetCreationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StageQuestion_AnswerValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StageQuestion_AnswerValidationError{
					field:  "CreationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StageQuestion_AnswerValidationError{
				field:  "CreationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StageQuestion_AnswerMultiError(errors)
	}

	return nil
}

// StageQuestion_AnswerMultiError is an error wrapping multiple validation
// errors returned by StageQuestion_Answer.ValidateAll() if the designated
// constraints aren't met.
type StageQuestion_AnswerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StageQuestion_AnswerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StageQuestion_AnswerMultiError) AllErrors() []error { return m }

// StageQuestion_AnswerValidationError is the validation error returned by
// StageQuestion_Answer.Validate if the designated constraints aren't met.
type StageQuestion_AnswerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StageQuestion_AnswerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StageQuestion_AnswerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StageQuestion_AnswerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StageQuestion_AnswerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StageQuestion_AnswerValidationError) ErrorName() string {
	return "StageQuestion_AnswerValidationError"
}

// Error satisfies the builtin error interface
func (e StageQuestion_AnswerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStageQuestion_Answer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StageQuestion_AnswerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StageQuestion_AnswerValidationError{}

// Validate checks the field values on UpdateStageQuestionsCostRequest_Question
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateStageQuestionsCostRequest_Question) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UpdateStageQuestionsCostRequest_Question with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// UpdateStageQuestionsCostRequest_QuestionMultiError, or nil if none found.
func (m *UpdateStageQuestionsCostRequest_Question) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateStageQuestionsCostRequest_Question) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuestionId

	// no validation rules for Cost

	if len(errors) > 0 {
		return UpdateStageQuestionsCostRequest_QuestionMultiError(errors)
	}

	return nil
}

// UpdateStageQuestionsCostRequest_QuestionMultiError is an error wrapping
// multiple validation errors returned by
// UpdateStageQuestionsCostRequest_Question.ValidateAll() if the designated
// constraints aren't met.
type UpdateStageQuestionsCostRequest_QuestionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateStageQuestionsCostRequest_QuestionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateStageQuestionsCostRequest_QuestionMultiError) AllErrors() []error { return m }

// UpdateStageQuestionsCostRequest_QuestionValidationError is the validation
// error returned by UpdateStageQuestionsCostRequest_Question.Validate if the
// designated constraints aren't met.
type UpdateStageQuestionsCostRequest_QuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) ErrorName() string {
	return "UpdateStageQuestionsCostRequest_QuestionValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStageQuestionsCostRequest_QuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStageQuestionsCostRequest_Question.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStageQuestionsCostRequest_QuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStageQuestionsCostRequest_QuestionValidationError{}
