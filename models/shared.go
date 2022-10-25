package models

import "time"

type (
	Match struct {
		Id         int32     `json:"id,omitempty"`
		Winner     int32     `json:"winner"`
		Map        string    `json:"map"`
		Status     int32     `json:"status,omitempty"`
		Score1     int32     `json:"score1"`
		Score2     int32     `json:"score2"`
		Datacenter int32     `json:"datacenter"`
		Type       int32     `json:"type"`
		CreatedAt  time.Time `json:"created_at"`
	}

	MatchWrapper struct {
		Match *Match `json:"match"`
	}

	CountWrapper struct {
		Aggregate *Count `json:"aggregate"`
	}

	Count struct {
		Count int32 `json:"count"`
	}
)
