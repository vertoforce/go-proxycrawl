# Go-ProxyCrawl

[![Go Report Card](https://goreportcard.com/badge/github.com/vertoforce/go-proxycrawl)](https://goreportcard.com/report/github.com/vertoforce/go-proxycrawl)
[![Documentation](https://godoc.org/github.com/vertoforce/go-proxycrawl?status.svg)](https://godoc.org/github.com/vertoforce/go-proxycrawl)

A simple library to use the proxy crawl [crawling api](https://proxycrawl.com/dashboard/api/docs)

## Usage

See the godoc for more examples on usage

```go
c := &proxycrawl.Client{
    NormalRequestToken:     os.Getenv("RequestToken"),
    JavascriptRequestToken: os.Getenv("JavascriptToken"),
}

params := &proxycrawl.RequestParameters{
    URL: "https://google.com",
}

resp, _ := c.MakeRequest(context.Background(), params, proxycrawl.JavascriptRequest)
```
