import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// This function can be marked `async` if using `await` inside
export function middleware(request: NextRequest) {

  // auth user can not visit login and sign up page
  if ((request.nextUrl.pathname.startsWith('/signup') ||
    request.nextUrl.pathname.startsWith('/login')) &&
    request.cookies.has('ecommerce-store-refresh-token')
  ) {
    const url = request.nextUrl.clone()
    url.pathname = "/"
    return NextResponse.redirect(url)
  }



  return NextResponse.next()
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: ['/signup', '/login'],
}
