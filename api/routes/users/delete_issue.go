package users

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/ent/generated/issue"
	"github.com/levisantosp/altamira-participa/api/ent/generated/user"
	"github.com/levisantosp/altamira-participa/api/middlewares"
	"github.com/levisantosp/altamira-participa/api/utils"
)

type DeleteIssueOutput struct {
	Body struct {
		DeleteCount int `json:"deleteCount"`
	}
}

func DeleteIssue(
	ctx context.Context,
	input *struct {
		UserID  int64 `path:"userId"`
		IssueID int64 `path:"issueId"`
	},
) (*DeleteIssueOutput, error) {
	userCtx := middlewares.MustGetUserFromContext(ctx)
	if userCtx.ID != input.UserID {
		return nil, utils.LogErr(huma.Error403Forbidden("Forbidden"))
	}

	count, err := db.Client.Issue.Delete().
		Where(issue.HasUserWith(user.IDEQ(input.UserID))).
		Where(issue.IDEQ(input.IssueID)).
		Exec(ctx)
	if err != nil {
		return nil, utils.LogErr(
			huma.Error500InternalServerError("Internal Server Error"),
		)
	}

	if count == 0 {
		return nil, utils.LogErr(
			huma.Error404NotFound("Demanda não encontrada"),
		)
	}

	res := DeleteIssueOutput{}
	res.Body.DeleteCount = count

	return &res, nil
}
