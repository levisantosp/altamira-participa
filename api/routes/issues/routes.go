package issues

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/middlewares"
)

func Routes(api huma.API) {
	group := huma.NewGroup(api, "/issues")
	group.UseMiddleware(middlewares.Auth(api, false))

	huma.Get(group, "", GetIssues)
	huma.Post(group, "", CreateIssue)
}
