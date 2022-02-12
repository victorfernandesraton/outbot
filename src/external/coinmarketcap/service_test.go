//+build e2e

package coinmarketcap_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.comvictorfernandesraton/outbot/src/config"
	"github.comvictorfernandesraton/outbot/src/entity"
	"github.comvictorfernandesraton/outbot/src/external/coinmarketcap"
)

func init() {
	// Root folder of this project
	if err := godotenv.Load(filepath.Join(config.ProjectRootPath, ".env")); err != nil {
		panic(err.Error())
	}
}

func TestWithGetValue(t *testing.T) {
	var data *entity.Crypto
	service := coinmarketcap.Service{
		Domain: os.Getenv("COINMARKETCAP_DOMAIN"),
		Key:    os.Getenv("COINMARKETCAP_KEY"),
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
			data = v
		}
	}
	if data == nil {
		t.Error("not find data")
	}
}
