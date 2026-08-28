package issues

import "github.com/danielgtaylor/huma/v2"

func Routes(api huma.API) {
	group := huma.NewGroup(api, "/issues")
	huma.Get(group, "", GetIssues)
}
