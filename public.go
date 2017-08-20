package poloniex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type data map[string]*json.RawMessage

type VolumeData struct {
	Volumes   map[string]map[string]float64
	TotalBTC  float64 `json:"totalBTC,string"`
	TotalETH  float64 `json:"totalETH,string"`
	TotalUSDT float64 `json:"totalUSDT,string"`
	TotalXMR  float64 `json:"totalXMR,string"`
	TotalXUSD float64 `json:"totalXUSD,string"`
}

// DailyVolume returns volume data for all markets over the past 24 hours.
func (c *Client) DailyVolume() (VolumeData, error) {
	d := make(data)
	if err := c.req("GET", "return24hVolume", nil, &d); err != nil {
		return VolumeData{}, err
	}
	res := VolumeData{Volumes: make(map[string]map[string]float64)}
	for k, v := range d {
		switch k {
		case "totalBTC":
			if err := parseTotalVolume(k, v, &res.TotalBTC); err != nil {
				return VolumeData{}, err
			}
		case "totalETH":
			if err := parseTotalVolume(k, v, &res.TotalETH); err != nil {
				return VolumeData{}, err
			}
		case "totalUSDT":
			if err := parseTotalVolume(k, v, &res.TotalUSDT); err != nil {
				return VolumeData{}, err
			}
		case "totalXMR":
			if err := parseTotalVolume(k, v, &res.TotalXMR); err != nil {
				return VolumeData{}, err
			}
		case "totalXUSD":
			if err := parseTotalVolume(k, v, &res.TotalXUSD); err != nil {
				return VolumeData{}, err
			}
		default:
			res.Volumes[k] = make(map[string]float64)
			t := make(map[string]*json.Number)
			err := json.Unmarshal(*v, &t)
			if err != nil {
				return VolumeData{}, err
			}
			for x, y := range t {
				n, err := y.Float64()
				if err != nil {
					return VolumeData{}, err
				}
				res.Volumes[k][x] = n
			}
		}
	}
	return res, nil
}

func parseTotalVolume(name string, data *json.RawMessage, s *float64) error {
	n := new(json.Number)
	if err := json.Unmarshal(*data, &n); err != nil {
		return errors.Wrapf(err, "error parsing %s", name)
	}
	var err error
	*s, err = n.Float64()
	return err
}

// Ticker contains the details for each ticker in TickerData.
type Ticker struct {
	Last          float64 `json:"last,string"`
	LowestAsk     float64 `json:"lowestAsk,string"`
	HighestBid    float64 `json:"highestBid,string"`
	PercentChange float64 `json:"percentChange,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
	ID            int64   `json:"id"`
	LowDay        float64 `json:"low24hr,string"`
	HighDay       float64 `json:"high24hr,string"`
	IsFrozen      int     `json:"isFrozen,string"`
}

// Ticker returns ticker data for all markets.
func (c *Client) Ticker() (map[string]Ticker, error) {
	d := make(data)
	if err := c.req("GET", "returnTicker", nil, &d); err != nil {
		return nil, err
	}
	res := make(map[string]Ticker)
	for k, v := range d {
		t := new(Ticker)
		if err := json.Unmarshal(*v, t); err != nil {
			return nil, err
		}
		res[k] = *t
	}
	return res, nil
}

// TradeHistory represents the data returned from a trade history request to poloniex.
type TradeHistory struct {
	GlobalTradeID int     `json:"globalTradeID"`
	TradeID       int     `json:"tradeID"`
	Date          string  `json:"date"`
	Type          string  `json:"type"`
	Rate          float64 `json:"rate"`
	Amount        float64 `json:"amount"`
	Total         float64 `json:"total"`
}

// TradeHistory returns the past trades for a given currency pair and timeframe.
func (c *Client) TradeHistory(pair string, start, end int) {}

// OrderBook represents the data returned from an order book request to poloniex.
type OrderBook struct {
	Asks [][]interface{} `json:"asks"`
	Bids [][]interface{} `json:"bids"`
}

// OrderBook returns the orderbook data for a given currency pair.
func (c *Client) OrderBook(pair string, depth int) {}

// Currency represents the data for each currency in Currencies.
type Currency struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	TransactionFee       float64 `json:"txFee,string"`
	MinimumConfirmations int     `json:"minConf"`
	Disabled             int     `json:"disabled"`
	Delisted             int     `json:"delisted"`
	Frozen               int     `json:"frozen"`
}

// Currencies returns the list of currencies available on poloniex.
func (c *Client) Currencies() (map[string]Currency, error) {
	d := make(data)
	if err := c.req("GET", "returnCurrencies", nil, &d); err != nil {
		return nil, err
	}
	res := make(map[string]Currency)
	for k, v := range d {
		cur := new(Currency)
		if err := json.Unmarshal(*v, cur); err != nil {
			return nil, err
		}
		res[k] = *cur
	}
	return res, nil
}

// LoanOrders represents the information returned from a loan orders request to poloniex.
type LoanOrders map[string][]Loan

// Loan contains the details for each loan found in LoanOrders.
type Loan struct {
	Rate     float64 `json:"rate"`
	Amount   float64 `json:"amount"`
	RangeMin int     `json:"rangeMin"`
	RangeMax int     `json:"rangeMax"`
}

// LoanOrders returns the list of loan orders for a given currency.
func (c *Client) LoanOrders(currency string) {}

// ChartData represents the data returned from a chart data request from poloniex
type ChartData struct {
	Date        int     `json:"date"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Open        float64 `json:"open"`
	Close       float64 `json:"close"`
	Volume      float64 `json:"volume"`
	QuoteVolume float64 `json:"quoteVolume"`
	WeightedAvg float64 `json:"weightedAverage"`
}

// ChartData returns the chart data for a given currency pair and timeframe.
func (c *Client) ChartData(pair string, start, end, period int) {}
