package entity_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.comvictorfernandesraton/outbot/src/entity"
)

func TestValidParser(t *testing.T) {
	crypto := &entity.Crypto{
		Name:   "Etherium",
		Symbol: "ETH",
		CreatedAt: sql.NullTime{
			Time:  time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
		Site: "www.eth.com",
	}
	data := &entity.Price{
		Crypto:     crypto,
		Evaluation: 20,
		Date: sql.NullTime{
			Time:  time.Date(2020, 4, 10, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}
	_, err := json.Marshal(data)
	if err != nil {
		t.Error("Invalid parse")
	}

}
