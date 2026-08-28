package issues

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/dtos"
	"github.com/levisantosp/altamira-participa/api/ent/generated/issue"
	"github.com/levisantosp/altamira-participa/api/middlewares"
	"github.com/levisantosp/altamira-participa/api/utils"
)

type CreateIssueOutput struct {
	Body dtos.Issue
}

func CreateIssue(
	ctx context.Context,
	input *struct {
		Body struct {
			Title       string       `json:"title" maxLength:"72" minLength:"3" required:"true"`
			Description string       `json:"description" maxLength:"65000" minLength:"10" required:"true"`
			Status      issue.Status `json:"status" enum:"open,closed,in_review" default:"open"`
		}
	},
) (*CreateIssueOutput, error) {
	userCtx := middlewares.MustGetUserFromContext(ctx)

	issue, err := db.Client.Issue.Create().
		SetTitle(input.Body.Title).
		SetDescription(input.Body.Description).
		SetStatus(input.Body.Status).
		SetUserID(userCtx.ID).
		Save(ctx)
	if err != nil {
		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	return &CreateIssueOutput{
		Body: dtos.IssueFrom(issue),
	}, nil
}
