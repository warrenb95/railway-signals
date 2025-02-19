package domain

type Signal struct {
	ID   int     `json:"signal_id"`
	Name *string `json:"signal_name"`
	ELR  string  `json:"elr"`
}

type Mileage struct {
	SignalID int      `json:"signal_id"`
	TrackID  int      `json:"track_id"`
	Mileage  *float64 `json:"mileage"`
}

type Track struct {
	ID     int    `json:"track_id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type TrackSignals struct {
	ID      int    `json:"track_id"`
	Source  string `json:"source"`
	Target  string `json:"target"`
	Signals []struct {
		ID      int      `json:"signal_id"`
		Name    *string  `json:"signal_name"`
		ELR     string   `json:"elr"`
		Mileage *float64 `json:"mileage"`
	}
}
