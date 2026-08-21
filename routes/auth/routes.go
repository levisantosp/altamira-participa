package auth

import "github.com/danielgtaylor/huma/v2"

func AuthRoutes(api huma.API) {
	group := huma.NewGroup(api, "/auth")
	huma.Post(group, "/sign-in/email", SignInWithEmail)
}
