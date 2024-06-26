package entity

import "time"

type Fixtures struct {
	ID                  int
	Name                string
	StartingAt          time.Time
	ResultInfo          string
	Leg                 string
	Details             interface{}
	Length              int
	Placeholder         bool
	HasOdds             bool
	StartingAtTimestamp int
}
