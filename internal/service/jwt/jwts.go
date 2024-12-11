package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/t1pcrips/auth/internal/config"
	"github.com/t1pcrips/auth/internal/errs"
	"github.com/t1pcrips/auth/internal/model"
	"github.com/t1pcrips/auth/internal/repository"
	"github.com/t1pcrips/auth/internal/service"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

const (
	id            = "id"
	email         = "email"
	username      = "username"
	role          = "role"
	iat           = "iat"
	exp           = "exp"
	authorization = "authorization"
	authPrefix    = "Bearer "
	accessToken   = "access"
	refreshToken  = "refresh"
)

type JWTServiceImpl struct {
	cacheRepository repository.CacheRepository
	secretConfig    *config.SecretsConfig
}

func NewJWTServiceImpl(secretConfig *config.SecretsConfig, cacheRepository repository.CacheRepository) service.JWTService {
	return &JWTServiceImpl{
		cacheRepository: cacheRepository,
		secretConfig:    secretConfig,
	}
}

func (s *JWTServiceImpl) GenerateAccessToken(user *model.User) (string, error) {
	return s.generateToken(user, accessToken)
}

func (s *JWTServiceImpl) GenerateRefreshToken(ctx context.Context, user *model.User) (string, error) {
	token, err := s.generateToken(user, refreshToken)
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("refresh_token: %d", user.Id)
	err = s.cacheRepository.Set(ctx, key, &model.RefreshToken{Token: token}, s.secretConfig.JWTRefreshTime)
	if err != nil {
		return "", errs.ErrSaveRedis
	}

	return token, nil
}

func (s *JWTServiceImpl) ValidateRefreshToken(ctx context.Context, signedToken string) (*model.User, error) {
	user, err := s.verifyToken(signedToken)
	if err != nil {
		return nil, err
	}
	// проверили валидность сущетсвует вообще такой или нет и нашли id user из токена
	key := fmt.Sprintf("refresh_token: %d", user.Id)

	tokenRedisBytes, err := s.cacheRepository.Get(ctx, key)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return nil, errs.ErrRedisInternalGet
	}

	var refreshToken model.RefreshToken

	err = json.Unmarshal(tokenRedisBytes, &refreshToken)
	if err != nil {
		return nil, errs.ErrFailedUnmarshaling
	}

	if signedToken == refreshToken.Token {
		return user, nil
	}

	return user, nil
}

func (s *JWTServiceImpl) ValidateTokenFromMeatadata(ctx context.Context) (*model.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md[authorization]
	if !ok || len(authHeader[0]) == 0 {
		return nil, errors.New("authorization header is not provide")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	user, err := s.verifyToken(accessToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *JWTServiceImpl) verifyToken(signedToken string) (*model.User, error) {
	token, err := jwt.Parse(signedToken, func(signedToken *jwt.Token) (interface{}, error) {
		return []byte(s.secretConfig.JWTRefresh), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errs.ErrInvalidToken
	}

	claims, err := s.userClaims(token.Claims.(jwt.MapClaims))
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *JWTServiceImpl) userClaims(claims jwt.MapClaims) (*model.User, error) {
	exp, err := getClaimString[float64](claims, exp)
	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > int64(exp) {
		return nil, errs.ErrInvalidToken
	}

	idInt, err := getClaimString[float64](claims, id)
	if err != nil {
		return nil, err
	}

	nameStr, err := getClaimString[string](claims, username)
	if err != nil {
		return nil, err
	}

	emailStr, err := getClaimString[string](claims, email)
	if err != nil {
		return nil, err
	}

	roleStr, err := getClaimString[string](claims, role)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:       int64(idInt),
		Username: nameStr,
		Email:    emailStr,
		Role:     roleStr,
	}, nil
}

func getClaimString[T any](claims jwt.MapClaims, key string) (T, error) {
	var nils T

	value, ok := claims[key]
	if !ok {
		return nils, errors.New("error with token Claims")
	}

	newValue, ok := value.(T)
	if !ok {
		return nils, errors.New("interface conversion error to type t")
	}

	return newValue, nil
}

func (s *JWTServiceImpl) generateToken(user *model.User, signed string) (string, error) {
	claims := jwt.MapClaims{
		id:       user.Id,
		email:    user.Email,
		username: user.Username,
		role:     user.Role,
		iat:      time.Now().Unix(),
		exp:      time.Now().Add(time.Duration(s.secretConfig.JWTAccessTime) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var signedToken string
	var err error

	switch signed {
	case "refresh":
		signedToken, err = token.SignedString([]byte(s.secretConfig.JWTRefresh))
		if err != nil {
			return "", errors.New("failed to signed token")
		}
	case "access":
		signedToken, err = token.SignedString([]byte(s.secretConfig.JWTAccess))
		if err != nil {
			return "", errors.New("failed to signed token")
		}
	}

	return signedToken, nil
}

// 1 - при логине мы идем в редис чтобы найти данные пользователя если он там есть то используем его данные роль ник и тд
// 2 - если нет то идем в бд и там сравниваем также по паролю и емейлу пароль хэш проверка
// 3 - понять как получать пользователя тип что именно мы должны передать туда
// 4 - разобраться как мы используем метадату и для чего она это раз потом
// 5 - если данные пользователя верны то мы получается можем для него создать токен акссес в котором - юзер ник и роль мб емейл
// 6 - создать рандомный рефереш токен задать время и сохранить его в бд думаю в редис лучше всего будет
// 7 - нахуя эта метадата нужна
