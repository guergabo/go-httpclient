package gohttp

import (
	"errors"
	"net/http"
)

// private struct method
// body can be of any type
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Unable to create an http request")
	}

	// don't want to copy struct, prob large
	return response, err
}
