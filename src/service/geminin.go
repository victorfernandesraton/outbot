package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"errors"

	"github.comvictorfernandesraton/outbot/src/config"
	"github.comvictorfernandesraton/outbot/src/entity"
)

type GemininService struct {
	Config *config.GemininConfig
}

func (service *GemininService) GetCryptoInfo(symbol string) (*entity.Crypto, error) {
	return nil, errors.New("Not implemented")
}
func (service *GemininService) GetCurrencyInfo(symbol string) (*entity.Price, error) {

	client := &http.Client{}
	var resDataArray [][]float64
	req, requErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/candles/%susd/1m", service.Config.Doamin, strings.ToLower(symbol)), nil)

	if requErr != nil {
		return nil, requErr
	}
	res, reqDoErr := client.Do(req)
	if reqDoErr != nil {
		fmt.Println(reqDoErr)
		return nil, reqDoErr
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return nil, readErr
	}
	errParse := json.Unmarshal(body, &resDataArray)
	if readErr != nil {
		fmt.Println(errParse)
		return nil, errParse
	}

	return &entity.Price{
		Crypto: &entity.Crypto{
			Symbol: symbol,
		},
		Evaluation: resDataArray[0][5],
		Date: sql.NullTime{
			Valid: true,
			Time:  time.Unix(int64(resDataArray[0][0]), 0),
		},
	}, nil
}
