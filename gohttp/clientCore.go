package gohttp

import (
	"errors"
	"net/http"
)

// private struct method
// body can be of any type
// pointer to http.Response also is useful because it allows us to return nil if something went wrong
// want to provide content-type management, we will do all of the transformation
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	// build http request
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("Unable to create a new request")
	}

	fullHeaders := c.getRequestHeaders(headers)
	request.Header = fullHeaders // overriding the headers, possible because request is a POINTER to http.Request so changes actual copy

	// get response
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("Unable to create an http request")
	}
	return response, err
}

// returns the finalize http.Header map after overriding any default configuration, map is a reference type
func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	// adding COMMON (default) headers to request
	for header, value := range c.Headers {
		// headers is a map[string][]string, will panic if they don't specify value
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// adding CUSTOM headers to request
	// THESE HEADERS OVERWRITE THE DEFAULT HEADERS
	for header, value := range requestHeaders {
		// headers is a map[string][]string, will panic if they don't specify value
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

// NewRequest requires an io.Readedr - by default assume JSON
// we need content-type specified in the headers to know the content-type
func (c *httpClient) getRequestBody(body interface{}) ([]byte, error) {

	if body == nil {
		return nil, nil
	}

}
