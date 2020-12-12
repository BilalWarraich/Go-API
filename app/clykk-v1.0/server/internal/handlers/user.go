package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/clykk-user-atif/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

func (h *Handler) Login(ctx *fasthttp.RequestCtx) {

	var user db.User
	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		fmt.Printf("[DEBUG] invalid request body, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	u, err := h.rDB.GetUser(user.Email, user.Password)
	if err != nil {
		if _, ok := err.(*db.UserNotFoundErr); ok {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Printf("[ERROR] unable to sign token string, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	if _, err = ctx.WriteString(tokenStr); err != nil {
		fmt.Printf("[ERROR] unable to write response, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

}
