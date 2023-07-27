package api

type API struct {
	Cfg *Config
	er  *ExchangeRates
}

type Config struct {
	CfgFile string
}

type ExchangeRates struct {
	Currencies map[string]currency `json:"currencies"`
}

type currency struct {
	TWD float64 `json:"TWD"`
	JPY float64 `json:"JPY"`
	USD float64 `json:"USD"`
}

type Response struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}
