package token

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

const (
	JwtIss     = "ri-im-A7RKqIHDfhBauNLY"
	PrivateKey = "OFzGavy25QSWcAMUlzxqJZGADVwzceSY2aH712PBWsCvnsyiAchQSWRW1jlMeulSyqB8gaQRryoRHtLIQmcmAZfctSHuw7GIkDml"
)

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

func ParseToken(ctx context.Context, token string) (*map[string]string, error) {
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
		if int(exp) < int(time.Now().Unix()) {
			return nil, errors.WithStack(fmt.Errorf("token has expired"))
		}

		extra := claims["extra"].(string)
		var token map[string]string
		err := json.Unmarshal([]byte(extra), &token)
		if err != nil {
			return nil, err
		}
		return &token, nil
	}
	return nil, nil
}
