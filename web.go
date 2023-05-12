package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Not really best practice but simplifying by creating one client available for reuse
var client = newHTTPClient()

func newHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
		},
		Timeout: 5 * time.Second,
	}
	return client
}

func addQueryParams(req *http.Request, cmd string) *http.Request {
	q := req.URL.Query()
	q.Add("cmd", cmd)
	req.URL.RawQuery = q.Encode()
	return req
}

func buildRequest(ctx context.Context, urlBase string, cmd string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlBase, nil)
	if err != nil {
		return nil, err
	}
	req = addQueryParams(req, cmd)
	return req, nil
}

func makeRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetPage(ctx context.Context, urlBase string, cmd string) ([]byte, string, error) {
	req, err := buildRequest(ctx, urlBase, cmd)
	if err != nil {
		return nil, "", err
	}
	b, err := makeRequest(client, req)
	return b, req.URL.String(), err
}
