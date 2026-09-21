package models

type Tick struct {
	Timestamp string
	Bid       string
	Ask       string
}

func (t Tick) String() string {
	return t.Timestamp + "," + t.Bid + "," + t.Ask
}
