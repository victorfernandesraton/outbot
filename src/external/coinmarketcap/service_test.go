package coinmarketcap_test

import (
	"fmt"
	"testing"

	"github.comvictorfernandesraton/outbot/src/entity"
	"github.comvictorfernandesraton/outbot/src/external/coinmarketcap"
)

func TestWithGetValue(t *testing.T) {
	var data *entity.Crypto
	service := coinmarketcap.Service{
		Domain: "https://pro-api.coinmarketcap.com/v1",
		Key:    "0e3ce94d-7ef6-44da-83d5-dba193390990",
	}

	res, err := service.GetCryptoInfo("ETH")
	if err != nil {
		t.Error(err.Error())
	}
	if len(res) == 0 {
		t.Error("not find data")
	}
	for _, v := range res {
		if v.Symbol == "ETH" {
			fmt.Println(v.Name)
			data = v
		}
	}
	if data == nil {
		t.Error("not find data")
	}
}
