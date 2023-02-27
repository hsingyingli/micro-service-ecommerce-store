import { CartContext } from "@/store/providers/CartProvider"
import { useContext } from "react"

const useCart = () => {
  return useContext(CartContext)
}

export {
  useCart
}
