package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"reason-im/internal/utils/caller"
	"reason-im/internal/utils/logger"
	"time"
)

const JwtIss = "ri-im-n98TmvynRdEl29Ko"

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("token")
		if err != nil {
			ctx.JSON(400, caller.ApiResp{Code: caller.ParamInvalid, Msg: "未登录"})
			ctx.Abort()
			return
		}
		token := cookie.Value
		if len(token) == 0 {
			ctx.Abort()
			ctx.JSON(400, caller.ApiResp{Code: caller.ParamInvalid, Msg: "未登录"})
			return
		}
		jwtToken, err := ParseJwtToken(token)
		if err != nil {
			logger.ErrorWithErr(ctx, "parse jwt token has failed", errors.WithStack(err))
			ctx.Abort()
			ctx.JSON(400, caller.ApiResp{Code: caller.ParamInvalid, Msg: "token已过期"})
			return
		}
		if jwtToken == nil {
			ctx.Abort()
			ctx.JSON(400, caller.ApiResp{Code: caller.ParamInvalid, Msg: "token失效，请重新登录"})
			return
		}
		ctx.Set("login_user_id", jwtToken.UserId)
		ctx.Next()
	}
}

func ParseJwtToken(token string) (*Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		extra := claims["extra"].(string)
		var token Token
		err := json.Unmarshal([]byte(extra), &token)
		if err != nil {
			return nil, err
		}
		return &token, nil
	}
	return nil, nil
}

func GenerateJwtToken(ctx *gin.Context, username string, userId int64) error {
	tokenStruct := &Token{
		Username: username,
		UserId:   userId,
		SignAt:   time.Now().Unix(),
	}
	marshal, err := json.Marshal(tokenStruct)
	if err != nil {
		logger.ErrorWithErr(ctx, "marshal token has failed", errors.WithStack(err))
		return nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"extra": string(marshal),
		"iss":   JwtIss,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	by := []byte(privateKey)
	signedString, err := token.SignedString(by)
	if err != nil {
		logger.ErrorWithErr(ctx, "sign token has failed", errors.WithStack(err))
		return nil
	}
	ctx.SetCookie("token", signedString, 60*60*24*7, "/", "localhost", false, true)
	return nil
}

const privateKey = "ZQaWJ0cMYFEx71hAmFcc8wr9Qfjky98kZtSrZzESiSN86fKjjNkaWeDkkpG3wzkBTAY4FxisenWxbr1ZiXcGW9TQnabfxmQzT45h"

type Token struct {
	Username string `json:"username"`
	UserId   int64  `json:"user_id"`
	SignAt   int64  `json:"sign_at"`
}
