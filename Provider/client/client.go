package air

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default AIR URL
const HostURL string = "http://localhost:5000"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Resource     string `json:"resource"`
}

type AuthResponse struct {
	Token string `json:"access_token"`
}

// NewClient -
func NewClient(host, clientId, clientSecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default AIR URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If clientId or clientSecret not provided, return empty client
	if clientId == nil || clientSecret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		GrantType:    "client_credentials",
		Resource:     "api://air-dev.gsk.onmicrosoft.com",
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
