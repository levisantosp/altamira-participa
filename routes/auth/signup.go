package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/db"
	"github.com/levisantosp/altamira-participa/ent/generated"
	"github.com/levisantosp/altamira-participa/redis"
	"github.com/levisantosp/altamira-participa/utils"
)

type SignUpWithEmailInput struct {
	Body struct {
		Email       string `json:"email" format:"email" required:"true"`
		Password    string `json:"password" minLength:"8" maxLength:"72" required:"true"`
		DisplayName string `json:"displayName" minLength:"1" maxLength:"100" required:"true"`
		Username    string `json:"username" minLength:"3" maxLength:"32" pattern:"^[a-zA-Z0-9_]+$" required:"true"`
	}
}

func SignUpWithEmail(
	ctx context.Context,
	input *SignUpWithEmailInput,
) (*SignInOutput, error) {
	hash, err := utils.GeneratePasswordHash(input.Body.Password)
	if err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	tx, err := db.Client.Tx(ctx)
	if err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	defer tx.Rollback()

	user, err := tx.User.Create().
		SetEmail(input.Body.Email).
		SetUsername(input.Body.Username).
		SetDisplayName(input.Body.DisplayName).
		Save(ctx)
	if err != nil {
		if generated.IsConstraintError(err) {
			return nil, huma.Error409Conflict(
				"O nome de usuário ou e-mail informado já está cadastrado no sistema.",
			)
		}

		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	_, err = tx.Account.Create().
		SetPassword(*hash).
		SetProvider("email").
		SetUser(user).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	sessionHash := make([]byte, 32)
	_, err = rand.Read(sessionHash)
	if err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	sessionId := hex.EncodeToString(sessionHash)

	session, err := json.Marshal(Session{
		ID:          sessionId,
		UserId:      strconv.FormatInt(user.ID, 10),
		Username:    user.Username,
		DisplayName: user.DisplayName,
		IsAdmin:     user.IsAdmin,
	})
	if err != nil {
		log.Println(err)
		return nil, huma.Error500InternalServerError("Internal Server Error")
	}

	ttl := 7 * 24 * time.Hour
	err = redis.Client.Set(ctx, "session:"+sessionId, session, ttl).Err()
	if err != nil {
		log.Println(err)
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
