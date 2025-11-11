package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/treasure33/passwork-client-go/internal/utils"
)

// Get a password by ID
func (c *Client) GetPassword(pwId string) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
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

// Search for password by name
func (c *Client) SearchPassword(request PasswordSearchRequest) (PasswordSearchResponse, error) {
	url := fmt.Sprintf("%s/passwords/search", c.BaseURL)
	method := http.MethodPost
	var responseObject PasswordSearchResponse
	var err error

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
	responseObject, err = utils.ParseJSONResponse[PasswordSearchResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) AddPassword(pwRequest PasswordRequest) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/passwords", c.BaseURL)
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
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
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
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
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
