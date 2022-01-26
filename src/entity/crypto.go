package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Crypto struct {
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"createdAt"`
	Site      string    `json:"site"`
}

func (c *Crypto) ParseJSON() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", errors.New(fmt.Sprintf("JSON Error parse, %v", err.Error()))
	}

	return string(data), nil
}
