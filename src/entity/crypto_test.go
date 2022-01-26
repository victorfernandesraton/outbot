package entity_test

import (
	"testing"
	"time"

	"github.comvictorfernandesraton/outbot/src/entity"
)

func TestValidJSONParse(t *testing.T) {
	data := &entity.Crypto{
		Name:      "Etherium",
		Symbol:    "ETH",
		CreatedAt: time.Date(2010, 4, 21, 0, 0, 0, 0, time.UTC),
		Site:      "www.eth.com",
	}
	result, err := data.ParseJSON()
	if err != nil {
		t.Errorf("Error in parse: %v", err.Error())
	}
	if result != string(`{"name":"Etherium","symbol":"ETH","createdAt":"2010-04-21T00:00:00Z","site":"www.eth.com"}`) {
		t.Errorf("Invalid return %v", result)
	}
}
