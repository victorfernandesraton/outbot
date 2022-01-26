package entity

import (
	"database/sql"
)

type Price struct {
	Crypto     Crypto
	Evaluation float64
	Date       sql.NullTime
}
