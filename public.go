package poloniex

import "encoding/json"

// Volume represents the data returned from a 24hr volume request to poloniex.
type Volume map[string]json.RawMessage

// DailyVolume returns volume data for all markets over the past 24 hours.
func (c *Client) DailyVolume() {}

// TickerData represents the data returned from a ticker request to poloniex.
type TickerData map[string]Ticker

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
func (c *Client) Ticker() {}

// TradeHistory represents the data returned from a trade history request to poloniex.
type TradeHistory struct {
	GlobalTradeID int     `json:"globalTradeID,int"`
	TradeID       int     `json:"tradeID,int"`
	Date          string  `json:"date,string"`
	Type          string  `json:"type,string"`
	Rate          float64 `json:"rate,float64"`
	Amount        float64 `json:"amount,float64"`
	Total         float64 `json:"total,float64"`
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

// Currencies represents the data returned from a currencies request to poloniex.
type Currencies map[string]Currency

// Currency represents the data for each currency in Currencies.
type Currency struct {
	ID                   int     `json:"id,int"`
	Name                 string  `json:"name,string"`
	TransactionFee       float64 `json:"txFee,float64"`
	MinimumConfirmations int     `json:"minConf,int"`
	Disabled             int     `json:"disabled,int"`
	Delisted             int     `json:"delisted,int"`
	Frozen               int     `json:"frozen,int"`
}

// Currencies returns the list of currencies available on poloniex.
func (c *Client) Currencies() {}

// LoanOrders represents the information returned from a loan orders request to poloniex.
type LoanOrders map[string][]Loan

// Loan contains the details for each loan found in LoanOrders.
type Loan struct {
	Rate     float64 `json:"rate,float64"`
	Amount   float64 `json:"amount,float64"`
	RangeMin int     `json:"rangeMin,int"`
	RangeMax int     `json:"rangeMax,int"`
}

// LoanOrders returns the list of loan orders for a given currency.
func (c *Client) LoanOrders(currency string) {}

// ChartData represents the data returned from a chart data request from poloniex
type ChartData struct {
	Date        int     `json:"date,int"`
	High        float64 `json:"high,float64"`
	Low         float64 `json:"low,float64"`
	Open        float64 `json:"open,float64"`
	Close       float64 `json:"close,float64"`
	Volume      float64 `json:"volume,float64"`
	QuoteVolume float64 `json:"quoteVolume,float64"`
	WeightedAvg float64 `json:"weightedAverage,float64"`
}

// ChartData returns the chart data for a given currency pair and timeframe.
func (c *Client) ChartData(pair string, start, end, period int) {}
