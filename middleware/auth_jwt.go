package middleware

import (
	"log"
	"net/http"
	"strings"

	"heroku-backend-a-cocreate/helper/header"
	"heroku-backend-a-cocreate/helper/response"
	"heroku-backend-a-cocreate/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := header.GetAuthorization(c)
		if authHeader == "" {
			response.BuildErrResponse(c, http.StatusNotFound, "Failed to process request", "No token found")
			return
		}

		bearer := strings.HasPrefix(authHeader, "Bearer")
		if !bearer {
			response.BuildErrResponse(c, http.StatusUnauthorized, "Failed to process request", "Bearer token rules")
			return
		}

		stringToken := strings.Split(authHeader, " ")[1]

		token, err := service.ValidateToken(stringToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id]:", claims["user_id"])
		} else {
			log.Println(err)
			response.BuildErrResponse(c, http.StatusUnauthorized, "Token is not valid", err.Error())
			return
		}

		c.Set("userData", token.Claims.(jwt.MapClaims))
		c.Next()
	}
}
