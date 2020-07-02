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
	// ProxyCrawl has a rate limit of 20 requests per second, we set it conservatively at 18
	rateLimit := ratelimit.NewBucketWithQuantum(time.Second, 18, 18)
	return &Client{
		NormalRequestToken:     NormalRequestToken,
		JavascriptRequestToken: JavascriptRequestToken,
		rateLimit:              rateLimit,
	}

}
