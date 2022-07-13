/*
Know how to develop your own http client
remove dependencies on third-party libraries
depend only on golang standard library
allow us to configure all aspects of http protocol
mocking capabilities
http configuration = new headers in the http request
no external libraries only the core of go
"a little copying is better than a little dependency"

HTTP Calls
----------
Request (method, URL, headers (allow us to configure request), body)
- size: 1.55 KB (most response size comes from the hearder not the body)
Response (status code, headers (allow us to configure response), body)

// if we don't have a response in 500ms we want to abort from the moment
// we send request, won't be there to handle response
// can time out to make the connection AND
// can time out to get the response from time of request
Connections ()
Timeouts ()

If we wanted to configure the headers we cannot use the default http
client because it only sends a request with out configuring the headers
at least with client.GET(), we can use http.NewRequest to configure with
request.Header.Set(). Even though it is easy, if the server gets slower
we have no way of configuring http client in order to not wait so long
and be able to configure a given time out. CPU goes to hell because
go will wait and hang. You will be leaking resources. http.Get uses a client
with out timeouts, so are dangerous to use on the internet.

scenario
30 different goroutines waiting for the response, that takes a lot of CPU
wait and try again? pause timeout on request.
if you have an application on production take a look on timeouts.
configure timeouts. issue with defaul http.Client. there are ways to
configure it but for consistency you don't want to keep copying the logic
but rather create a package to use for http requests. problem when resource
server gets slower. creating a wrapper of http.Client{}. want metrics...
http client is just used a lot and want to create custom one that gives
you everything you might need out of the box. before requests send this
metrics to datadog. etc. single point where changes can be made.

hard to get 100% coverage using default http.Client, can't test error
after because don't have control over what the http.Client returns so
you can't test the logic afterwards, the way http.Client works does
not provide mocking features

Go Modules
----------
the new standard way of working with dependencies in go since 1.14
before 1.14 modules were beta and used dep and GOPATH.

GOROOT - the Go installation folder, where go is installed. defines where
your go SDK is located. you do not need to change this variable unless you
plan to use different go versions.

GOPATH - inside is three different folders:
defines go workspaces on the computer, every dependency you download
on any project gets download on this GOPATH.
before 1.13, every go project needed to be clone inside your GOPATH
- it defines the root of your workspace

go < 1.13 = no modules available. dep was the dependency management system
go 1.13 = modules as BETA
go 1.14 < = modules are the standard for managing dependencies
go 1.8 = default GOPATH ~/go
1. src / location of go source code
	github.com
		username
			repo
2. pkg / location of compiled package code
	mod
		github.com
			stripe
				dependency right here
3. bin / location of compiled executable programs by Go

package v. program (go module, defines a program)

(github becomes the path you are using, can be outside of GOPATH)
(if we want to use a given dependency in our application, )
(want to name the mod github path because it will made it standard,
for others to download and get it)
module github.com/guergabo/go-httpclient

go mod tidy to clean up and update the go.mod

Package Organization
--------------------
working on a library therefore, does not need a main.go
main.go, main package, and func main() are required for executables only

defining the public interface (
	don't put all your code in the root package (module)
	then public interface gets messing, work with
	sub packages
)
package name can't have '-'
							[module]     /[package]
import "github.com/guergabo/go-httpclient/gohttp"

func main() {
	gohttp.Client
}

packages encapsulate behavior

Structs
-------
- Collection of fields

type Client struct {
	Attr string
	Attr2 string
}


*/
package gohttp

import "net/http"

// public interface, loosely coupled
type HttpClient interface {
	SetHeaders(http.Header)

	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)
}

// can't access or initialize interface directly
// so create function to return implementation
// can't have pointer to interface, just interface cause
// it can't be instantiated
// CAN DEFINE DEFAULT HEADERS HERE, allow default configuration establish
func NewClient() HttpClient {
	// returns a pointer to httpClient
	client := &httpClient{}
	return client
}

// for a struct to be consider http client must implement functions above
// provides all functionality for the http client
type httpClient struct {
	Headers http.Header
}

// method not a package function, gohttp.Get() versus client.Get()
// upper case means exported and can be used outside of package
// struct is NOT a reference type therefore in order to make receiver
// functions that actually effect the struct you must pass a pointer
// to it. Example: for structs that contain a sync.Mutex or similar
// synchronizing field (they musn’t be copied).
// go is pass by value unless you say otherwise. Slices and maps
// are reference types though.
// http.Header is alias for ->
// type Header map[string][]string
// need to define it as part of the interface if not will be accesisble
// but not in the logical unti
func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
