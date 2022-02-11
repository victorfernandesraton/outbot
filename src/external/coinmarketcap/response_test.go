package coinmarketcap_test

import (
	"encoding/json"
	"testing"

	"github.comvictorfernandesraton/outbot/src/external/coinmarketcap"
)

func TestJsonUnarshallResponse(t *testing.T) {
	data := []byte(`{"data": { "BTC": { "id": 1, "name": "Bitcoin", "symbol": "BTC", "category": "coin", "description": "Bitcoin (BTC) is a cryptocurrency . Users are able to generate BTC through the process of mining. Bitcoin has a current supply of 18,952,550. The last known price of Bitcoin is 44,050.71974828 USD and is up 0.53 over the last 24 hours. It is currently trading on 9126 active market(s) with $32,791,328,911.49 traded over the last 24 hours. More information can be found at https://bitcoin.org/.", "slug": "bitcoin", "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png", "subreddit": "bitcoin", "notice": "", "tags": [ "mineable", "pow", "sha-256", "store-of-value", "state-channel", "coinbase-ventures-portfolio", "three-arrows-capital-portfolio", "polychain-capital-portfolio", "binance-labs-portfolio", "blockchain-capital-portfolio", "boostvc-portfolio", "cms-holdings-portfolio", "dcg-portfolio", "dragonfly-capital-portfolio", "electric-capital-portfolio", "fabric-ventures-portfolio", "framework-ventures-portfolio", "galaxy-digital-portfolio", "huobi-capital-portfolio", "alameda-research-portfolio", "a16z-portfolio", "1confirmation-portfolio", "winklevoss-capital-portfolio", "usv-portfolio", "placeholder-ventures-portfolio", "pantera-capital-portfolio", "multicoin-capital-portfolio", "paradigm-portfolio" ], "tag-names": [ "Mineable", "PoW", "SHA-256", "Store Of Value", "State Channel", "Coinbase Ventures Portfolio", "Three Arrows Capital Portfolio", "Polychain Capital Portfolio", "Binance Labs Portfolio", "Blockchain Capital Portfolio", "BoostVC Portfolio", "CMS Holdings Portfolio", "DCG Portfolio", "DragonFly Capital Portfolio", "Electric Capital Portfolio", "Fabric Ventures Portfolio", "Framework Ventures Portfolio", "Galaxy Digital Portfolio", "Huobi Capital Portfolio", "Alameda Research Portfolio", "a16z Portfolio", "1Confirmation Portfolio", "Winklevoss Capital Portfolio", "USV Portfolio", "Placeholder Ventures Portfolio", "Pantera Capital Portfolio", "Multicoin Capital Portfolio", "Paradigm Portfolio" ], "tag-groups": [ "OTHERS", "ALGORITHM", "ALGORITHM", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY", "CATEGORY" ], "urls": { "website": [ "https://bitcoin.org/" ], "twitter": [], "message_board": [ "https://bitcointalk.org" ], "chat": [], "facebook": [], "explorer": [ "https://blockchain.coinmarketcap.com/chain/bitcoin", "https://blockchain.info/", "https://live.blockcypher.com/btc/", "https://blockchair.com/bitcoin", "https://explorer.viabtc.com/btc" ], "reddit": [ "https://reddit.com/r/bitcoin" ], "technical_doc": [ "https://bitcoin.org/bitcoin.pdf" ], "source_code": [ "https://github.com/bitcoin/bitcoin" ], "announcement": [] }, "platform": null, "date_added": "2013-04-28T00:00:00.000Z", "twitter_username": "", "is_hidden": 0, "date_launched": null, "contract_address": [], "self_reported_circulating_supply": null, "self_reported_tags": null, "self_reported_market_cap": null } } }`)

	res := new(coinmarketcap.ResponseJSON)
	if err := json.Unmarshal(data, &res); err != nil {
		t.Errorf("not valid parse %s", err.Error())
	}
	if len(res.Coin) < 1 {
		t.Error("Invalid data")
	}
	if res.Coin["BTC"].Symbol != "BTC" {
		t.Error("Invalid data")
	}

}
