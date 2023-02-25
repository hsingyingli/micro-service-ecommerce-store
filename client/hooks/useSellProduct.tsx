import { SellContext } from "@/store/providers/SellProvider"
import { useContext } from "react"

const useSellProduct = () => {
  return useContext(SellContext)
}

export {
  useSellProduct
}
