package marshaler_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/xescugc/marshaler"
)

type testURL struct {
	URL marshaler.URL `json:"url"`
}

func TestNewURL(t *testing.T) {
	u, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("error parsing the url %s", err)
	}
	ts := testURL{URL: marshaler.URL{URL: u}}
	te := testURL{URL: marshaler.NewURL(u)}

	if ts != te {
		t.Fatalf("expect %+v to be equal to %+v", ts, te)
	}
}

func TestURLMarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, err := url.Parse("http://example.com")
		if err != nil {
			t.Fatalf("error parsing the url %s", err)
		}

		eu := fmt.Sprintf(`{"url":"%s"}`, u)
		ts := testURL{URL: marshaler.URL{URL: u}}

		b, err := json.Marshal(ts)
		if err != nil {
			t.Fatalf("error marshalling the struct %s", err)
		}

		if string(b) != eu {
			t.Errorf("expected %q to be equal to %q", b, eu)
		}
	})
}

func TestURLUnmarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, err := url.Parse("http://example.com")
		if err != nil {
			t.Fatalf("error parsing the url %s", err)
		}

		eu := []byte(fmt.Sprintf(`{"url": "%s"}`, u))
		ets := testURL{URL: marshaler.URL{URL: u}}
		var ts testURL

		err = json.Unmarshal(eu, &ts)
		if err != nil {
			t.Fatalf("error unmarshaling the struct %s", err)
		}

		if !reflect.DeepEqual(ets, ts) {
			t.Errorf("expected %+v to be equal to %+v", ets, ts)
		}
	})
}
