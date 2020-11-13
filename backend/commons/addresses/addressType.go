package addresses

type Address struct {
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
	Zipcode      string `json:"zipcode"`
	City         string `json:"city"`
	Country      string `json:"country"`
}

type GeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
