package marshaler_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/xescugc/marshaler"
	"golang.org/x/text/currency"
)

type testCurrencyUnit struct {
	Unit marshaler.CurrencyUnit `json:"unit"`
}

func TestNewCurrencyUnit(t *testing.T) {
	ts := testCurrencyUnit{Unit: marshaler.CurrencyUnit{Unit: currency.EUR}}
	te := testCurrencyUnit{Unit: marshaler.NewCurrencyUnit(currency.EUR)}

	if ts != te {
		t.Fatalf("expect %+v to be equal to %+v", ts, te)
	}
}

func TestCurrencyUnitMarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		eu := fmt.Sprintf(`{"unit":"%s"}`, currency.EUR.String())
		ts := testCurrencyUnit{Unit: marshaler.CurrencyUnit{Unit: currency.EUR}}

		b, err := json.Marshal(ts)
		if err != nil {
			t.Fatalf("error marshalling the struct %s", err)
		}

		if string(b) != eu {
			t.Errorf("expected %q to be equal to %q", b, eu)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		eu := fmt.Sprintf(`{"unit":"%s"}`, currency.XXX.String())
		ts := testCurrencyUnit{}

		b, err := json.Marshal(ts)
		if err != nil {
			t.Fatalf("error marshalling the struct %s", err)
		}

		if string(b) != eu {
			t.Errorf("expected %q to be equal to %q", b, eu)
		}
	})
}

func TestCurrencyUnitUnmarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		eu := []byte(fmt.Sprintf(`{"unit": "%s"}`, currency.EUR.String()))
		ets := testCurrencyUnit{Unit: marshaler.CurrencyUnit{Unit: currency.EUR}}
		var ts testCurrencyUnit

		err := json.Unmarshal(eu, &ts)
		if err != nil {
			t.Fatalf("error unmarshaling the struct %s", err)
		}

		if !reflect.DeepEqual(ets, ts) {
			t.Errorf("expected %+v to be equal to %+v", ets, ts)
		}
	})
	t.Run("Error", func(t *testing.T) {
		eu := []byte(`{"unit": "potato"}`)
		var ts testCurrencyUnit

		err := json.Unmarshal(eu, &ts)
		if err.Error() != "currency: tag is not well-formed" {
			t.Fatalf("error unmarshaling the struct %s", err)
		}
	})
}
