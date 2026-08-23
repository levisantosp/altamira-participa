package users

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/levisantosp/altamira-participa/api/plugins"
)

func Routes(api huma.API) {
	group := huma.NewGroup(api, "/users")
	group.UseMiddleware(plugins.Auth(api, false))

	huma.Get(group, "/{userId}/issues", GetIssues)
}
