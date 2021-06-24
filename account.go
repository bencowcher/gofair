package gofair

type AccountFundsResponse struct {
	AvailableToBetBalance float64 `json:"availableToBetBalance"`
	Exposure              float64 `json:"exposure"`
	RetainedCommission    float64 `json:"retainedCommission"`
	ExposureLimit         float64 `json:"exposureLimit"`
	DiscountRate          float64 `json:"discountRate"`
	PointsBalance         int     `json:"pointsBalance"`
	Wallet                string  `json:"wallet"`
}

func (acc Account) GetAccountFunds() (*AccountFundsResponse, error) {
	url := createUrl(api_account_url, "getAccountFunds/")

	var res AccountFundsResponse

	err := acc.client.Request(url, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil

}
