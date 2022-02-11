package coinmarketcap

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.comvictorfernandesraton/outbot/src/entity"
)

type Service struct {
	Domain string
	Key    string
}

func (s *Service) GetCryptoInfo(symbol string) ([]*entity.Crypto, error) {

	var response ResponseJSON
	var result []*entity.Crypto
	url := fmt.Sprintf("%s/cryptocurrency/info?symbol=%s", s.Domain, symbol)

	req, errBuildRequest := http.NewRequest("GET", url, nil)

	if errBuildRequest != nil {
		return nil, errors.New("Error on build request")
	}

	req.Header.Add("X-CMC_PRO_API_KEY", s.Key)

	res, errHandleRequest := http.DefaultClient.Do(req)

	if errHandleRequest != nil {
		return nil, errors.New("Error on send request")
	}

	defer res.Body.Close()

	body, errBuildResponse := ioutil.ReadAll(res.Body)

	if errBuildResponse != nil {
		return nil, errors.New("Error on build buff")
	}

	errParseResponse := json.Unmarshal(body, &response)
	if errParseResponse != nil {
		return nil, errors.New("Error on parse response data")
	}

	for _, v := range response.Coin {
		createdAt, errParsedTime := time.Parse("2013-04-28T00:00:00.000Z", v.DateLaunched)
		coin := entity.Crypto{
			Name:   v.Name,
			Symbol: v.Symbol,
			Site:   v.Urls.Website[0],
		}
		if errParsedTime != nil {
			coin.CreatedAt = sql.NullTime{
				Time:  createdAt,
				Valid: true,
			}
		}
		result = append(result, &coin)
	}

	return result, nil
}
