package marshaler

import (
	"encoding/json"

	"golang.org/x/text/currency"
)

// URL extends the currency.Unit to add implementations
// to the Marshaler and Unmarshaler interfaces
type CurrencyUnit struct {
	currency.Unit
}

// NewCurrencyUnit returns a new marshaler.CurrencyUnit from u
func NewCurrencyUnit(u currency.Unit) CurrencyUnit {
	return CurrencyUnit{
		Unit: u,
	}
}

// UnmarshalJSON transforms the b to a string
func (u *CurrencyUnit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	unit, err := currency.ParseISO(s)
	if err != nil {
		return err
	}

	u.Unit = unit
	return nil
}

// MarshalJSON transforms the CurrencyUnit to a String
// if it's not defined it'll use the default one which
// is currency.XXX
func (u CurrencyUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Unit.String())
}
