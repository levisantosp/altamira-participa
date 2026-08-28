package users

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/dtos"
	"github.com/levisantosp/altamira-participa/api/ent/generated"
	"github.com/levisantosp/altamira-participa/api/ent/generated/issue"
	"github.com/levisantosp/altamira-participa/api/ent/generated/user"
	"github.com/levisantosp/altamira-participa/api/middlewares"
	"github.com/levisantosp/altamira-participa/api/utils"
)

type EditIssueOutput struct {
	Body dtos.Issue
}

func EditIssue(ctx context.Context, input *struct {
	UserID  int64 `path:"userId"`
	IssueID int64 `path:"issueId"`
},
) (*EditIssueOutput, error) {
	userCtx := middlewares.MustGetUserFromContext(ctx)
	if userCtx.ID != input.UserID {
		return nil, utils.LogErr(huma.Error403Forbidden("Forbidden"))
	}

	issue, err := db.Client.Issue.UpdateOneID(input.IssueID).
		Where(issue.HasUserWith(user.IDEQ(input.UserID))).
		Save(ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, utils.LogErr(
				huma.Error404NotFound("Demanda não encontrada"),
			)
		}

		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	return &EditIssueOutput{
		Body: dtos.IssueFrom(issue),
	}, nil
}
