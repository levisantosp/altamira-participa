package users

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/db"
	"github.com/levisantosp/altamira-participa/ent/generated/issue"
	"github.com/levisantosp/altamira-participa/ent/generated/user"
	"github.com/levisantosp/altamira-participa/utils"
)

type Issue struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      issue.Status `json:"status"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type GetUserIssuesOutputBody = utils.PaginatedResponse[Issue]

type GetUserIssuesOutput struct {
	Body GetUserIssuesOutputBody
}

func GetUserIssues(
	ctx context.Context,
	input *struct {
		ID     int64  `path:"id"`
		Status string `          query:"status" enum:"open,closed,in_review"`
		Limit  int    `          query:"limit"                               minimum:"1" maximum:"100" default:"10"`
		Page   int    `          query:"page"                                minimum:"1"               default:"1"`
	},
) (*GetUserIssuesOutput, error) {
	query := db.Client.Issue.Query().
		Where(issue.HasUserWith(user.IDEQ(input.ID))).
		Order(issue.ByCreatedAt(sql.OrderDesc())).
		Offset((input.Page - 1) * input.Limit).
		Limit(input.Limit + 1)

	if input.Status != "" {
		query = query.Where(issue.StatusEQ(issue.Status(input.Status)))
	}

	issues, err := query.All(ctx)
	if err != nil {
		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	items := make([]Issue, 0, len(issues))

	for _, item := range issues {
		items = append(items, Issue{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &GetUserIssuesOutput{
		Body: utils.PaginatedResponseFrom(items, input.Page, input.Limit),
	}, nil
}
