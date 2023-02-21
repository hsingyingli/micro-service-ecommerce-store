import { MainLayout } from '@/components/Layouts'
import AuthProvider from '@/store/providers/AuthProvider'
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { Toaster } from 'react-hot-toast';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <AuthProvider>
        <MainLayout>
          <Component {...pageProps} />
        </MainLayout>
      </AuthProvider>
      <Toaster />
    </>
  )
}
