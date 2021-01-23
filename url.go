package marshaler

import (
	"encoding/json"
	"net/url"
)

// URL extends the url.URL to add implementations
// to the Marshaler and Unmarshaler interfaces
type URL struct {
	*url.URL
}

// NewURL returns a new marshaler.URL from u
func NewURL(u *url.URL) URL {
	return URL{
		URL: u,
	}
}

// UnmarshalJSON transforms the b to a string
func (u *URL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	url, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.URL = url
	return nil
}

// MarshalJSON transforms the url to a String
func (u URL) MarshalJSON() ([]byte, error) {
	if u.URL == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(u.URL.String())
}
