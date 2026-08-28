package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/middlewares"
	"github.com/levisantosp/altamira-participa/api/redis"
	"github.com/levisantosp/altamira-participa/api/utils"
)

type GetMeOutput struct {
	Body middlewares.Session
}

func GetMe(ctx context.Context, input *struct {
	Session http.Cookie `cookie:"session"`
},
) (*GetMeOutput, error) {
	if input.Session.Name == "" {
		return nil, huma.Error404NotFound("Not Found")
	}

	raw, err := redis.Client.Get(ctx, "session:"+input.Session.Value).Result()
	if err != nil {
		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	var session middlewares.Session
	if err := json.Unmarshal([]byte(raw), &session); err != nil {
		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	return &GetMeOutput{
		Body: session,
	}, nil
}
