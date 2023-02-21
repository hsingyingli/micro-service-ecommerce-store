import Head from "next/head"
import React from "react"
import Header from "../Headers"

interface Props {
  children: React.ReactNode
}

const MainLayout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <Head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="description" content={`Micro-Service-Ecommerce-Store`} />
        <meta name="author" content="Hsing Ying Li" />
        <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon" />
        <title>Micro-Service-Ecommerce-Store</title>
      </Head>
      <div className="w-screen min-h-screen bg-primary-500">
        <Header />
        {children}
      </div>
    </>
  )
}

export {
  MainLayout
}
