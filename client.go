package billingo

import (
	"net/http"
)

const ServerURL = "https://api.billingo.hu/v3"

type apiKeyInjector struct {
	client *http.Client
	token  string
}

func newApiKeyInjector(client *http.Client, token string) *apiKeyInjector {
	return &apiKeyInjector{
		client: client,
		token:  token,
	}
}

func (h *apiKeyInjector) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-API-KEY", h.token)
	return h.client.Do(req)
}

func ToPtr[T any](v T) *T {
	return &v
}

// New creates a new Client with the given token.
func New(token string) (*ClientWithResponses, error) {
	return NewClientWithResponses(ServerURL, WithHTTPClient(newApiKeyInjector(http.DefaultClient, token)))
}
