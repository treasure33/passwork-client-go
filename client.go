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
func (c *Client) Login() error {
	url := fmt.Sprintf("%s/auth/login/%s", c.BaseURL, c.apiKey)

	response, _, err := c.sendRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	responseObject, err := utils.ParseJSONResponse[LoginResponse](response)
	if err != nil {
		return err
	}

	if responseObject.Status != "success" {
		return fmt.Errorf("login failed, status: %s", responseObject.Status)
	}

	c.sessionToken = responseObject.Data.Token

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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Passwork-Auth", c.sessionToken)

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
