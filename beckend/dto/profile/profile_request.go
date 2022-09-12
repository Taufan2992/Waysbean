package profiledto

type ProfileResponse struct {
	ID       int    `json:"id"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
}
