package coinmarketcap

type ResponseJSON struct {
	Coin map[string]*CryptoInfo `json:"data"`
}

type CryptoInfo struct {
	Category                      string        `json:"category"`
	ContractAddress               []interface{} `json:"contract_address"`
	DateAdded                     string        `json:"date_added"`
	DateLaunched                  string        `json:"date_launched"`
	Description                   string        `json:"description"`
	ID                            int64         `json:"id"`
	IsHidden                      int64         `json:"is_hidden"`
	Logo                          string        `json:"logo"`
	Name                          string        `json:"name"`
	Notice                        string        `json:"notice"`
	Platform                      interface{}   `json:"platform"`
	SelfReportedCirculatingSupply interface{}   `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         interface{}   `json:"self_reported_market_cap"`
	SelfReportedTags              interface{}   `json:"self_reported_tags"`
	Slug                          string        `json:"slug"`
	Subreddit                     string        `json:"subreddit"`
	Symbol                        string        `json:"symbol"`
	Tag_groups                    []string      `json:"tag-groups"`
	Tag_names                     []string      `json:"tag-names"`
	Tags                          []string      `json:"tags"`
	TwitterUsername               string        `json:"twitter_username"`
	Urls                          struct {
		Announcement []interface{} `json:"announcement"`
		Chat         []interface{} `json:"chat"`
		Explorer     []string      `json:"explorer"`
		Facebook     []interface{} `json:"facebook"`
		MessageBoard []string      `json:"message_board"`
		Reddit       []string      `json:"reddit"`
		SourceCode   []string      `json:"source_code"`
		TechnicalDoc []string      `json:"technical_doc"`
		Twitter      []interface{} `json:"twitter"`
		Website      []string      `json:"website"`
	} `json:"urls"`
}
