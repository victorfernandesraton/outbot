package entity

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Crypto struct {
	Name      string       `json:"name"`
	Symbol    string       `json:"symbol"`
	CreatedAt sql.NullTime `json:"createdAt"`
	Site      string       `json:"site"`
}

func (c *Crypto) ParseJSON() (string, error) {
	var dataStr interface{}
	if c.CreatedAt.Valid {
		dataStr = struct {
			Name      string    `json:"name"`
			Symbol    string    `json:"symbol"`
			CreatedAt time.Time `json:"createdAt,omitempty"`
			Site      string    `json:"site,omitempty"`
		}{
			Name:      c.Name,
			Symbol:    c.Symbol,
			Site:      c.Site,
			CreatedAt: c.CreatedAt.Time,
		}
	} else {
		dataStr = struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			Site   string `json:"site"`
		}{
			Name:   c.Name,
			Symbol: c.Symbol,
			Site:   c.Site,
		}
	}
	data, err := json.Marshal(dataStr)
	if err != nil {
		return "", errors.New(fmt.Sprintf("JSON Error parse, %v", err.Error()))
	}
	return string(data), nil

}
