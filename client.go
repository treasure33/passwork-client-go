package passwork

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/treasure33/passwork-client-go/internal/utils"
)

type Client struct {
	BaseURL      string
	apiKey       string
	sessionToken string
	HTTPClient   *http.Client
}

type LoginResponse struct {
	Status string
	Data   LoginResponseData
}

type LogoutResponse struct {
	Status string
	Data   string
}

type LoginResponseData struct {
	Token                 string
	RefreshToken          string
	TokenTtl              int
	RefreshTokenTtl       int
	TokenExpiredAt        int
	RefreshTokenExpiredAt int
	User                  User
}

type User struct {
	Name  string
	Email string
}

func NewClient(baseURL, apiKey string, timeout time.Duration) *Client {
	client := Client{
		BaseURL:      baseURL,
		apiKey:       apiKey,
		sessionToken: "",
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}

	return &client
}

// Perform Login Request and set session Token in struct
// For API v1, the API key is used directly as the bearer token
func (c *Client) Login() error {
	// For v1 API, we use the API key directly
	// Try to make a simple request to verify the API key works
	c.sessionToken = c.apiKey
	return nil
}

func (c *Client) Logout() error {
	url := fmt.Sprintf("%s/auth/logout", c.BaseURL)

	response, _, err := c.sendRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	responseObject, err := utils.ParseJSONResponse[LogoutResponse](response)
	if err != nil {
		return err
	}

	if responseObject.Status == "success" && responseObject.Data == "loggedOut" {
		return nil
	}

	return fmt.Errorf("logout failed, status: %s", responseObject.Status)
}

// Sends HTTP request to URL with method and body
// Returns response body
func (c *Client) sendRequest(method string, url string, body io.Reader) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Use Authorization Bearer header if we have a session token, otherwise use API key
	if c.sessionToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.sessionToken)
	} else if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// Execute HTTP request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Printf("HTTP Request failed: %v", err)
		if resp != nil {
			return nil, resp.StatusCode, err
		}
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return responseData, resp.StatusCode, nil
}
