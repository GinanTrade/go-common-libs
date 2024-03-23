package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Path: lib/http.go

func PostData(endpoint string, jsonData interface{}, headers map[string]string) ([]byte, error) {
	// Create a new HTTP POST request with the JSON payload
	jsondata, err := json.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsondata))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return respBody, nil
}
func GetData(endpoint string, queryString map[string]string, headers map[string]string) ([]byte, error) {
	// Create a new GET POST request with the query string
	endpoint = endpoint + "?"

	for key, value := range queryString {
		value = url.QueryEscape(value)
		endpoint += key + "=" + value
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return respBody, nil
}
