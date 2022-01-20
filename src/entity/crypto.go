package entity

import (
	"net/url"
	"time"
)

type Crypto struct {
	Name      string
	Symbol    string
	CreatedAt time.Time
	Site      url.URL
}
