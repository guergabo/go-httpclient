package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/guergabo/go-httpclient/gohttp"
)

// singleton
/* A singleton design pattern restricts the instantiation of a class to a single
instance. This is done in order to privde coordinated access to a certian resource,
throughout an entire software system. Singletons are sometimes considered to be an
alternative to global variables or static classes. Same client for different request,
in this case. Won't change its state. Create a basic header with httpclient so it
can be a more useful singleton. Changing only the unique parts of each request. Always
send default headers. */
var (
	client = newDefaultClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

}

// default configuration for whatever you are working with, default and specific options above
func newDefaultClient() gohttp.HttpClient {

	client := gohttp.NewClient()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func createUser(user User) {

	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bs, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bs))
}
