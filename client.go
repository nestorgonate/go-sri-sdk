package gosrisdk

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
	headers    map[string]string
	Taxpayer   *TaxpayerService
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: "https://srienlinea.sri.gob.ec",
		headers: make(map[string]string),
	}
	client.Taxpayer = &TaxpayerService{
		client: client,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}
