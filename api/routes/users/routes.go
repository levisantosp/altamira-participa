package users

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/middlewares"
)

func Routes(api huma.API) {
	group := huma.NewGroup(api, "/users")
	group.UseMiddleware(middlewares.Auth(api, false))

	huma.Get(group, "/{userId}/issues", GetIssues)
	huma.Delete(group, "/{userId}/issues/{issueId}", DeleteIssue)
	huma.Put(group, "/{userId}/issues/{issueId}", EditIssue)
	huma.Delete(group, "", DeleteAccount)
}
