// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bootstrap/v1/bootstrap.proto

package bootstrapv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _bootstrap_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Bootstrap) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetServer() == nil {
		return BootstrapValidationError{
			field:  "Server",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetServer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Server",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetOriginServer() == nil {
		return BootstrapValidationError{
			field:  "OriginServer",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOriginServer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "OriginServer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetLogging() == nil {
		return BootstrapValidationError{
			field:  "Logging",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLogging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Logging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetCache() == nil {
		return BootstrapValidationError{
			field:  "Cache",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetCache()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Cache",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMetricsSink() == nil {
		return BootstrapValidationError{
			field:  "MetricsSink",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMetricsSink()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "MetricsSink",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetAdmin() == nil {
		return BootstrapValidationError{
			field:  "Admin",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAdmin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Admin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Server) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAddress() == nil {
		return ServerValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerValidationError is the validation error returned by Server.Validate if
// the designated constraints aren't met.
type ServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerValidationError) ErrorName() string { return "ServerValidationError" }

// Error satisfies the builtin error interface
func (e ServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerValidationError{}

// Validate checks the field values on Upstream with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Upstream) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAddress() == nil {
		return UpstreamValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Ads

	return nil
}

// UpstreamValidationError is the validation error returned by
// Upstream.Validate if the designated constraints aren't met.
type UpstreamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamValidationError) ErrorName() string { return "UpstreamValidationError" }

// Error satisfies the builtin error interface
func (e UpstreamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstream.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamValidationError{}

// Validate checks the field values on Logging with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Logging) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Path

	if _, ok := Logging_Level_name[int32(m.GetLevel())]; !ok {
		return LoggingValidationError{
			field:  "Level",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// LoggingValidationError is the validation error returned by Logging.Validate
// if the designated constraints aren't met.
type LoggingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoggingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoggingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoggingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoggingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoggingValidationError) ErrorName() string { return "LoggingValidationError" }

// Error satisfies the builtin error interface
func (e LoggingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogging.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoggingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoggingValidationError{}

// Validate checks the field values on Cache with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Cache) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTtl() == nil {
		return CacheValidationError{
			field:  "Ttl",
			reason: "value is required",
		}
	}

	if d := m.GetTtl(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return CacheValidationError{
				field:  "Ttl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return CacheValidationError{
				field:  "Ttl",
				reason: "value must be greater than or equal to 0s",
			}
		}

	}

	// no validation rules for MaxEntries

	return nil
}

// CacheValidationError is the validation error returned by Cache.Validate if
// the designated constraints aren't met.
type CacheValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CacheValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CacheValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CacheValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CacheValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CacheValidationError) ErrorName() string { return "CacheValidationError" }

// Error satisfies the builtin error interface
func (e CacheValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCache.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CacheValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CacheValidationError{}

// Validate checks the field values on SocketAddress with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SocketAddress) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateHostname(m.GetAddress()); err != nil {
		if ip := net.ParseIP(m.GetAddress()); ip == nil {
			return SocketAddressValidationError{
				field:  "Address",
				reason: "value must be a valid hostname, or ip address",
			}
		}
	}

	if m.GetPortValue() > 65535 {
		return SocketAddressValidationError{
			field:  "PortValue",
			reason: "value must be less than or equal to 65535",
		}
	}

	return nil
}

func (m *SocketAddress) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// SocketAddressValidationError is the validation error returned by
// SocketAddress.Validate if the designated constraints aren't met.
type SocketAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketAddressValidationError) ErrorName() string { return "SocketAddressValidationError" }

// Error satisfies the builtin error interface
func (e SocketAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketAddressValidationError{}

// Validate checks the field values on Admin with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Admin) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAddress() == nil {
		return AdminValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdminValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AdminValidationError is the validation error returned by Admin.Validate if
// the designated constraints aren't met.
type AdminValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminValidationError) ErrorName() string { return "AdminValidationError" }

// Error satisfies the builtin error interface
func (e AdminValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminValidationError{}

// Validate checks the field values on MetricsSink with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MetricsSink) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Type.(type) {

	case *MetricsSink_Statsd:

		if v, ok := interface{}(m.GetStatsd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetricsSinkValidationError{
					field:  "Statsd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return MetricsSinkValidationError{
			field:  "Type",
			reason: "value is required",
		}

	}

	return nil
}

// MetricsSinkValidationError is the validation error returned by
// MetricsSink.Validate if the designated constraints aren't met.
type MetricsSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsSinkValidationError) ErrorName() string { return "MetricsSinkValidationError" }

// Error satisfies the builtin error interface
func (e MetricsSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsSinkValidationError{}

// Validate checks the field values on Statsd with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Statsd) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAddress() == nil {
		return StatsdValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsdValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRootPrefix()) < 1 {
		return StatsdValidationError{
			field:  "RootPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetFlushInterval() == nil {
		return StatsdValidationError{
			field:  "FlushInterval",
			reason: "value is required",
		}
	}

	if d := m.GetFlushInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return StatsdValidationError{
				field:  "FlushInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return StatsdValidationError{
				field:  "FlushInterval",
				reason: "value must be greater than or equal to 0s",
			}
		}

	}

	return nil
}

// StatsdValidationError is the validation error returned by Statsd.Validate if
// the designated constraints aren't met.
type StatsdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsdValidationError) ErrorName() string { return "StatsdValidationError" }

// Error satisfies the builtin error interface
func (e StatsdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatsd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsdValidationError{}
