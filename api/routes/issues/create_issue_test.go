package issues

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/levisantosp/altamira-participa/api/dtos"
	"github.com/levisantosp/altamira-participa/api/tests"

	_ "github.com/levisantosp/altamira-participa/api/ent/generated/runtime"
)

func TestCreateIssue(t *testing.T) {
	tests.Setup(t)

	_, api := humatest.New(t)

	Routes(api)

	user := tests.CreateUser(t)
	session := tests.CreateSession(user, t)

	t.Run("should create an issue", func(t *testing.T) {
		res := api.Post(
			"/issues",
			tests.GetCookie(session.ID),
			map[string]any{
				"title":       strings.Repeat("a", 3),
				"description": strings.Repeat("a", 10),
			},
		)

		if res.Code != http.StatusCreated {
			t.Fatalf("expected status %d got %d", http.StatusCreated, res.Code)
		}

		var issue dtos.Issue
		if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should reject unauthenticated request", func(t *testing.T) {
		res := api.Post(
			"/issues",
			map[string]any{
				"title":       strings.Repeat("a", 3),
				"description": strings.Repeat("a", 10),
			},
		)

		if res.Code != http.StatusUnauthorized {
			t.Fatalf(
				"expected status %d got %d",
				http.StatusUnauthorized,
				res.Code,
			)
		}
	})

	t.Run("should reject missing title", func(t *testing.T) {
		res := api.Post(
			"/issues",
			tests.GetCookie(session.ID),
			map[string]any{
				"description": strings.Repeat("a", 10),
			},
		)

		if res.Code != http.StatusUnprocessableEntity {
			t.Fatalf(
				"expected status %d got %d",
				http.StatusUnprocessableEntity,
				res.Code,
			)
		}
	})

	t.Run("should reject missing description", func(t *testing.T) {
		res := api.Post(
			"/issues",
			tests.GetCookie(session.ID),
			map[string]any{
				"title": strings.Repeat("a", 3),
			},
		)

		if res.Code != http.StatusUnprocessableEntity {
			t.Fatalf(
				"expected status %d got %d",
				http.StatusUnprocessableEntity,
				res.Code,
			)
		}
	})

	t.Run("should reject title shorter than 3 characters", func(t *testing.T) {
		res := api.Post(
			"/issues",
			tests.GetCookie(session.ID),
			map[string]any{
				"title":       strings.Repeat("a", 2),
				"description": strings.Repeat("a", 10),
			},
		)

		if res.Code != http.StatusUnprocessableEntity {
			t.Fatalf(
				"expected status %d got %d",
				http.StatusUnprocessableEntity,
				res.Code,
			)
		}
	})

	t.Run("should reject title longer than 72 characters", func(t *testing.T) {
		res := api.Post(
			"/issues",
			tests.GetCookie(session.ID),
			map[string]any{
				"title":       strings.Repeat("a", 73),
				"description": strings.Repeat("a", 10),
			},
		)

		if res.Code != http.StatusUnprocessableEntity {
			t.Fatalf(
				"expected status %d got %d",
				http.StatusUnprocessableEntity,
				res.Code,
			)
		}
	})

	t.Run(
		"should reject description shorter than 10 characters",
		func(t *testing.T) {
			res := api.Post(
				"/issues",
				tests.GetCookie(session.ID),
				map[string]any{
					"title":       strings.Repeat("a", 3),
					"description": strings.Repeat("a", 9),
				},
			)

			if res.Code != http.StatusUnprocessableEntity {
				t.Fatalf(
					"expected status %d got %d",
					http.StatusUnprocessableEntity,
					res.Code,
				)
			}
		},
	)

	t.Run(
		"should reject description longer than 65000 characters",
		func(t *testing.T) {
			res := api.Post(
				"/issues",
				tests.GetCookie(session.ID),
				map[string]any{
					"title":       strings.Repeat("a", 3),
					"description": strings.Repeat("a", 65_001),
				},
			)

			if res.Code != http.StatusUnprocessableEntity {
				t.Fatalf(
					"expected status %d got %d",
					http.StatusUnprocessableEntity,
					res.Code,
				)
			}
		},
	)
}
