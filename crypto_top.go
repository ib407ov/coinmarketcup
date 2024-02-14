package CoinMarketCup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinData struct {
	Symbol  string  `json:"symbol"`
	CmcRank float64 `json:"cmc_rank"`
}

type CoinMarketCapResponse struct {
	Data []CoinData
}

func GetDataCryptoTop() (CoinMarketCapResponse, error) {
	key := "1d4c8f1c-3c85-407e-83f8-2b22df3230cc"
	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?CMC_PRO_API_KEY=%s&start=1&limit=5000", key)
	resp, err := http.Get(url)

	var emptyResponse CoinMarketCapResponse
	if err != nil {
		fmt.Println("Помилка при виконанні HTTP-запиту:", err)
		return emptyResponse, err
	}
	defer resp.Body.Close()

	// Read body request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Помилка при читанні відповіді:", err)
		return emptyResponse, err
	}

	// Unmarshal JSON
	var cmcResponse CoinMarketCapResponse
	err = json.Unmarshal(body, &cmcResponse)
	if err != nil {
		fmt.Println("Помилка при розборі JSON:", err)
		return emptyResponse, err
	}

	return cmcResponse, nil
}
