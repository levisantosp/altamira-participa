import { NextRequest, NextResponse } from 'next/server'
import { publicRoutes } from './config'

export default function (request: NextRequest) {
  const sessionCookie = request.cookies.get('session')
  if (!sessionCookie && !publicRoutes.has(request.nextUrl.pathname)) {
    const redirectUrl = request.nextUrl.clone()
    redirectUrl.pathname = '/sign-in'
    return NextResponse.redirect(redirectUrl)
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!_next/static|_next/image|favicon.ico).*)']
}
