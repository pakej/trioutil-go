package TrioUtil

import (
	"net/http"
	"net/url"
	"strings"
)

// returns the endpoint after parsing the phone and message parameters
func getEndpoint(phone string, message string) string {
	var Url *url.URL
	Url, err := url.Parse("http://cloudsms.trio-mobile.com/index.php/api/bulk_mt?")

	if err != nil {
		return ""
	}

	parameters := url.Values{}
	parameters.Add("api_key", "my-api-key")
	parameters.Add("action", "send")
	parameters.Add("to", phone)
	parameters.Add("msg", message)
	parameters.Add("sender_id", "CLOUDSMS")
	parameters.Add("content_type", "1")
	parameters.Add("mode", "shortcode")
	Url.RawQuery = parameters.Encode()

	// replaces all occurences of '+' to '%20' since the encoder
	// encodes spaces to '+' instead of '%20'
	str := strings.Replace(Url.String(), "+", "%20", -1)
	return str
}

// SendSms - send SMS to TRIO service provider
func SendSms(phone string, message string) (*http.Response, error) {
	endpoint := getEndpoint(phone, message)
	resp, err := http.Post(endpoint, "", nil)

	return resp, err
}
