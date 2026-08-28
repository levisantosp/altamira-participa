import { queryOptions, useQuery } from '@tanstack/react-query'
import {
  getAuthMeQueryOptions,
  postAuthSignInEmail,
  postAuthSignUpEmail
} from 'api-client'
import { queryClient } from './query-client'

const query = queryOptions({
  ...getAuthMeQueryOptions(),
  staleTime: 5 * 60_000
})

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
  },
  async fetchSession() {
    return await queryClient.query(query).catch(() => null)
  },
  useSession() {
    return useQuery({
      ...query,
      staleTime: 5 * 60_000
    })
  }
} as const
