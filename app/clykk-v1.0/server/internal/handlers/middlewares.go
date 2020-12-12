package handlers

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

const (
	jwtSecret = "secret"
)

func AuthUser(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		headers := make(map[string]string)

		ctx.Request.Header.VisitAll(func(key, value []byte) {
			headers[string(key)] = string(value)
		})

		authHeader, ok := headers["Authorization"]
		if !ok {
			ctx.WriteString(unAuthorizedAccessErr)
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.WriteString(unAuthorizedAccessErr)
			return
		}

		token, err := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.WriteString(unAuthorizedAccessErr)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.WriteString(unAuthorizedAccessErr)
			return
		}

		ctx.SetUserValue("user_id", claims["user_id"])
		next(ctx)
		return
	}

}
