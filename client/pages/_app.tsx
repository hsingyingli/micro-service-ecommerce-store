import { MainLayout } from '@/components/Layouts'
import AuthProvider from '@/store/providers/AuthProvider'
import CartProvider from '@/store/providers/CartProvider';
import OrderProvider from '@/store/providers/OrderProvicer';
import SellProvider from '@/store/providers/SellProvider';
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { Toaster } from 'react-hot-toast';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <AuthProvider>
        <SellProvider>
          <CartProvider>
            <OrderProvider>
              <MainLayout>
                <Component {...pageProps} />
              </MainLayout>
            </OrderProvider>
          </CartProvider>
        </SellProvider>
      </AuthProvider>
      <Toaster />
    </>
  )
}
