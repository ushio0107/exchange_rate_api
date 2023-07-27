package api

import (
	"github.com/spf13/viper"
)

// NewSettings gets the config of the exchangeRates, and return it to the api.
func NewSettings(cfg *Config) (*API, error) {
	er := NewExchangeRates()
	if err := cfg.ReadConfig(er); err != nil {
		return nil, err
	}

	return &API{
		er: er,
	}, nil
}

// ReadConfig reads the config file.
func (cfg *Config) ReadConfig(er *ExchangeRates) error {
	vp := viper.New()
	vp.SetConfigFile(cfg.CfgFile)
	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	if err := vp.Unmarshal(er); err != nil {
		return err
	}

	return nil
}
