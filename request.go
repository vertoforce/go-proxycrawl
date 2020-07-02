package proxycrawl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// Format Indicates the response format, either json or html. Defaults to html.
type Format string

// Device - Optionally, if you don't want to specify a user_agent but you want to have the requests from a specific device, you can use this parameter.
// There are two options available: desktop and mobile.
type Device string

// Country - A two letter code representing a country
type Country string

// RequestType to Send, can be javascript or normal
type RequestType int

// Constants
const (
	baseURL = "https://api.proxycrawl.com"

	HTMLFormat Format = "html"
	JSONFormat Format = "json"

	MobileDevice  Device = "mobile"
	DesktopDevice Device = "desktop"

	CountryAustria       = "AT"
	CountryAustralia     = "AU"
	CountryBelarus       = "BY"
	CountryCanada        = "CA"
	CountryChina         = "CN"
	CountryCzech         = "CZ"
	CountryGermany       = "DE"
	CountryEstonia       = "EE"
	CountrySpain         = "ES"
	CountryFrance        = "FR"
	CountryUnitedKingdom = "GB"
	CountryHong          = "HK"
	CountryIsrael        = "IL"
	CountryIndia         = "IN"
	CountryJapan         = "JP"
	CountryKuwait        = "KW"
	CountryLatvia        = "LV"
	CountryMoldova       = "MD"
	CountryNetherlands   = "NL"
	CountryPoland        = "PL"
	CountryRomania       = "RO"
	CountryRussia        = "RU"
	CountryTurkey        = "TR"
	CountryUkraine       = "UA"
	CountryUnitedStates  = "US"

	NormalRequest RequestType = iota
	JavascriptRequest
)

// RequestParameters used to make a request
type RequestParameters struct {
	// You will need a url to crawl. Make sure it starts with http or https.
	URL string `label:"url"`
	// Indicates the response format, either json or html. Defaults to html.
	// If format html is used, ProxyCrawl will send you back the response parameters in the headers.
	Format Format `label:"format"`
	// If you want to make the request with a custom user agent, you can pass it here and our servers will forward it to the requested url.
	UserAgent string `label:"user_agent"`
	// If you are using the javascript token, you can optionally pass page_wait parameter to wait an amount of milliseconds before the browser captures the resulting html code.
	// This is useful in cases where the page takes some seconds to render or some ajax needs to be loaded before the html is being captured.
	PageWait int64 `label:"page_wait"`
	// If you are using the javascript token, you can optionally pass ajax_wait parameter to wait for the ajax requests to finish before getting the html response.
	AjaxWait int64 `label:"ajax_wait"`
	// If you are using the javascript token, you can optionally pass css_click_selector parameter to click an element in the page before the browser captures the resulting html code.
	// It must be a full and valid CSS selector, for example #some-button or .some-other-button and properly encoded.
	CSSClickSelector string `label:"css_click_selector"`
	// Optionally, if you don't want to specify a user_agent but you want to have the requests from a specific device, you can use this parameter.
	// There are two options available: desktop and mobile.
	Device Device `label:"device"`
	// Optionally, if you need to get the cookies that the original website sets on the response, you can use the &get_cookies=true parameter.
	// The cookies will come back in the header (or in the json response if you use &format=json) as original_set_cookie.
	GetCookies bool `label:"get_cookies"`
	// Optionally, if you need to get the headers that the original website sets on the response, you can use the &get_headers=true parameter.
	// The headers will come back in the header (or in the json response if you use &format=json) as original_header_name.
	GetHeaders bool `label:"get_headers"`
	// If you need to use the same proxy for subsequent requests, you can use the &proxy_session= parameter.
	// The &proxy_session= parameter can be any value. Simply send a new value to create a new proxy session (this will allow you to continue using the same proxy for all subsequent requests with that proxy session value). Sessions expire 30 seconds after the last API call.
	ProxySession string `label:"proxy_session"`
	// If you need to send the cookies that come back on every request to all subsequent calls, you can use the &cookies_session= parameter.
	// The &cookies_session= parameter can be any value. Simply send a new value to create a new cookies session (this will allow you to send the returned cookies from the subsequent calls to the next API calls with that cookies session value). Sessions expire in 300 seconds after the last API call.
	CookiesSession string `label:"cookies_session"`
	// If you are using the javascript token, you can optionally pass &screenshot=true parameter to get a screenshot in JPEG format of the whole crawled page.
	// ProxyCrawl will send you back the screenshot_url in the response headers (or in the json response if you use &format=json).
	// The screenshot_url expires in one hour.
	Screenshot bool `label:"screenshot"`
	// Optionally pass &store=true parameter to store a copy of the API response in the ProxyCrawl Cloud Storage
	// ProxyCrawl will send you back the storage_url in the response headers (or in the json response if you use &format=json).
	Store bool `label:"store"`
	// Returns back the information parsed according to the specified scraper. Check the list of all the available data scrapers to see which one to choose.
	// https://proxycrawl.com/dashboard/api/scrapers
	// The response will come back as JSON.
	// Scraper is optional parameter. If you don't use it, you will receive back the full HTML of the page so you can scrape it freely.
	Sraper bool `label:"scraper"`
	// Optionally, if you need to get the scraped data of the page that you requested, you can pass &autoparse=true parameter.
	// The response will come back as JSON. The structure of the response varies depending on the URL that you sent.
	// &autoparse=true is an optional parameter. If you don't use it, you will receive back the full HTML of the page so you can scrape it freely.
	Autoparse bool `label:"autoparse"`
	// If you want your requests to be geolocated from a specific country, you can use the &country= parameter, like &country=US (two-character country code).
	// Please take into account that specifying a country can reduce the amount of successful requests you get back, so use it wisely and only when geolocation crawls are required.
	Country Country `label:"country"`
	// If you want to crawl onion websites over the Tor network, you can pass the &tor_network=true parameter.
	TorNetwork bool `label:"tor_network"`
}

// ToURL Converts the request to be ready to be placed in a url
// Will only return values that are not "", "false" or "0"
func (r *RequestParameters) ToURL() *url.Values {
	typeOf := reflect.TypeOf(*r)
	valueOf := reflect.ValueOf(*r)

	values := url.Values{}
	for i := 0; i < typeOf.NumField(); i++ {
		value := fmt.Sprintf("%v", valueOf.Field(i).Interface())
		if value == "" || value == "false" || value == "0" {
			continue
		}
		label := typeOf.Field(i).Name
		if l := typeOf.Field(i).Tag.Get("label"); l != "" {
			label = l
		}
		values.Add(label, value)
	}

	return &values
}

// MakeRequest Makes a normal or javascript request
func (c *Client) MakeRequest(ctx context.Context, params *RequestParameters, requestType RequestType) (*http.Response, error) {
	urlValues := params.ToURL()
	switch requestType {
	default:
		urlValues.Add("token", c.NormalRequestToken)
	case JavascriptRequest:
		urlValues.Add("token", c.JavascriptRequestToken)
	}

	// Wait for a token
	c.rateLimit.Wait(1)

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/?%s", baseURL, urlValues.Encode()), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	return resp, err
}
