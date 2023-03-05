import { useAuth } from "@/hooks/useAuth";
import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import React, { createContext, useEffect, useState } from "react";

type Product = {
  id: number
  uid: number
  title: string
  price: number
  amount: number
  description: string
  image_name: string
  createdAt: string
  updatedAt: string
}


interface SellContextInterface {
  products: Array<Product>
  removeProduct: (id: number) => void
  addProduct: (product: Product) => void
}

const initState: SellContextInterface = {
  products: [],
  removeProduct: (id: number) => { },
  addProduct: (product: Product) => { }
}

const SellContext = createContext<SellContextInterface>(initState)

interface Props {
  children: React.ReactNode
}

const SellProvider: React.FC<Props> = ({ children }) => {
  const [products, setProducts] = useState<Array<Product>>([])
  const { user } = useAuth()
  const axiosPrivate = useAxiosPrivate("product")


  useEffect(() => {
    const fetchSellProduct = async () => {
      try {
        const res = await axiosPrivate.get("/v1/product/auth/all")
        setProducts(res.data as Array<Product>)
      } catch (error) {
      }
    }

    if (user === null) {
      setProducts([])
      return
    }

    fetchSellProduct()

  }, [user])

  const removeProduct = (id: number) => {
    setProducts((prev) => prev.filter((product) => product.id != id))
  }

  const addProduct = (product: Product) => {
    setProducts((prev) => {
      prev.push(product)
      return prev
    })
  }

  return (
    <SellContext.Provider value={{ products, removeProduct, addProduct }}>
      {children}
    </SellContext.Provider>
  )
}


export default SellProvider
export {
  SellContext
}

export type {
  Product
}

