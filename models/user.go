package models

import "time"

type (
	UserMatchesResponse struct {
		User  *User         `json:"user"`
		Count *CountWrapper `json:"count"`
	}

	User struct {
		Id       int32          `json:"id"`
		Name     string         `json:"name"`
		Avatar   string         `json:"avatar"`
		LastSeen time.Time      `json:"last_seen"`
		Matches  []*UserMatches `json:"users_matches"`
	}

	UserMatches struct {
		MatchId           int32   `json:"match_id"`
		UserId            int32   `json:"user_id"`
		Team              int32   `json:"team"`
		AvgDamagePerRound float32 `json:"average_damage_per_round"`
		Kills             int32   `json:"kills"`
		Deaths            int32   `json:"deaths"`
		HSRatio           float32 `json:"headshot_ratio"`
		RoundsPlayed      int32   `json:"rounds_played"`
		HltvRating        float32 `json:"hltv_rating"`
		Match             *Match  `json:"match"`
	}
)
