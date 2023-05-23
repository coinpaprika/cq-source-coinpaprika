package client

type Spec struct {
	AccessToken string `json:"access_token"`
	StartDate   string `json:"start_date"`
	Interval    string `json:"interval"`
}
