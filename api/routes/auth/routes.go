package auth

import "github.com/danielgtaylor/huma/v2"

func Routes(api huma.API) {
	group := huma.NewGroup(api, "/auth")
	huma.Post(group, "/sign-in/email", SignInWithEmail)
	huma.Post(group, "/sign-up/email", SignUpWithEmail)
	huma.Get(group, "/me", GetMe)
}
