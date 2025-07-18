package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Doer represents minimal interface requiring basic http.Client functionality
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Client implements a client for Server.
type Client struct {
	// It is good practice to add an underlying http.Client
	// This allows us to:
	// 		1. Set connection parameters (timeout, cookies, redirect handling)
	//		2. Specify Transport (low level connection parameters - buffer sizes, protocols, tls)
	//		3. Substitute with a different client for testing
	client Doer
	URI    string
}

func NewClient(URI string, client Doer) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		client: client,
		URI:    URI,
	}
}

func (c Client) GetSize(data []byte) (int, error) {
	request, err := http.NewRequest(http.MethodGet, c.URI, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	respData, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	out, err := strconv.Atoi(string(respData))
	if err != nil {
		return 0, ErrInvalidResponse

		// Wrapping the general error with our domain error for handling down the line
		//return 0, errors.Join(ErrInvalidResponse, err)
	}
	return out, nil
}
