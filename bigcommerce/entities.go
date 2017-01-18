package bigcommerce

// count is a generic object used for returning resource counts.
type count struct {
	Count int `json:"count"`
}

// AddressEntities defines a list of the AddressEntity object.
type AddressEntities []AddressEntity

// AddressEntity describes the address entity.
type AddressEntity struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Company     string `json:"company"`
	Street1     string `json:"street_1"`
	Street2     string `json:"street_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	CountryIso2 string `json:"country_iso2"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
