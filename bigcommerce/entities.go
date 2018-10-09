package bigcommerce

import (
	"strconv"
	"time"
)

// count is a generic object used for returning resource counts.
type count struct {
	Count int `json:"count"`
}

// AddressEntities defines a list of the AddressEntity object.
type AddressEntities []AddressEntity

// AddressEntity describes the address entity.
type AddressEntity struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Company        string `json:"company"`
	Street1        string `json:"street_1"`
	Street2        string `json:"street_2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zip            string `json:"zip"`
	Country        string `json:"country"`
	CountryIso2    string `json:"country_iso2"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	ShippingMethod string `json:"shipping_method,omitempty"`
}

// BCTime converts to/from Big Commerce Time format (RFC1123).
// BC returns "" to indicate unset (null) values.
// The internal time will be nil in those cases.
type BCTime struct {
	t *time.Time
}

// NewBCTime returns a wrapped time.
// The supplied time can be nil.
func NewBCTime(t *time.Time) BCTime {
	return BCTime{t: t}
}

// Time returns the wrapped time.
func (b BCTime) Time() *time.Time {
	return b.t
}

func (b *BCTime) UnmarshalJSON(text []byte) error {
	s, err := strconv.Unquote(string(text))
	if err != nil {
		return err
	}
	// In BC an empty string is null.
	if s == "" {
		return nil
	}
	t, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return err
	}
	*b = BCTime{t: &t}
	return nil
}

func (b BCTime) MarshalJSON() (text []byte, err error) {
	if b.t == nil || b.t.IsZero() {
		return []byte(`""`), nil
	}
	s := b.t.Format(time.RFC1123Z)
	return []byte(strconv.Quote(s)), nil
}
