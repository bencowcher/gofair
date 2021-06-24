package gofair

import (
	"fmt"
	"time"
)

type TimeRangeFilter struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type MarketFilter struct {
	TextQuery          string          `json:"textQuery,omitempty"`
	EventTypeIds       []string        `json:"eventTypeIds,omitempty"`
	MarketCountries    []string        `json:"marketCountries,omitempty"`
	MarketIds          []string        `json:"marketIds,omitempty"`
	EventIds           []string        `json:"eventIds,omitempty"`
	CompetitionIds     []string        `json:"competitionIds,omitempty"`
	BSPOnly            bool            `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  bool            `json:"turnInPLayEnabled,omitempty"`
	InPlayOnly         bool            `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []string        `json:"marketBettingTypes,omitempty"`
	MarketTypeCodes    []string        `json:"marketTypeCodes,omitempty"`
	MarketStartTime    TimeRangeFilter `json:"marketStartTime,omitempty"`
	WithOrders         string          `json:"withOrders,omitempty"`
}

type Params struct {
	MarketFilter     MarketFilter `json:"filter,omitempty"`
	MaxResults       int          `json:"maxResults,omitempty"`
	Granularity      string       `json:"granularity,omitempty"`
	MarketProjection []string     `json:"marketProjection,omitempty"`
	Sort             string       `json:"sort,omitempty"`
	Locale           string       `json:"locale,omitempty"`
}

type ExBestOffersOverrides struct {
	BestPricesDepth          int     `json:"bestPricesDepth,omitempty"`
	RollupModel              string  `json:"rollupModel,omitempty"`
	RollupLimit              int     `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float64 `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int     `json:"rollupLiabilityFactor,omitempty"`
}

type PriceProjection struct {
	PriceData             []string              `json:"priceData,omitempty"`
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise,omitempty"`
	RolloverStakes        bool                  `json:"rolloverStakes,omitempty"`
}

type ListMarketBookRequest struct {
	MarketIds                     []string        `json:"marketIds,omitempty"`
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"`
	MatchProjection               string          `json:"matchProjection,omitempty"`
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  *time.Time      `json:"matchedSince,omitempty"`
	BetIDs                        []string        `json:"betIds,omitempty"`
}

type Decimal float64

func (n Decimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

type MarketVersion struct {
	Version int64 `json:"version,omitempty"`
}

type PlaceOrderRequest struct {
	MarketID            string             `json:"marketId,omitempty"`
	Instructions        []PlaceInstruction `json:"instructions,omitempty"`
	CustomerRef         string             `json:"customerRef,omitempty"`
	MarketVersion       *MarketVersion     `json:"marketVersion,omitempty"`
	CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
	Async               bool               `json:"async,omitempty"`
}

type PlaceInstruction struct {
	OrderType          string              `json:"orderType,omitempty"`
	SelectionID        int64               `json:"selectionId,omitempty"`
	Handicap           Decimal             `json:"handicap,omitempty"`
	Side               string              `json:"side,omitempty"`
	LimitOrder         *LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string              `json:"customerOrderRef,omitempty"`
}

type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef,omitempty"`
	Status             string                   `json:"status,omitempty"`
	ErrorCode          string                   `json:"errorCode,omitempty"`
	MarketID           string                   `json:"marketId,omitempty"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports,omitempty"`
}

type PlaceInstructionReport struct {
	Status              string           `json:"status,omitempty"`
	ErrorCode           string           `json:"errorCode,omitempty"`
	OrderStatus         string           `json:"orderStatus,omitempty"`
	Instruction         PlaceInstruction `json:"instruction,omitempty"`
	BetID               string           `json:"betId,omitempty"`
	PlacedDate          time.Time        `json:"placedDate,omitempty"`
	AveragePriceMatched float64          `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64          `json:"sizeMatched,omitempty"`
}

type LimitOrder struct {
	Size            Decimal `json:"size,omitempty"`
	Price           Decimal `json:"price,omitempty"`
	PersistenceType string  `json:"persistenceType,omitempty"`
	TimeInForce     string  `json:"timeInForce,omitempty"`
	MinFillSize     Decimal `json:"minFillSize,omitempty"`
	BetTargetType   string  `json:"betTargetType,omitempty"`
	BetTargetSize   Decimal `json:"betTargetSize,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability Decimal `json:"liability,omitempty"`
	Price     Decimal `json:"price,omitempty"`
}

type MarketOnCloseOrder struct {
	Liability Decimal `json:"liability,omitempty"`
}
