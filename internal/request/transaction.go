package request

type (
	TransactionRequest struct {
		Otr       int    `json:"otr" validate:"required"`
		UuidLimit string `json:"uuid_limit" validate:"required,max=100"`
		AssetName string `json:"asset_name" validate:"required,max=100"`
		Platform  string `json:"platform" validate:"required,max=100"`
	}

	VerifyOtpRequest struct {
		Otp string `json:"otp" validate:"required,min=6,max=6"`
	}
)
