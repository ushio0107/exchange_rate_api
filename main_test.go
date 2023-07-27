package main

import (
	"encoding/json"
	"exchange_rate_api/api"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHandleCurrencyConversion(t *testing.T) {
	a, err := api.NewServer(&api.Config{
		CfgFile: "./config/currencies_config_test.json",
	})
	if err != nil {
		log.Fatal("Fail: ", err)
	}

	er := a.APIExchangeRates()
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		sourceCurrencies := []string{"USD", "JPY", "TWD"}
		targetCurrencies := []string{"USD", "JPY", "TWD"}
		source := sourceCurrencies[rand.Intn(len(sourceCurrencies))]
		target := targetCurrencies[rand.Intn(len(targetCurrencies))]
		// Generate a random amount with two decimal places.
		amount := float64(rand.Intn(10000)*100) + float64(rand.Intn(100))/100.0

		url := "/convert?source=" + source + "&target=" + target + "&amount=" + api.FormatAmount(amount)
		req, err := http.NewRequest("GET", url, nil)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		a.NewRouter().ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var response api.Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		convertedAmount, err := er.ConvertCurrency(source, target, amount)
		require.NoError(t, err)
		expectedResponse := api.Response{
			Msg:    "success",
			Amount: api.FormatAmount(convertedAmount),
		}
		require.Equal(t, expectedResponse, response)
	}
}
