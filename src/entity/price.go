package entity

import "time"

type Price struct {
	Crypto     Crypto
	Evaluation float64
	Date       time.Time
}
