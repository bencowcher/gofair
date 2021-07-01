package gofair

import (
	"time"
)

type EventType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EventTypeResult struct {
	MarketCount int       `json:"marketCount"`
	EventType   EventType `json:"eventType"`
}

type Competition struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CompetitionResult struct {
	MarketCount       int         `json:"marketCount"`
	CompetitionRegion string      `json:"competitionRegion"`
	Competition       Competition `json:"competition"`
}

type TimeRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type TimeRangeResult struct {
	MarketCount int       `json:"marketCount"`
	TimeRange   TimeRange `json:"timeRange"`
}

type Event struct {
	Id          string `json:"id"`
	OpenDate    string `json:"openDate"`
	TimeZone    string `json:"timezone"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Venue       string `json:"venue"`
}

type EventResult struct {
	MarketCount int   `json:"marketCount"`
	Event       Event `json:"event"`
}

type MarketTypeResult struct {
	MarketCount int    `json:"marketCount"`
	MarketType  string `json:"marketType"`
}

type CountryResult struct {
	MarketCount int    `json:"marketCount"`
	CountryCode string `json:"countryCode"`
}

type VenueResult struct {
	MarketCount int    `json:"marketCount"`
	Venue       string `json:"venue"`
}

type MarketCatalogueDescription struct {
	BettingType        string    `json:"bettingType"`
	BSPMarket          bool      `json:"bspMarket"`
	DiscountAllowed    bool      `json:"discountAllowed"`
	MarketBaseRate     float32   `json:"marketBaseRate"`
	MarketTime         time.Time `json:"marketTime"`
	MarketType         string    `json:"marketType"`
	PersistenceEnabled bool      `json:"persistenceEnabled"`
	Regulator          string    `json:"regulator"`
	Rules              string    `json:"rules"`
	RulesHasDate       bool      `json:"rulesHasDate"`
	SuspendDate        time.Time `json:"suspendTime"`
	TurnInPlayEnabled  bool      `json:"turnInPlayEnabled"`
	Wallet             string    `json:"wallet"`
	EachWayDivisor     float32   `json:"eachWayDivisor"`
	Clarifications     string    `json:"clarifications"`
}

type Metadata struct {
	RunnerId int `json:"runnerId"`
}

type RunnerCatalogue struct {
	SelectionId  int     `json:"selectionId"`
	RunnerName   string  `json:"runnerName"`
	SortPriority int     `json:"sortPriority"`
	Handicap     float32 `json:"handicap"`
	//Metadata		*metadata	`json:"metadata"`  //todo
}

type PriceSize struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type MarketBook struct {
	MarketID              string  `json:"marketId"`
	IsMarketDataDelayed   bool    `json:"isMarketDataDelayed"`
	Status                string  `json:"status"`
	BetDelay              int     `json:"betDelay"`
	BspReconciled         bool    `json:"bspReconciled"`
	Complete              bool    `json:"complete"`
	Inplay                bool    `json:"inplay"`
	NumberOfWinners       int     `json:"numberOfWinners"`
	NumberOfRunners       int     `json:"numberOfRunners"`
	NumberOfActiveRunners int     `json:"numberOfActiveRunners"`
	TotalMatched          float64 `json:"totalMatched"`
	TotalAvailable        float64 `json:"totalAvailable"`
	CrossMatching         bool    `json:"crossMatching"`
	RunnersVoidable       bool    `json:"runnersVoidable"`
	Version               int64   `json:"version"`
	Runners               []struct {
		SelectionID  int64   `json:"selectionId"`
		Handicap     float64 `json:"handicap"`
		Status       string  `json:"status"`
		TotalMatched float64 `json:"totalMatched"`
		Ex           struct {
			AvailableToBack []PriceSize `json:"availableToBack"`
			AvailableToLay  []PriceSize `json:"availableToLay"`
			TradedVolume    []PriceSize `json:"tradedVolume"`
		} `json:"ex"`
		SP struct {
			NearPrice         float64     `json:"nearPrice"`
			FarPrice          float64     `json:"farPrice"`
			ActualSP          interface{} `json:"actualSP"`
			BackStakeTaken    []PriceSize `json:"backStakeTaken"`
			LayLiabilityTaken []PriceSize `json:"layLiabilityTaken"`
		} `json:"sp"`
		Orders  []Order `json:"orders"`
		Matches []Match `json:"matches"`
	} `json:"runners"`
}

type Order struct {
	BetID           string    `json:"betId"`
	OrderType       string    `json:"orderType"`
	OrderStatus     string    `json:"orderStatus"`
	PersistenceType string    `json:"persistenceType"`
	Side            string    `json:"side"`
	Size            Decimal   `json:"size"`
	Price           Decimal   `json:"price"`
	BSPLiability    Decimal   `json:"bspLiability"`
	PlacedDate      time.Time `json:"placedDate"`

	AvgPriceMatched Decimal `json:"avgPriceMatched"`
	SizeMatched     Decimal `json:"sizeMatched"`
	SizeRemaining   Decimal `json:"sizeRemaining"`
	SizeLapsed      Decimal `json:"sizeLapsed"`
	SizeCancelled   Decimal `json:"sizeCancelled"`
	SizeVoided      Decimal `json:"sizeVoided"`

	CustomerOrderRef    string `json:"customerOrderRef"`
	CustomerStrategyRef string `json:"customerStrategyRef"`
}
type Match struct {
	BetID       string    `json:"betId"`
	MatchID     string    `json:"matchId"`
	Side        string    `json:"side"`
	Size        Decimal   `json:"size"`
	Price       Decimal   `json:"price"`
	MatchedDate time.Time `json:"matchDate"`
}

type MarketCatalogue struct {
	MarketId                   string                     `json:"marketId"`
	MarketName                 string                     `json:"marketName"`
	TotalMatched               float32                    `json:"totalMatched"`
	MarketStartTime            time.Time                  `json:"marketStartTime"`
	Competition                Competition                `json:"competition"`
	Event                      Event                      `json:"event"`
	EventType                  EventType                  `json:"eventType"`
	MarketCatalogueDescription MarketCatalogueDescription `json:"description"`
	Runners                    []RunnerCatalogue          `json:"runners"`
}

// PlaceOrders to place new orders into market.
func (b *Betting) PlaceOrders(req PlaceOrderRequest) (*PlaceExecutionReport, error) {
	url := createUrl(api_betting_url, "placeOrders/")
	res := PlaceExecutionReport{}

	err := b.client.Request(url, &req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (b *Betting) ListEventTypes(filter MarketFilter) ([]EventTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEventTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []EventTypeResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCompetitions(filter MarketFilter) ([]CompetitionResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCompetitions/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []CompetitionResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListTimeRanges(filter MarketFilter, granularity string) ([]TimeRangeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listTimeRanges/")

	// build request
	params := new(Params)
	params.MarketFilter = filter
	params.Granularity = granularity

	var response []TimeRangeResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListEvents(filter MarketFilter) ([]EventResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEvents/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []EventResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketTypes(filter MarketFilter) ([]MarketTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []MarketTypeResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCountries(filter MarketFilter) ([]CountryResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCountries/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []CountryResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListVenues(filter MarketFilter) ([]VenueResult, error) {
	// create url
	url := createUrl(api_betting_url, "listVenues/")

	// build request
	params := new(Params)
	params.MarketFilter = filter

	var response []VenueResult

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketCatalogue(filter MarketFilter, marketProjection []string, sort string, maxResults int) ([]MarketCatalogue, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketCatalogue/")

	// build request
	params := new(Params)
	params.MarketFilter = filter
	params.MarketProjection = marketProjection
	params.Sort = sort
	params.MaxResults = maxResults

	var response []MarketCatalogue

	// make request
	err := b.client.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketBook(req ListMarketBookRequest) ([]MarketBook, error) {
	url := createUrl(api_betting_url, "listMarketBook/")

	markets := []MarketBook{}

	err := b.client.Request(url, &req, &markets)

	if err != nil {
		return nil, err
	}

	return markets, nil

}
