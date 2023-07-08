package middleware

import (
	"ginchat/common"
	"ginchat/common/response"
	"ginchat/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(service.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.App.Config.Jwt.Secret), nil
		})
		if err != nil || service.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < common.App.Config.Jwt.RefreshGracePeriod {
			lock := common.Lock("refresh_token_lock", common.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := service.JwtService.GetUserInfo(GuardName, claims.Id)
				if err != nil {
					common.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := service.JwtService.CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = service.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
