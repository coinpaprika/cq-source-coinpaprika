package client

type Spec struct {
	AccessToken  string   `json:"access_token"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Interval     string   `json:"interval"`
	Tickers      []string `json:"tickers"`
	ApiDebug     bool     `json:"api_debug"`
	RateNumber   int      `json:"rate_number"`
	RateDuration string   `json:"rate_duration"`
}
