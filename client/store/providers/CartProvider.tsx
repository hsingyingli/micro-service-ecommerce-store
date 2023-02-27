import { useAuth } from "@/hooks/useAuth"
import React, { createContext, useEffect, useState } from "react"
import { Product } from "./SellProvider"

type CartItem = {
  id: number,
  imageData: string,
  title: string,
  price: number,
  amount: number,
}

interface CartContextInterface {
  products: Array<Product>
  addProduct: (product: Product) => void
  removeProduct: (id: number) => void
}

const initState: CartContextInterface = {
  products: [],
  addProduct: (product: Product) => { },
  removeProduct: (id: number) => { }
}

const CartContext = createContext<CartContextInterface>(initState)

interface Props {
  children: React.ReactNode
}

const CartProvider: React.FC<Props> = ({ children }) => {
  const [products, setProducts] = useState<Array<Product>>([])
  const { user } = useAuth()
  useEffect(() => {
    const fetchCart = async () => {

    }

    if (user === null) {
      setProducts([])
      return
    }

    fetchCart()

  }, [user])


  const addProduct = (product: Product) => {
    setProducts((prev) => {
      prev.push(product)
      return prev
    })
  }

  const removeProduct = (id: number) => {
    setProducts((prev) => prev.filter((product) => product.id !== id))
  }

  return (
    <CartContext.Provider value={{ products, addProduct, removeProduct }}>
      {children}
    </CartContext.Provider>
  )
}

export default CartProvider

export {
  CartContext
}

