package middleware

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/config"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/model"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/request"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/response"
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/util/authentication"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	jwt2 "gopkg.in/dgrijalva/jwt-go.v3"
	"sync"
	"time"
)

// singleton mode for jwt.GinJWTMiddleware
var (
	jwtMiddleware *jwt.GinJWTMiddleware
	once          sync.Once
)

func newJWTMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      config.GetConfig().JWT.Realm,
		Key:        []byte(config.GetConfig().JWT.Key),
		Timeout:    config.GetConfig().JWT.Timeout * time.Minute,
		MaxRefresh: config.GetConfig().JWT.MaxRefreshDelay * time.Minute,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"username": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(claims jwt2.MapClaims) interface{} {
			// TODO: get user info from DB
			if username, ok := claims["username"]; ok {
				if username == "admin" {
					return &model.User{
						Username:  username.(string),
						Privilege: authentication.Admin,
					}
				} else {
					return &model.User{
						Username:  username.(string),
						Privilege: authentication.StandardUser,
					}
				}
			}

			return nil
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginRequest := &request.LoginRequest{}
			if err := c.Bind(loginRequest); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			username := loginRequest.Username
			password := loginRequest.Password
			// TODO: perform auth with DB
			if username == "test" && password == "test" ||
				username == "admin" && password == "admin" {
				log.Infof("user %v log in", username)
				return &model.User{
					Username: username,
				}, nil
			} else {
				return nil, jwt.ErrFailedAuthentication
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if user, ok := data.(*model.User); ok {
				return authentication.AuthCheck(c.HandlerName(), user)
			}
			return false
		},
		Unauthorized: func(ctx *gin.Context, code int, msg string) {
			response.NewStandardError(ctx, code, msg)
		},
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func GetGinJWTMiddleware() *jwt.GinJWTMiddleware {
	once.Do(func() {
		jwtMiddleware = newJWTMiddleware()
	})
	return jwtMiddleware
}
