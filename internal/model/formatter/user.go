package formatter

type (
	CheckAccountFormatter struct {
		Uuid string `json:"uuid"`
	}

	LoginFormatter struct {
		Uuid  string `json:"uuid"`
		Token string `json:"token"`
	}
)
