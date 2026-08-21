package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/db"
	"github.com/levisantosp/altamira-participa/ent/generated/account"
	"github.com/levisantosp/altamira-participa/ent/generated/user"
	"github.com/levisantosp/altamira-participa/redis"
	"github.com/levisantosp/altamira-participa/utils"
)

type SignInWithEmailInput struct {
	Body struct {
		Email    string `json:"email" format:"email" required:"true"`
		Password string `json:"password" minLength:"8" maxLength:"72" required:"true"`
	}
}

type Session struct {
	ID          string `json:"id"`
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	IsAdmin     bool   `json:"isAdmin"`
}

type SignInOutput struct {
	Status    int           `json:"-"`
	Location  string        `         header:"Location"`
	SetCookie []http.Cookie `         header:"Set-Cookie"`
}

func SignInWithEmail(
	ctx context.Context,
	input *SignInWithEmailInput,
) (*SignInOutput, error) {
	account, err := db.Client.Account.Query().
		Where(account.HasUserWith(user.EmailEQ(input.Body.Email))).
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound(
			"Usuário não encontrado. Verifique o e-mail.",
		)
	}

	ok := utils.VerifyPassword(input.Body.Password, account.Password)
	if !ok {
		return nil, huma.Error401Unauthorized(
			"Credenciais inválidas. Verifique o e-mail e/ou senha.",
		)
	}

	sessionHash := make([]byte, 32)
	_, err = rand.Read(sessionHash)
	if err != nil {
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	sessionId := hex.EncodeToString(sessionHash)

	session, err := json.Marshal(Session{
		ID:          sessionId,
		UserId:      strconv.FormatInt(account.Edges.User.ID, 10),
		Username:    account.Edges.User.Username,
		DisplayName: account.Edges.User.DisplayName,
		IsAdmin:     account.Edges.User.IsAdmin,
	})
	if err != nil {
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	ttl := 7 * 24 * time.Hour
	err = redis.Client.Set(ctx, "session:"+sessionId, session, ttl).Err()
	if err != nil {
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	sessionCookie := http.Cookie{
		Name:     "session",
		Value:    sessionId,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(ttl.Seconds()),
		Path:     "/",
	}

	return &SignInOutput{
		Status:    http.StatusFound,
		Location:  utils.Env.WebURL,
		SetCookie: []http.Cookie{sessionCookie},
	}, nil
}
