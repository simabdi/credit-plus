package request

type (
	LoginRequest struct {
		PhoneNumber string `json:"phone_number" validate:"required"`
	}

	VerifyPinRequest struct {
		Uuid string `json:"uuid" validate:"required,max=100"`
		Pin  string `json:"pin" validate:"required,min=6"`
	}
)

type RegistrationRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type SetPinRequest struct {
	Pin string `json:"pin" validate:"required,min=6,max=6"`
}
