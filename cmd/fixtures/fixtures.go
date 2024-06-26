package fixtures

type Fixture struct {
	ID                  int         `json:"id"`
	SportID             int         `json:"sport_id"`
	LeagueID            int         `json:"league_id"`
	SeasonID            int         `json:"season_id"`
	StageID             int         `json:"stage_id"`
	GroupID             interface{} `json:"group_id"`
	AggregateID         interface{} `json:"aggregate_id"`
	RoundID             int         `json:"round_id"`
	StateID             int         `json:"state_id"`
	VenueID             int         `json:"venue_id"`
	Name                string      `json:"name"`
	StartingAt          string      `json:"starting_at"`
	ResultInfo          string      `json:"result_info"`
	Leg                 string      `json:"leg"`
	Details             interface{} `json:"details"`
	Length              int         `json:"length"`
	Placeholder         bool        `json:"placeholder"`
	HasOdds             bool        `json:"has_odds"`
	HasPremiumOdds      bool        `json:"has_premium_odds"`
	StartingAtTimestamp int         `json:"starting_at_timestamp"`
}

// Response represents the full response structure.
type Response struct {
	Data []Fixture `json:"data"`
}
