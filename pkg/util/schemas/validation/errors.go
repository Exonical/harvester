// Package validation re-exports wrangler validation error codes so that
// consuming packages no longer need a direct wrangler import.
package validation

import wranglervalidation "github.com/rancher/wrangler/v3/pkg/schemas/validation"

type ErrorCode = wranglervalidation.ErrorCode

var (
	Unauthorized     = wranglervalidation.Unauthorized
	PermissionDenied = wranglervalidation.PermissionDenied
	NotFound         = wranglervalidation.NotFound
	MethodNotAllowed = wranglervalidation.MethodNotAllowed
	Conflict         = wranglervalidation.Conflict

	InvalidDateFormat  = wranglervalidation.InvalidDateFormat
	InvalidFormat      = wranglervalidation.InvalidFormat
	InvalidReference   = wranglervalidation.InvalidReference
	NotNullable        = wranglervalidation.NotNullable
	NotUnique          = wranglervalidation.NotUnique
	MinLimitExceeded   = wranglervalidation.MinLimitExceeded
	MaxLimitExceeded   = wranglervalidation.MaxLimitExceeded
	MinLengthExceeded  = wranglervalidation.MinLengthExceeded
	MaxLengthExceeded  = wranglervalidation.MaxLengthExceeded
	InvalidOption      = wranglervalidation.InvalidOption
	InvalidCharacters  = wranglervalidation.InvalidCharacters
	MissingRequired    = wranglervalidation.MissingRequired
	InvalidCSRFToken   = wranglervalidation.InvalidCSRFToken
	InvalidAction      = wranglervalidation.InvalidAction
	InvalidBodyContent = wranglervalidation.InvalidBodyContent
	InvalidType        = wranglervalidation.InvalidType
	ActionNotAvailable = wranglervalidation.ActionNotAvailable
	InvalidState       = wranglervalidation.InvalidState

	ServerError        = wranglervalidation.ServerError
	ClusterUnavailable = wranglervalidation.ClusterUnavailable

	ErrComplete = wranglervalidation.ErrComplete
)
