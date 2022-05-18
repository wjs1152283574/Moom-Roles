// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/roles/service/v1/roles.proto

package v1

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

// Validate checks the field values on CreateSuperUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSuperUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSuperUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSuperUserRequestMultiError, or nil if none found.
func (m *CreateSuperUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSuperUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateSuperUserRequestMultiError(errors)
	}

	return nil
}

// CreateSuperUserRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSuperUserRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSuperUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSuperUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSuperUserRequestMultiError) AllErrors() []error { return m }

// CreateSuperUserRequestValidationError is the validation error returned by
// CreateSuperUserRequest.Validate if the designated constraints aren't met.
type CreateSuperUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSuperUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSuperUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSuperUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSuperUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSuperUserRequestValidationError) ErrorName() string {
	return "CreateSuperUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSuperUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSuperUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSuperUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSuperUserRequestValidationError{}

// Validate checks the field values on CreateSuperUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSuperUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSuperUserResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSuperUserResponseMultiError, or nil if none found.
func (m *CreateSuperUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSuperUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateSuperUserResponseMultiError(errors)
	}

	return nil
}

// CreateSuperUserResponseMultiError is an error wrapping multiple validation
// errors returned by CreateSuperUserResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateSuperUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSuperUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSuperUserResponseMultiError) AllErrors() []error { return m }

// CreateSuperUserResponseValidationError is the validation error returned by
// CreateSuperUserResponse.Validate if the designated constraints aren't met.
type CreateSuperUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSuperUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSuperUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSuperUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSuperUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSuperUserResponseValidationError) ErrorName() string {
	return "CreateSuperUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSuperUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSuperUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSuperUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSuperUserResponseValidationError{}
