package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type (
	Option struct {
		ID          uint   `json:"id"`
		DisplayText string `json:"Display_text"`
	}
	UserPreferencesData struct {
		ID          uint     `json:"-"`
		DisplayName string   `json:"Display_name"`
		DisplayText string   `json:"Display_text,omitempty"`
		Options     []Option `json:"Options"`
	}

	PreferenceRequestBody struct {
		Location struct {
			City      string `json:"city"`
			Latitude  string `json:"lat"`
			Longitude string `json:"long"`
		} `json:"location"`
		UserPreferences map[string]*UserPreferencesData `json:"user_preferences"`
	}
)

func (h *Handler) UserPreferencesGET(ctx *fasthttp.RequestCtx) {
	res, err := h.rDB.GetUserPreferences(ctx.UserValue("user_id"))
	if err != nil {
		fmt.Printf("[ERROR] unable to fetch user pref, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	dataBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("[ERROR] unable to marshal user pref response, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.Write(dataBytes)
}

func (h *Handler) UserPreferencesPOST(ctx *fasthttp.RequestCtx) {
	var (
		userPref PreferenceRequestBody
	)

	err := json.Unmarshal(ctx.Request.Body(), &userPref)
	if err != nil {
		fmt.Printf("[DEBUG] invalid request body, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	optIds := getOptIds(userPref)

	userID := ctx.UserValue("user_id")

	if err = h.wDB.CreateUserPrefrences(userID, optIds); err != nil {
		fmt.Printf("[ERROR] unable to create user preferences, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

}

func (h *Handler) UserPreferencesPUT(ctx *fasthttp.RequestCtx) {

	var userPref PreferenceRequestBody

	err := json.Unmarshal(ctx.Request.Body(), &userPref)
	if err != nil {
		fmt.Printf("[DEBUG] invalid request body, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	optIds := getOptIds(userPref)

	if err = h.wDB.UpdateUserPrefrences(ctx, ctx.UserValue("user_id"), optIds); err != nil {
		fmt.Printf("[ERROR] unable to create user preferences, err : %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
}

func getOptIds(userPref PreferenceRequestBody) (optIds []uint) {
	for _, opts := range userPref.UserPreferences {
		for _, opt := range opts.Options {
			optIds = append(optIds, opt.ID)
		}
	}
	return
}
