package models

// BEGIN monthly aggregates
type AggregateTotal struct {
	Euros float32 `json:"euros,omitempty"`
	Kwh float32 `json:"kwh,omitempty"`
	RushHour float32 `json:"rush_hour,omitempty"`
	VallHour float32 `json:"vall_hour,omitempty"`
}

type AggregateDay struct {
	Day string `json:"day,omitempty"`
	Total map[string]AggregateTotal `json:"total,omitempty"`
	AvgPrice float32 `json:"avg_price,omitempty"`
	WeekDay int8 `json:"week_day,omitempty"`
	IsWeekend bool `json:"is_weekend,omitempty"`
	TotalHour map[string]float32 `json:"total_hour,omitempty"`
}

type EquivalenceConsume struct {
	KwCO2 float32 `json:"kw_co2,omitempty"`
	Km float32 `json:"km,omitempty"`
	BulbDays15W float32 `json:"bulb_days_15w,omitempty"`
}

type AggregateContractMonth struct {
	Total map[string]AggregateTotal `json:"total,omitempty"`
	Days map[string]AggregateDay `json:"days,omitempty"`
	AvgPriceKwh float32 `json:"avg_price_kwh,omitempty"`
  PercentConsumeSlot map[string]float32 `json:"percent_consume_slot,omitempty"`
  Week map[string]map[string]float32 `json:"week,omitempty"`
	EquivalenceConsume *EquivalenceConsume `json:"equivalence_consume,omitempty"`
	PercentConsumeBetweenWeek float32 `json:"percent_consume_between_week,omitempty"`
	PercentConsumeWeekend float32 `json:"percent_consume_weekend,omitempty"`
}

type AggregateContractMonthPayload struct {
	JsonPayload AggregateContractMonth `json:"payload"`
}

type AggregateContractMonthPayloadInterface struct {
	JsonPayload interface{} `json:"payload"`
}
// END monthly aggregates

// BEGIN yearly aggregates
type TotalSeason struct {
	Total map[string]AggregateTotal `json:"total,omitempty"`
	Label string `json:"label,omitempty"`
}

type TotalPerSeason struct {
	Autumn *TotalSeason `json:"autumn,omitempty"`
	Winter *TotalSeason `json:"winter,omitempty"`
	Spring *TotalSeason `json:"spring,omitempty"`
	Summer *TotalSeason `json:"summer,omitempty"`
}

type AggregateContractYear struct {
	TotalEuros float32 `json:"total_euros,omitempty"`
	TotalKwh float32 `json:"total_kwh,omitempty"`
	ComparativeData map[string]map[string]float32 `json:"comparative,omitempty"`
	AvgPriceKwh float32 `json:"avg_price_kwh,omitempty"`
	TotalConsumePerSeason *TotalPerSeason `json:"total_consume_per_season,omitempty"`
	AvgConsumePerHour map[string]float32 `json:"avg_consume_per_hour,omitempty"`
}

type AggregateContractYearPayload struct {
	JsonPayload AggregateContractYear `json:"payload"`
}

type AggregateContractYearPayloadInterface struct {
	JsonPayload interface{} `json:"payload"`
}

// END yearly aggregates
