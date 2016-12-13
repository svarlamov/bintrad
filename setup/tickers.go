package setup

import (
	"github.com/svarlamov/bintrad/models"
	"time"
	"net/http"
	"encoding/json"
)

func setupTickers() error {
	tickers := []string{"JBLU"}
	for _, key := range tickers {
		resp := models.YahooFinanceResponse{}
		err := getJson("https://query1.finance.yahoo.com/v7/finance/chart/" + key + "?range=1y&interval=30m&indicators=quote&includeTimestamps=true&includePrePost=false&corsDomain=finance.yahoo.com", &resp)
		if err != nil {
			return err
		}
		ticker := models.Ticker{
			Ticker: resp.Chart.Result[0].Meta.Symbol,
			Name:   resp.Chart.Result[0].Meta.Symbol,
			Type:   resp.Chart.Result[0].Meta.InstrumentType,
		}
		err = ticker.Create()
		if err != nil {
			return err
		}
		startTime := time.Unix(resp.Chart.Result[0].Timestamp[0], 0)
		periodSeconds := 1800
		for ind := range resp.Chart.Result[0].Timestamp {
			tickerData := models.TickerData{
				TickerId: ticker.Id,
				OpensAt:  startTime.Add(time.Duration(ind) * time.Duration(periodSeconds) * time.Second),
				ClosesAt: startTime.Add(time.Duration(ind+1) * time.Duration(periodSeconds) * time.Second),
				Period:   int64(periodSeconds),
				Open:     resp.Chart.Result[0].Indicators.Quote[0].Open[ind],
				High:     resp.Chart.Result[0].Indicators.Quote[0].High[ind],
				Low:      resp.Chart.Result[0].Indicators.Quote[0].Low[ind],
				Close:    resp.Chart.Result[0].Indicators.Quote[0].Close[ind],
				Volume:   resp.Chart.Result[0].Indicators.Quote[0].Volume[ind],
			}
			err = tickerData.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}