package models

import "time"

type (
	MatchResponse struct {
		Match *MatchInstance `json:"match"`
	}

	MatchInstance struct {
		Match
		Leader            int64        `json:"leader"`
		DemoUploaded      bool         `json:"demo_uploaded"`
		Team1NameOverride string       `json:"team1_name_override"`
		Team2NameOverride string       `json:"team2_name_override"`
		Matches           []*UserMatch `json:"users_matches"`
	}

	UserMatch struct {
		User              *MatchUser `json:"user"`
		Private           bool       `json:"private"`
		UserId            int32      `json:"user_id"`
		Team              int32      `json:"team"`
		Kills             int32      `json:"kills"`
		Assists           int32      `json:"assists"`
		Deaths            int32      `json:"deaths"`
		HSRatio           float64    `json:"headshot_ratio"`
		AvgDamagePerRound float32    `json:"average_damage_per_round"`
		BombsPlanted      int32      `json:"bomb_planted"`
		BombsDefused      int32      `json:"bomb_defused"`
		RoundsPlayed      int32      `json:"rounds_played"`
		ClutchKills       int32      `json:"clutch_kills"`
		IsCoach           bool       `json:"is_coach"`
		HltvRating        float64    `json:"hltv_rating"`
		Premium           *Premium   `json:"premium"`
	}

	MatchUser struct {
		Name              string    `json:"name"`
		SubscriptionLevel int32     `json:"subscription_level"`
		Subscription      time.Time `json:"subscription"`
		Avatar            string    `json:"avatar"`
	}

	Premium struct {
		UserId                int64   `json:"user_id"`
		T1V1Clutches          int32   `json:"total_1v1_clutches"`
		T1V2Clutches          int32   `json:"total_1v2_clutchesme"`
		T1V3Clutches          int32   `json:"total_1v3_clutches"`
		T1V4Clutches          int32   `json:"total_1v4_clutches"`
		T1V5Clutches          int32   `json:"total_1v5_clutches"`
		T1kRounds             int32   `json:"total_1k_rounds"`
		T2kRounds             int32   `json:"total_2k_rounds"`
		T3kRounds             int32   `json:"total_3k_rounds"`
		T4kRounds             int32   `json:"total_4k_rounds"`
		T5kRounds             int32   `json:"total_5k_rounds"`
		HltvRating            float32 `json:"hltv_rating"`
		FlashAssists          int32   `json:"flash_assists"`
		FlashThrows           int32   `json:"flash_throws"`
		FlashEnemyHits        int32   `json:"flash_enemy_hits"`
		FlashEnemyDuration    float32 `json:"flash_enemy_duration"`
		FlashEnemyPercent     float32 `json:"flash_enemy_percent"`
		FlashTeammateHits     int32   `json:"flash_teammate_hits"`
		FlashTeammatePercent  float32 `json:"flash_teammate_percent"`
		FlashTeammateDuration float32 `json:"flash_teammate_duration"`
		Damage                int32   `json:"damage_done"`
		RoundsFired           int32   `json:"rounds_fired"`
		RoundsHit             int32   `json:"rounds_hit"`
		RoundsHitHead         int32   `json:"rounds_hit_head"`
		RoundsHitUpperTorso   int32   `json:"rounds_hit_upper_torso"`
		RoundsHitLowerTorso   int32   `json:"rounds_hit_lower_torso"`
		RoundsHitLeftArm      int32   `json:"rounds_hit_left_arm"`
		RoundsHitRightArm     int32   `json:"rounds_hit_right_arm"`
		RoundsHitLeftLeg      int32   `json:"rounds_hit_left_leg"`
		RoundsHitRightLeg     int32   `json:"rounds_hit_right_leg"`
		IsSubscribed          bool    `json:"is_subscribed"`
		FirstKills            int32   `json:"first_kills"`
	}
)
