// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/capture.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Connection with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Connection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetLocalAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				field:  "LocalAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				field:  "RemoteAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ConnectionValidationError is the validation error returned by
// Connection.Validate if the designated constraints aren't met.
type ConnectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionValidationError) ErrorName() string { return "ConnectionValidationError" }

// Error satisfies the builtin error interface
func (e ConnectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.EventSelector.(type) {

	case *Event_Read_:

		if v, ok := interface{}(m.GetRead()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "Read",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Event_Write_:

		if v, ok := interface{}(m.GetWrite()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "Write",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on Trace with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trace) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TraceValidationError{
				field:  "Connection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TraceValidationError is the validation error returned by Trace.Validate if
// the designated constraints aren't met.
type TraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceValidationError) ErrorName() string { return "TraceValidationError" }

// Error satisfies the builtin error interface
func (e TraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceValidationError{}

// Validate checks the field values on Event_Read with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event_Read) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Data

	return nil
}

// Event_ReadValidationError is the validation error returned by
// Event_Read.Validate if the designated constraints aren't met.
type Event_ReadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Event_ReadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Event_ReadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Event_ReadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Event_ReadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Event_ReadValidationError) ErrorName() string { return "Event_ReadValidationError" }

// Error satisfies the builtin error interface
func (e Event_ReadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent_Read.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Event_ReadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Event_ReadValidationError{}

// Validate checks the field values on Event_Write with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Event_Write) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Data

	// no validation rules for EndStream

	return nil
}

// Event_WriteValidationError is the validation error returned by
// Event_Write.Validate if the designated constraints aren't met.
type Event_WriteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Event_WriteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Event_WriteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Event_WriteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Event_WriteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Event_WriteValidationError) ErrorName() string { return "Event_WriteValidationError" }

// Error satisfies the builtin error interface
func (e Event_WriteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent_Write.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Event_WriteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Event_WriteValidationError{}
