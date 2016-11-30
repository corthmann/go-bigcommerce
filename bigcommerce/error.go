package bigcommerce

import (
	"fmt"
)

type APIError struct {
	Errors []ErrorDetail
}

type ErrorDetail struct {
	Code    int32  `json:"status"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		return fmt.Sprintf("bigcommerce: %d %v", err.Code, err.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if len(e.Errors) == 0 {
		return true
	}
	return false
}

// relevantError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}
