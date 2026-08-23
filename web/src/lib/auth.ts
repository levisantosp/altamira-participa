import { postAuthSignInEmail, postAuthSignUpEmail } from 'api-client'

export const auth = {
  signIn: {
    async email(email: string, password: string) {
      await postAuthSignInEmail({
        body: {
          email,
          password
        }
      })
    }
  },
  signUp: {
    async email(body: {
      email: string
      password: string
      username: string
      displayName: string
    }) {
      await postAuthSignUpEmail({ body })
    }
  }
} as const
