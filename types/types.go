package types

type TokenData struct {
	TokenName       string `json:"token_name"`
	TokenPrice      string `json:"token_price"`
	PercentUpOrDown string `json:"percent_up_or_down"`
	UpOrDown        string `json:"up_or_down"`
	ChartUrl        string `json:"chart_url"`
	TokenSymbol     string `json:"token_symbol"`
	TokenLogo       string `json:"token_logo"`
}

type Beneficiary struct {
	BPhoneNo        string
	BName           string
	ServiceProvider string
}
