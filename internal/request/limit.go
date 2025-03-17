package request

type LimitRequest struct {
	Amount int `json:"amount" validate:"required"`
}
