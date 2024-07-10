package entity

type Fixtures struct {
	ID                  int
	Name                string
	StartingAt          string
	ResultInfo          string
	Leg                 string
	Details             interface{}
	Length              int
	Placeholder         bool
	HasOdds             bool
	StartingAtTimestamp int
}
