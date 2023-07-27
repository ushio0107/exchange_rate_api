package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func NewExchangeRates() *ExchangeRates {
	return &ExchangeRates{}
}

func (a *API) APIExchangeRates() *ExchangeRates {
	return a.er
}

func (er *ExchangeRates) ConvertCurrency(source, target string, amount float64) (float64, error) {
	source = strings.ToLower(source)
	sourceRates, found := er.Currencies[source]
	if !found {
		return 0, errors.New("Failed to find source rate")
	}

	switch target {
	case "TWD":
		return amount * sourceRates.TWD, nil
	case "JPY":
		return amount * sourceRates.JPY, nil
	case "USD":
		return amount * sourceRates.USD, nil
	default:
		return 0, errors.New("Failed to find target rate")
	}
}

func FormatAmount(amount float64) string {
	return fmt.Sprintf("$%s", strconv.FormatFloat(amount, 'f', 2, 64))
}
