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
  const [products, setProducts] = useState<Array<CartItem>>([])
  const axiosPrivate = useAxiosPrivate("cart")
  const { user } = useAuth()
  useEffect(() => {
    const fetchCart = async () => {
      const res = await axiosPrivate.get("/v1/cart", {
        headers: {
          "Content-Type": "application/json"
        }
      })
      console.log(res.data)
    }

    if (user === null) {
      setProducts([])
      return
    }

    fetchCart()

  }, [user])


  const addProduct = (item: CartItem) => {
    setProducts((prev) => {
      prev.push(item)
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

export type {
  CartItem
}
