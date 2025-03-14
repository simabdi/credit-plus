package resource

import (
	"credit-plus/internal/model/entity"
	"credit-plus/internal/model/formatter"
)

func CheckAccountResource(user entity.User) formatter.CheckAccountFormatter {
	Resource := formatter.CheckAccountFormatter{
		Uuid: user.Uuid,
	}

	return Resource
}

func LoginResource(user entity.User, token string) formatter.LoginFormatter {
	Resource := formatter.LoginFormatter{
		Uuid:  user.Uuid,
		Token: token,
	}

	return Resource
}
