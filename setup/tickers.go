package setup

import (
	"github.com/svarlamov/bintrad/models"
	"time"
)

func setupTickers() error {
	foo := models.Ticker{
		Id:     1,
		Ticker: "FOO",
		Name:   "Foo Inc.",
		Type:   "EQUITY",
	}
	err := foo.Create()
	if err != nil {
		return err
	}
	open := []float64{90, 91.3, 88.7, 93.7, 83.2}
	high := []float64{91.7, 91.5, 93.3, 93.9, 85.3}
	low := []float64{87.6, 86.5, 88.5, 84.1, 82.8}
	close := []float64{91.1, 87.3, 93.1, 84.3, 84.8}
	volume := []float64{1001, 800, 2000, 200, 5000}
	startTime := time.Now().AddDate(0, -1, 0)
	periodSeconds := 300
	for ind := range open {
		tickerData := models.TickerData{
			TickerId: foo.Id,
			OpensAt:  startTime.Add(time.Duration(ind) * time.Duration(periodSeconds) * time.Second),
			ClosesAt: startTime.Add(time.Duration(ind+1) * time.Duration(periodSeconds) * time.Second),
			Open:     open[ind],
			High:     high[ind],
			Low:      low[ind],
			Close:    close[ind],
			Volume:   volume[ind],
		}
		err = tickerData.Create()
		if err != nil {
			return err
		}
	}
	return nil
}
