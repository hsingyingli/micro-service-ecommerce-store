import { OrderContext } from "@/store/providers/OrderProvicer"
import { useContext } from "react"

const useOrder = () => {
  return useContext(OrderContext)
}

export { useOrder }
