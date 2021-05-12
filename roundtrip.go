package proxycrawl

import (
	"fmt"
	"net/http"
	"time"
)

// JavascriptClient is a http.RoundTripper that makes javascript requests through proxycrawl
type JavascriptClient struct {
	c        *Client
	PageWait time.Duration
}

// NormalClient is a http.RoundTripper that makes normal requests through proxycrawl
type NormalClient struct {
	c *Client
}

func (c *Client) NewJavascriptRoundTripper() *JavascriptClient {
	return &JavascriptClient{c: c, PageWait: time.Second * 3}
}

func (c *Client) NewNormalClientRoundTripper() *NormalClient {
	return &NormalClient{c: c}
}

func (j *JavascriptClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return j.c.RoundTrip(req, JavascriptRequest, j.PageWait)
}

func (n *NormalClient) RoundTrip(req *http.Request) (*http.Response, error) {
	return n.c.RoundTrip(req, NormalRequest, 0)
}

// RoundTrip makes a single round trip with the provided request and requestType.
// If you want a http.RoundTripper, see NewJavascriptRoundTripper and NewNormalClientRoundTripper.
// Params:
// - pageWait - How long to wait for the page to load for javascript requests
func (c *Client) RoundTrip(req *http.Request, requestType RequestType, pageWait time.Duration) (*http.Response, error) {
	resp, err := c.MakeRequest(req.Context(), &RequestParameters{
		URL:        req.URL.String(),
		UserAgent:  req.Header.Get("User-Agent"),
		PageWait:   pageWait.Milliseconds(),
		GetHeaders: true,
	}, requestType)
	if err != nil {
		return nil, fmt.Errorf("proxycrawl error: %w", err)
	}

	return resp, nil
}
