package entity_test

import (
	"database/sql"
	"testing"
	"time"

	"github.comvictorfernandesraton/outbot/src/entity"
)

func TestValidJSONParse(t *testing.T) {
	data := &entity.Crypto{
		Name:   "Etherium",
		Symbol: "ETH",
		CreatedAt: sql.NullTime{
			Time:  time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
		Site: "www.eth.com",
	}
	result, err := data.ParseJSON()
	if err != nil {
		t.Errorf("Error in parse: %v", err.Error())
	}
	if result != string(`{"name":"Etherium","symbol":"ETH","createdAt":"2019-04-21T00:00:00Z","site":"www.eth.com"}`) {
		t.Errorf("Invalid return %v", result)
	}
}

func TestWithNullDate(t *testing.T) {
	data := &entity.Crypto{
		Name:   "Etherium",
		Symbol: "ETH",
		CreatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		Site: "www.eth.com",
	}
	result, err := data.ParseJSON()
	if err != nil {
		t.Errorf("Error in parse: %v", err.Error())
	}
	if result != string(`{"name":"Etherium","symbol":"ETH","site":"www.eth.com"}`) {
		t.Errorf("Invalid return %v", result)
	}
}

func TestWIthJSONNotHaveSite(t *testing.T) {
	data := &entity.Crypto{
		Name:   "Etherium",
		Symbol: "ETH",
		CreatedAt: sql.NullTime{
			Time:  time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}
	res, err := data.ParseJSON()
	if err != nil {
		t.Errorf("Error in parse %v", err.Error())
	}
	if res != string(`{"name":"Etherium","symbol":"ETH","createdAt":"2019-04-21T00:00:00Z"}`) {
		t.Errorf("Unexpected return, %v got %v", `{"name":"Etherium","symbol":"ETH","createdAt":"2019-04-21T00:00:00Z"}`, res)
	}

}
