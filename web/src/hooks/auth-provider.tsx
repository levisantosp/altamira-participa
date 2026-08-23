'use client'

import { QueryClientProvider } from '@tanstack/react-query'
import type { Session } from 'api-client'
import { useRouter } from 'next/navigation'
import { createContext, useContext, useEffect } from 'react'
import { Spinner } from 'ui'
import { publicRoutes } from '../config'
import { auth } from '../lib/auth'
import { queryClient } from '../lib/query-client'

type AuthContext = {
  session: Session | null | undefined
  isPending: boolean
}

const AuthContext = createContext<AuthContext | null>(null)

function AuthStateProvider(props: { children: React.ReactNode }) {
  const session = auth.useSession()
  const router = useRouter()

  useEffect(() => {
    if (session.isPending) return

    const isPublicRoute = publicRoutes.get(window.location.pathname)

    if (!session.data && isPublicRoute) return

    if (session.data && isPublicRoute?.action === 'redirect') {
      router.push('/')
    }

    if (!session.data && !isPublicRoute) {
      router.push('/sign-in')
    }
  }, [session.isPending, session.data, router])

  if (session.isPending) {
    return (
      <div className='fixed inset-0 flex gap-1 items-center justify-center bg-background'>
        <Spinner className='size-6' />
        <span>Carregando sessão...</span>
      </div>
    )
  }

  return (
    <AuthContext.Provider
      value={{
        session: session.data,
        isPending: session.isPending
      }}
    >
      {props.children}
    </AuthContext.Provider>
  )
}

export function AuthProvider(props: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthStateProvider>{props.children}</AuthStateProvider>
    </QueryClientProvider>
  )
}

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) {
    throw new Error('"useAuth" must be used inside of "AuthProvider"')
  }

  return ctx
}
