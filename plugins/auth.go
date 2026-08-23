package plugins

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/db"
	"github.com/levisantosp/altamira-participa/ent/generated"
	"github.com/levisantosp/altamira-participa/redis"
	"github.com/levisantosp/altamira-participa/utils"
)

type Session struct {
	ID          string `json:"id"`
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	IsAdmin     bool   `json:"isAdmin"`
}

type ctxKey int

const (
	userCtxKey ctxKey = iota
	sessionCtxKey
)

func Auth(
	api huma.API,
	adminOnly bool,
) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		cookie, err := huma.ReadCookie(ctx, "session")
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		raw, err := redis.Client.Get(ctx.Context(), "session:"+cookie.Value).
			Result()
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		var session Session
		if err := json.Unmarshal([]byte(raw), &session); err != nil {
			utils.LogErr(err)
			huma.WriteErr(
				api,
				ctx,
				http.StatusInternalServerError,
				"Internal Server Error",
			)
			return
		}

		userId, err := strconv.ParseInt(session.UserId, 10, 64)
		if err != nil {
			utils.LogErr(err)
			huma.WriteErr(
				api,
				ctx,
				http.StatusInternalServerError,
				"Internal Server Error",
			)
			return
		}

		user, err := db.Client.User.Get(ctx.Context(), userId)
		if err != nil {
			if generated.IsNotFound(err) {
				huma.WriteErr(api, ctx, http.StatusUnauthorized, "Unauthorized")
				return
			}

			utils.LogErr(err)

			huma.WriteErr(
				api,
				ctx,
				http.StatusInternalServerError,
				"Internal Server Error",
			)
			return
		}

		if adminOnly && !user.IsAdmin {
			huma.WriteErr(
				api,
				ctx,
				http.StatusForbidden,
				"Forbidden",
			)
			return
		}

		ctx = huma.WithValue(ctx, userCtxKey, user)
		ctx = huma.WithValue(ctx, sessionCtxKey, session)

		next(ctx)
	}
}

func GetUserFromContext(ctx context.Context) (*generated.User, bool) {
	user, ok := ctx.Value(userCtxKey).(*generated.User)
	return user, ok
}

func GetSessionFromContext(ctx context.Context) (Session, bool) {
	session, ok := ctx.Value(sessionCtxKey).(Session)
	return session, ok
}
