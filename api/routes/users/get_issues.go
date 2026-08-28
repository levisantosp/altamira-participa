package users

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/dtos"
	"github.com/levisantosp/altamira-participa/api/ent/generated/issue"
	"github.com/levisantosp/altamira-participa/api/ent/generated/user"
	"github.com/levisantosp/altamira-participa/api/utils"
)

type GetUserIssuesOutputBody = utils.PaginatedResponse[dtos.Issue]

type GetUserIssuesOutput struct {
	Body GetUserIssuesOutputBody
}

func GetIssues(
	ctx context.Context,
	input *struct {
		UserID int64  `path:"userId"`
		Status string `          query:"status" enum:"open,closed,in_review"`
		Limit  int    `          query:"limit"                               minimum:"1" maximum:"100" default:"10"`
		Page   int    `          query:"page"                                minimum:"1"               default:"1"`
	},
) (*GetUserIssuesOutput, error) {
	query := db.Client.Issue.Query().
		Where(issue.HasUserWith(user.IDEQ(input.UserID))).
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

	items := make([]dtos.Issue, 0, len(issues))

	for _, item := range issues {
		items = append(items, dtos.IssueFrom(item))
	}

	return &GetUserIssuesOutput{
		Body: utils.PaginatedResponseFrom(items, input.Page, input.Limit),
	}, nil
}
