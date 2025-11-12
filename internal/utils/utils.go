package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// ParseJSONResponse unmarshals JSON data into a given struct type.
// Supports both direct JSON and base64-encoded JSON (API v1 format).
func ParseJSONResponse[T any](data []byte) (T, error) {
	var responseObject T
	
	// Try to detect if this is API v1 base64 format
	var base64Response struct {
		Format  string `json:"format"`
		Content string `json:"content"`
	}
	
	if err := json.Unmarshal(data, &base64Response); err == nil && base64Response.Format == "base64" {
		// Decode base64 content
		decoded, err := base64.StdEncoding.DecodeString(base64Response.Content)
		if err != nil {
			return responseObject, fmt.Errorf("failed to decode base64 content: %w", err)
		}
		
		// Parse the decoded JSON
		err = json.Unmarshal(decoded, &responseObject)
		if err != nil {
			return responseObject, fmt.Errorf("failed to parse decoded JSON: %w", err)
		}
		return responseObject, nil
	}
	
	// Otherwise, parse as direct JSON (API v4 format)
	err := json.Unmarshal(data, &responseObject)
	if err != nil {
		return responseObject, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return responseObject, nil
}
