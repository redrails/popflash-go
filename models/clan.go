package models

type (
	ClanMatchesResponse struct {
		Clan  []*Clan       `json:"clan"`
		Count *CountWrapper `json:"count"`
	}

	Clan struct {
		Id      int32           `json:"id"`
		Name    string          `json:"name"`
		Matches []*MatchWrapper `json:"clan_matches"`
	}
)
