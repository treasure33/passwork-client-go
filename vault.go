package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/treasure33/passwork-client-go/internal/utils"
)

func (c *Client) GetVault(vaultId string) (VaultResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
	method := http.MethodGet
	var responseObject VaultResponse
	var err error

	response, _, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	responseObject, err = utils.ParseJSONResponse[VaultResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) AddVault(vaultRequest VaultAddRequest) (VaultOperationResponse, error) {
	url := fmt.Sprintf("%s/vaults", c.BaseURL)
	method := http.MethodPost
	var responseObject VaultOperationResponse
	var err error

	body, err := json.Marshal(vaultRequest)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	response, _, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	responseObject, err = utils.ParseJSONResponse[VaultOperationResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" && responseObject.Code != "vaultCreated" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) EditVault(vaultId string, request VaultEditRequest) (VaultOperationResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
	method := http.MethodPut
	var responseObject VaultOperationResponse

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
	responseObject, err = utils.ParseJSONResponse[VaultOperationResponse](response)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Code)
	}

	return responseObject, nil
}

func (c *Client) DeleteVault(vaultId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
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
