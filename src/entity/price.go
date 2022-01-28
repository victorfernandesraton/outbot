package entity

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Price struct {
	Crypto     Crypto
	Evaluation float64
	Date       sql.NullTime
}

func (p *Price) MarshalJSON() ([]byte, error) {
	var dataStr interface{}
	if p.Crypto.CreatedAt.Valid == true {
		dataStr = struct {
			Crypto     Crypto    `json:"crypto"`
			Evaluation float64   `json:"evaluation"`
			Date       time.Time `json:"date"`
		}{
			Crypto:     p.Crypto,
			Evaluation: p.Evaluation,
			Date:       p.Date.Time,
		}
	} else {
		dataStr = struct {
			Crypto     Crypto  `json:"crypto"`
			Evaluation float64 `json:"evaluation"`
		}{
			Crypto:     p.Crypto,
			Evaluation: p.Evaluation,
		}
	}
	data, err := json.Marshal(dataStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("JSON Invalid parse %v", err.Error()))
	}
	return data, nil
}
