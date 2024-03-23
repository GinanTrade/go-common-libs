package middlewares

import (
	"errors"
	"strings"
	"time"

	"github.com/GinanTrade/go-commmon-libs/config"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenHeader string = c.GetHeader("Authorization")

		if tokenHeader == "" {

			c.AbortWithStatusJSON(401, authentication.ResponseWebFailed{
				Type:    "UNAUTHORIZED",
				Message: "invalid token provided",
				Status:  "failed",
			})
			return
		}

		var tokenArray = strings.Split(tokenHeader, " ")

		if len(tokenArray) < 2 {
			c.AbortWithStatusJSON(401, authentication.ResponseWebFailed{
				Type:    "UNAUTHORIZED",
				Message: "invalid token format",
				Status:  "failed",
			})
			return
		}

		var isValidBearer bool = tokenArray[0] == "Bearer"

		if !isValidBearer {
			c.AbortWithStatusJSON(401, authentication.ResponseWebFailed{
				Type:    "UNAUTHORIZED",
				Message: "Invalid bearer token",
				Status:  "failed",
			})
			return
		}
		tokenHeader = tokenArray[1]
		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				panic(model.UnauthorizedError{Message: "unexpected signing method"})
			}

			return []byte(config.Config.Get("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, model.ResponseWebFailed{
				Type:    "UNAUTHORIZED",
				Message: err.Error(),
				Status:  "failed",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			panic(errors.New("INTERNAL SERVER ERROR"))
		}

		var data model.Auth = model.Auth{
			LastName:  claims["last_name"].(string),
			FirstName: claims["first_name"].(string),
			Iat:       claims["iat"].(float64),
			Exp:       claims["exp"].(float64),
			AccountId: claims["account_id"].(string),
			Email:     claims["email"].(string),
			Jti:       claims["jti"].(string),
			TokenType: claims["token_type"].(string),
			CompanyId: claims["company_id"].(float64),
		}

		providedTimeStamp := int64(data.Exp)

		now := time.Now()

		providedTime := time.Unix(providedTimeStamp, 0)

		if providedTime.Before(now) {
			c.AbortWithStatusJSON(401, model.ResponseWebFailed{
				Type:    "UNAUTHORIZED",
				Message: "tokenis expired",
				Status:  "failed",
			})
			return
		}

		c.Set("user", data)
		c.Next()
	}
}
