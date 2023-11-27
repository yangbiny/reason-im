package token

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

const (
	JwtIss     = "ri-im-A7RKqIHDfhBauNLY"
	PrivateKey = "OFzGavy25QSWcAMUlzxqJZGADVwzceSY2aH712PBWsCvnsyiAchQSWRW1jlMeulSyqB8gaQRryoRHtLIQmcmAZfctSHuw7GIkDml"
)

type Claims struct {
	claims map[string]interface{}
}

func (claims *Claims) UserId() int64 {
	return claims.claims["userId"].(int64)
}

func (claims *Claims) KeyAsString(key string) string {
	return claims.claims[key].(string)
}

func (claims *Claims) HasExpire() bool {
	expireAt := claims.claims["expire"].(int)
	return expireAt < int(time.Now().Unix())
}

func ApplyToken(ctx context.Context, userId int64, claims map[string]string) (*string, error) {
	if len(claims) == 0 {
		return nil, fmt.Errorf("claims is empty")
	}
	marshal, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"extra":   string(marshal),
		"Subject": userId,
		"iss":     JwtIss,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	by := []byte(PrivateKey)
	signedString, err := token.SignedString(by)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &signedString, nil
}

func ParseToken(ctx context.Context, token string) (*Claims, error) {
	if len(token) == 0 {
		return nil, errors.WithStack(fmt.Errorf("token is empty"))
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(PrivateKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		exp := claims["exp"].(float64)
		extra := claims["extra"].(string)
		var token map[string]interface{}
		err := json.Unmarshal([]byte(extra), &token)
		if err != nil {
			return nil, err
		}
		token["userId"], err = strconv.Atoi(claims["Subject"].(string))
		token["expire"] = int(exp)
		return &Claims{
			claims: token,
		}, nil
	}
	return nil, nil
}
