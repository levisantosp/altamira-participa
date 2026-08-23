import { NextRequest, NextResponse } from 'next/server'
import { publicRoutes } from './config'

export default function (request: NextRequest) {
  const isPublicRoute = publicRoutes.get(request.nextUrl.pathname)
  const sessionCookie = request.cookies.get('session')

  if (!sessionCookie && !isPublicRoute) {
    const redirectUrl = request.nextUrl.clone()
    redirectUrl.pathname = '/sign-in'
    return NextResponse.redirect(redirectUrl)
  }

  if (sessionCookie && isPublicRoute?.action === 'redirect') {
    const redirectUrl = request.nextUrl.clone()
    redirectUrl.pathname = '/'
    return NextResponse.redirect(redirectUrl)
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!_next/static|_next/image|favicon.ico).*)']
}
