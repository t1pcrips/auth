package converter_cache

import (
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository/cache/model_cache"
)

func ToSetRepoModelUser(info *model.User) *model_cache.User {
	return &model_cache.User{
		Email:    info.Email,
		Password: info.Password,
	}
}
