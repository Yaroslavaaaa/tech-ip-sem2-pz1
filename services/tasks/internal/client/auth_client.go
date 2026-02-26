package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"tech-ip-sem2/shared/httpx"
)

type AuthClient struct {
	httpClient *httpx.Client
}

type VerifyResponse struct {
	Valid   bool   `json:"valid"`
	Subject string `json:"subject"`
	Error   string `json:"error"`
}

func NewAuthClient(baseURL string, timeout time.Duration) *AuthClient {
	return &AuthClient{
		httpClient: httpx.NewClient(baseURL, timeout),
	}
}

func (c *AuthClient) VerifyToken(ctx context.Context, token string) (string, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.httpClient.Get(ctx, "/v1/auth/verify", headers)
	if err != nil {
		return "", fmt.Errorf("auth service request failed: %w", err)
	}
	defer resp.Body.Close()

	var verifyResp VerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&verifyResp); err != nil {
		return "", fmt.Errorf("failed to decode auth response: %w", err)
	}

	if resp.StatusCode != http.StatusOK || !verifyResp.Valid {
		return "", fmt.Errorf("token invalid: %s", verifyResp.Error)
	}

	return verifyResp.Subject, nil
}
