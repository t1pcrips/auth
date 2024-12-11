package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/t1pcrips/auth/internal/converter"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/utils"
)

func (s *AuthServiceImpl) Login(ctx context.Context, info *model.User) (*model.Tokens, error) {
	keyEmail, err := utils.SecureHash(info.Email)
	if err != nil {
		return nil, err
	}

	keyEmail = "Email: " + keyEmail

	userRedisBytes, err := s.cacheRepository.Get(ctx, keyEmail)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return nil, errs.ErrRedisInternalGet
	}

	if userRedisBytes != nil {
		var userRedis model.User

		err = json.Unmarshal(userRedisBytes, &userRedis)
		if err != nil {
			return nil, errs.ErrFailedUnmarshaling
		}

		err = utils.CheckSecureHash(info.Password, userRedis.Password)
		if err != nil {
			return nil, errs.ErrInvalidPasswords
		}

		return s.generateTokensPair(ctx, &userRedis)
	}

	userPg, err := s.userRepository.GetByParams(ctx, converter.ToParamsByEmail(info.Email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found with these credentials")
		}
		return nil, err
	}

	err = utils.CheckSecureHash(info.Password, userPg.Password)
	if err != nil {
		return nil, errs.ErrInvalidPasswords
	}

	user := &model.User{
		Id:       userPg.Id,
		Email:    userPg.Email,
		Password: userPg.Password,
		Username: userPg.Name,
		Role:     string(userPg.Role),
	}

	err = s.cacheRepository.Set(ctx, keyEmail, user, s.secretsTimeRedisLive)
	if err != nil {
		return nil, errs.ErrSaveRedis
	}

	return s.generateTokensPair(ctx, user)
}

func (s *AuthServiceImpl) generateTokensPair(ctx context.Context, info *model.User) (*model.Tokens, error) {
	accesToken, err := s.jwtService.GenerateAccessToken(info)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(ctx, info)
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		AccessToken:  accesToken,
		RefreshToken: refreshToken,
	}, nil
}

// 1 - при логине мы идем в редис чтобы найти данные пользователя если он там есть то используем его данные роль ник и тд +

// 2 - если нет то идем в бд и там сравниваем также по паролю и емейлу пароль хэш проверка +

// 3 - понять как получать пользователя тип что именно мы должны передать туда +

// 5 - если данные пользователя верны то мы получается можем для него создать токен акссес в котором - юзер ник и роль мб емейл +

// 6 - создать рандомный рефереш токен задать время и сохранить его в бд думаю в редис лучше всего будет +

// 7 - нахуя эта метадата нужна

// redis - сохраняем пользователя я так понял полностью ключ будет хэшированный емейл

// функция сверки хэшированного и не хэшированного при помощи bcrypto
