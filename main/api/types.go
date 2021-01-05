package api

type Interval string

const (
	Minute Interval = "1m"
	Hour   Interval = "1h"
	Day    Interval = "1D"
)