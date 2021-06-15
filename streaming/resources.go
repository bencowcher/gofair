package streaming

import "time"

type RunnerDefinition struct {
	SortPriority     int32     `json:"sortPriority"`
	RemovalDate      time.Time `json:"removalDate"`
	SelectionId      int64     `json:"id"`
	Handicap         float64   `json:"hc"`
	AdjustmentFactor float64   `json:"adjustmentFactor"`
	Bsp              float64   `json:"bsp"`
	Status           string    `json:"status"`
}

type MarketDefinition struct {
	Venue                  string             `json:"venue"`
	SettledTime            string             `json:"settledTime"`
	Timezone               string             `json:"timezone"`
	EachWayDivisor         float64            `json:"eachWayDivisor"`
	Regulators             []string           `json:"regulators"`
	MarketType             string             `json:"marketType"`
	MarketBaseRate         float64            `json:"marketBaseRate"`
	NumberOfWinners        int32              `json:"numberOfWinners"`
	CountryCode            string             `json:"countryCode"`
	LineMaxUnit            float64            `json:"lineMaxUnit"`
	Inplay                 bool               `json:"inPlay"`
	BetDelay               int32              `json:"betDelay"`
	BspMarket              bool               `json:"bspMarket"`
	BettingType            string             `json:"bettingType"`
	NumberOfActiveRunners  int32              `json:"numberOfActiveRunners"`
	LineMinUnit            float64            `json:"lineMinUnit"`
	EventId                string             `json:"eventId"`
	CrossMatching          bool               `json:"crossMatching"`
	RunnersVoidable        bool               `json:"runnersVoidable"`
	TurnInPlayEnabled      bool               `json:"turnInPlayEnabled"`
	SuspendTime            string             `json:"suspendTime"`
	DiscountAllowed        bool               `json:"discountAllowed"`
	PersistenceEnabled     bool               `json:"persistenceEnabled"`
	Runners                []RunnerDefinition `json:"runners"`
	Version                int64              `json:"version"`
	EventTypeId            string             `json:"eventTypeId"`
	Complete               bool               `json:"complete"`
	OpenDate               string             `json:"openDate"`
	MarketTime             string             `json:"marketTime"`
	BspReconciled          bool               `json:"bspReconciled"`
	LineInterval           float64            `json:"lineInterval"`
	Status                 string             `json:"status"`
	PriceLadderDescription string             `json:"priceLadderDescription"`
	KeyLineDefinition      string             `json:"keyLineDefinition"`
	Name                   string             `json:"name"`
}

type RunnerChange struct {
	SelectionId     int64       `json:"id"`
	Handicap        float64     `json:"hc"`
	TradedVolume    float64     `json:"tv"`
	LastTradedPrice float64     `json:"ltp"`
	Traded          [][]float64 `json:"trd"`
	//StartingPriceNear 		float64	`json:"spn"`
	//StartingPriceFar 		float64	`json:"spf"`
	StartingPriceBack          [][]float64 `json:"spb"`
	StartingPriceLay           [][]float64 `json:"spl"`
	AvailableToBack            [][]float64 `json:"atb"`
	AvailableToLay             [][]float64 `json:"atl"`
	BestAvailableToBack        [][]float64 `json:"batb"`
	BestAvailableToLay         [][]float64 `json:"batl"`
	BestDisplayAvailableToBack [][]float64 `json:"bdatb"`
	BestDisplayAvailableToLay  [][]float64 `json:"bdatl"`
}

type MarketChange struct {
	Image            bool              `json:"img"`
	Conflated        bool              `json:"con"`
	MarketId         string            `json:"id"`
	TradedVolume     float64           `json:"tv"`
	RunnerChange     []RunnerChange    `json:"rc"`
	MarketDefinition *MarketDefinition `json:"marketDefinition"`
}

type MarketChangeMessage struct {
	MarketChanges []MarketChange `json:"mc"`
	PublishTime   int            `json:"pt"`
	Operation     string         `json:"op"`
	ChangeType    string         `json:"ct"`
	InitialClk    string         `json:"initialClk"`
	Clk           string         `json:"clk"`
	HeartbeatMs   int64          `json:"heartbeatMs"`
	ConflateMs    int64          `json:"conflateMs"`
	SegmentType   string         `json:"segmentType"`
	Status        int32          `json:"status"`
}

type MarketBook struct {
	PublishTime           int
	MarketId              string
	Status                string
	BetDelay              int32
	BspReconciled         bool
	Complete              bool
	Inplay                bool
	NumberOfWinners       int32
	NumberOfRunners       int
	NumberOfActiveRunners int32
	TotalMatched          float64
	CrossMatching         bool
	RunnersVoidable       bool
	Version               int64
	Runners               []Runner
}

type Runner struct {
	SelectionId      int64
	Handicap         float64
	Status           string
	AdjustmentFactor float64
	LastPriceTraded  float64
	TotalMatched     float64
	RemovalDate      time.Time
	EX               ExchangePrices
}

type ExchangePrices struct {
	AvailableToBack []PriceSize
	AvailableToLay  []PriceSize
	TradedVolume    []PriceSize
}
