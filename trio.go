/*
Package TrioUtil is a client for TrioMobile (http://trio-mobile.com) that utilises TrioMobile's API.
*/
package TrioUtil

import (
	"net/http"
	"net/url"
	"strings"
)

// Trio represents the Trio Instance
type Trio struct {
	token  string
	params url.Values
}

// New - creates a new instance of Trio
func New(token string) *Trio {
	params := url.Values{}
	params.Set("api_key", token)
	params.Set("action", "send")
	params.Set("sender_id", "CLOUDSMS")
	params.Set("content_type", "1")
	params.Set("mode", "shortcode")

	return &Trio{
		token,
		params,
	}
}

// SendSms - send SMS to TRIO service provider
func (t *Trio) SendSms(phone string, message string) (*http.Response, error) {
	// gets all the predefined url parameters
	url_parameters := t.params
	// set SendSms func specific url parameters
	url_parameters.Set("to", phone)
	url_parameters.Set("msg", message)
	// parse the url parameters into an endpoint
	endpoint := getEndpoint(url_parameters)
	// sends the POST request to TrioMobile
	resp, err := http.Post(endpoint, "", nil)

	return resp, err
}

// returns the endpoint after parsing url parameters
func getEndpoint(url_parameters url.Values) string {
	var Url *url.URL
	Url, err := url.Parse("http://cloudsms.trio-mobile.com/index.php/api/bulk_mt?")

	if err != nil {
		return ""
	}

	parameters := url_parameters
	Url.RawQuery = parameters.Encode()

	// replaces all occurences of '+' to '%20' since the encoder
	// encodes spaces to '+' instead of '%20'
	str := strings.Replace(Url.String(), "+", "%20", -1)
	return str
}
