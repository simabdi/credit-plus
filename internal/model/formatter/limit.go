package formatter

type LimitFormatter struct {
	Uuid          string `json:"uuid"`
	Tenor         int    `json:"tenor"`
	Amount        int    `json:"amount"`
	CurrentAmount int    `json:"current_amount"`
}
