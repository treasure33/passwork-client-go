package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/treasure33/passwork-client-go/internal/utils"
)

// GetPassword Get a password by ID
func (c *Client) GetPassword(pwId string) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/items/%s", c.BaseURL, pwId)
	method := http.MethodGet
	var responseObject PasswordResponse
	var err error

	// HTTP request
	response, _, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[PasswordResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

// SearchPassword Search for password by name
func (c *Client) SearchPassword(request PasswordSearchRequest) (PasswordSearchResponse, error) {
	baseURL := fmt.Sprintf("%s/items/search", c.BaseURL)
	method := http.MethodGet
	var responseObject PasswordSearchResponse
	var err error

	// Build query parameters
	params := make([]string, 0)
	if request.Query != "" {
		params = append(params, fmt.Sprintf("query=%s", request.Query))
	}
	if request.VaultId != "" {
		params = append(params, fmt.Sprintf("vaultId=%s", request.VaultId))
	}
	if len(request.Colors) > 0 {
		colorStrs := make([]string, len(request.Colors))
		for i, color := range request.Colors {
			colorStrs[i] = fmt.Sprintf("%d", color)
		}
		params = append(params, fmt.Sprintf("colors=%s", strings.Join(colorStrs, ",")))
	}
	if len(request.Tags) > 0 {
		params = append(params, fmt.Sprintf("tags=%s", strings.Join(request.Tags, ",")))
	}
	if request.IncludeShared {
		params = append(params, "includeShared=true")
	}

	// Add query parameters to URL if any
	url := baseURL
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", baseURL, strings.Join(params, "&"))
	}

	// HTTP request
	response, _, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[PasswordSearchResponse](response)
	if err != nil {
		return responseObject, err
	}

	// Check status only if it's present (API v4 format)
	// API v1 doesn't return Status field
	if responseObject.Status != "" && responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) AddPassword(pwRequest PasswordRequest) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/items", c.BaseURL)
	method := http.MethodPost
	var responseObject PasswordResponse
	var err error

	body, err := json.Marshal(pwRequest)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	response, _, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[PasswordResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) EditPassword(pwId string, request PasswordRequest) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/items/%s", c.BaseURL, pwId)
	method := http.MethodPut
	var responseObject PasswordResponse

	body, err := json.Marshal(request)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	response, _, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[PasswordResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) DeletePassword(pwId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/items/%s", c.BaseURL, pwId)
	method := http.MethodDelete
	var responseObject DeleteResponse

	// HTTP request
	response, _, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[DeleteResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}
