package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newHTTPClient() *http.Client {
	client := &http.Client{
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

func addUserAgent(req *http.Request, userAgent string) *http.Request {
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "application/json")
	return req
}

func buildRequest(ctx context.Context, urlBase string, cmd string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlBase, nil)
	if err != nil {
		return nil, err
	}
	req = addQueryParams(req, cmd)
	// req = addUserAgent(req, userAgent)
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

func GetPage(ctx context.Context, urlBase string, cmd string) ([]byte, error) {
	client := newHTTPClient()
	req, err := buildRequest(ctx, urlBase, cmd)
	if err != nil {
		return nil, err
	}
	return makeRequest(client, req)
}
