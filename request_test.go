package proxycrawl_test

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/vertoforce/go-proxycrawl"
)

func TestRequest(t *testing.T) {
	c := &proxycrawl.Client{
		NormalRequestToken:     os.Getenv("RequestToken"),
		JavascriptRequestToken: os.Getenv("JavascriptToken"),
	}

	params := &proxycrawl.RequestParameters{
		URL: "https://google.com",
	}
	resp, err := c.MakeRequest(context.Background(), params, proxycrawl.JavascriptRequest)
	if err != nil {
		t.Error(err)
		return
	}

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.Contains(string(html), "About") {
		t.Errorf("did not get valid html")
	}
}
