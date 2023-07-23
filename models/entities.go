package models

type Bet struct {
	ID        int64   `json:"id"`
	Event     string  `json:"event"`
	Selection string  `json:"selection"`
	Stake     float64 `json:"stake"`
	Odds      float64 `json:"odds"`
	Status    string  `json:"status"`
}
