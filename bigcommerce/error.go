package bigcommerce

import (
	"fmt"
)

type APIError []struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details struct {
		Errors []struct {
			Type    string `json:"type"`
			Product struct {
				ID             int    `json:"id"`
				Name           string `json:"name"`
				InventoryLevel int    `json:"inventory_level"`
				URL            string `json:"url"`
				Resource       string `json:"resource"`
			} `json:"product"`
		} `json:"errors"`
	} `json:"details"`
}

func (e APIError) Error() string {
	if len(e) > 0 {
		err := e[0]
		return fmt.Sprintf("bigcommerce: %d %v", err.Status, err.Message)
	}
	return ""
}

// Empty returns true if empty. Otherwise, at least 1 error message/code is
// present and false is returned.
func (e APIError) Empty() bool {
	if len(e) == 0 {
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
