package organizations

type Organization struct {
	OrganizationId string `json:"organizationId"`
	DisplayName    string `json:"displayName"`
	StreetName     string `json:"streetName"`
	Number         string `json:"number"`
	ZipCode        string `json:"zipCode"`
	city           string `json:"city"`
	country        string `json:"country"`
}

type AddOrganization struct {
	Organization
	UserId string `json:"userId"`
}
