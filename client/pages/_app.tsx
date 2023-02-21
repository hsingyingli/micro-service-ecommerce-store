import { MainLayout } from '@/components/Layouts'
import AuthProvider from '@/store/providers/AuthProvider'
import '@/styles/globals.css'
import type { AppProps } from 'next/app'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <MainLayout>
      <AuthProvider>
        <Component {...pageProps} />
      </AuthProvider>
    </MainLayout>
  )
}
