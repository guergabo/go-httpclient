package main

import (
	"fmt"
	"io/ioutil"

	"github.com/guergabo/go-httpclient/gohttp"
)

func main() {

	client := gohttp.NewClient()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bs, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bs))
}
