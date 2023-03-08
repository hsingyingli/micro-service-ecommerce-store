import { useAuth } from "@/hooks/useAuth"
import useAxiosPrivate from "@/hooks/useAxiosPrivate"
import React, { createContext, useEffect, useState } from "react"

type CartItem = {
  id: number,
  pid: number,
  amount: number,
  title: string,
  price: number,
  image_name: string,
  created_at: string,
  updated_at: string
}

interface CartContextInterface {
  products: Array<CartItem>
  addProduct: (item: CartItem) => void
  removeProduct: (id: number) => void
}

const initState: CartContextInterface = {
  products: [],
  addProduct: (item: CartItem) => { },
  removeProduct: (id: number) => { }
}

const CartContext = createContext<CartContextInterface>(initState)

interface Props {
  children: React.ReactNode
}

const CartProvider: React.FC<Props> = ({ children }) => {
  const [isFetch, setIsFetch] = useState(false)
  const [products, setProducts] = useState<Array<CartItem>>([])
  const axiosPrivate = useAxiosPrivate("cart")
  const { user } = useAuth()
  useEffect(() => {
    const fetchCart = async () => {
      try {
        const res = await axiosPrivate.get("/v1/cart", {
          headers: {
            "Content-Type": "application/json"
          }
        })
        setProducts(res.data || [])
      } catch (error) {
        setProducts([])
      } finally {
        setIsFetch(true)
      }
    }

    if (user === null) {
      setIsFetch(false)
      setProducts([])
      return
    }
    if (!isFetch) {
      fetchCart()
    }

  }, [user])

  const addProduct = (item: CartItem) => {
    setProducts([...products, item])
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

export type {
  CartItem
}
