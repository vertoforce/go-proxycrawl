package proxycrawl_test

import (
	"context"
	"fmt"
	"os"

	"github.com/vertoforce/go-proxycrawl"
)

func ExampleClient_MakeRequest() {
	c := &proxycrawl.Client{
		NormalRequestToken:     os.Getenv("RequestToken"),
		JavascriptRequestToken: os.Getenv("JavascriptToken"),
	}

	params := &proxycrawl.RequestParameters{
		URL: "https://google.com",
	}
	// This does NOT check for errors
	resp, _ := c.MakeRequest(context.Background(), params, proxycrawl.JavascriptRequest)
	fmt.Println(resp.StatusCode)

	// Output: 200
}
