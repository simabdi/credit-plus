package formatter

type LimitFormatter struct {
	Uuid   string `json:"uuid"`
	Tenor  int    `json:"tenor"`
	Amount int    `json:"amount"`
}
