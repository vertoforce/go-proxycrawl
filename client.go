package proxycrawl

import (
	"time"

	"github.com/juju/ratelimit"
)

// Client for proxycrawl
type Client struct {
	NormalRequestToken     string
	JavascriptRequestToken string

	rateLimit *ratelimit.Bucket
}

// New creates a new client with the rate limit set up
func New(NormalRequestToken, JavascriptRequestToken string) *Client {
	// ProxyCrawl has a rate limit of 20 requests per second, which means approx 1 request every 50 milliseconds
	// We set it to conservatively 1 request every 60 milliseconds
	rateLimit := ratelimit.NewBucketWithQuantum(time.Millisecond*60, 1, 1)
	return &Client{
		NormalRequestToken:     NormalRequestToken,
		JavascriptRequestToken: JavascriptRequestToken,
		rateLimit:              rateLimit,
	}

}
