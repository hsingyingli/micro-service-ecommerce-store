import useAxiosPrivate from '@/hooks/useAxiosPrivate'
import { createContext, useEffect, useState } from 'react'

type OrderItem = {
  id: number
  oid: number
  pid: number
  price: number
  amount: number
  title: string
  image_name: string
}

type Order = {
  id: number
  items: Array<OrderItem>
}

interface OrderContextInterface {
  orders: Array<Order>
  removeOrder: (id: number) => void
  addOrder: (order: Order) => void
}


const initState: OrderContextInterface = {
  orders: [],
  removeOrder: () => { },
  addOrder: () => { }
}

const OrderContext = createContext<OrderContextInterface>(initState)
interface Props {
  children: React.ReactNode
}

const OrderProvider: React.FC<Props> = ({ children }) => {
  const [orders, setOrders] = useState<Array<Order>>([])
  const axiosPrivate = useAxiosPrivate("order")

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const res = await axiosPrivate.get("/v1/order")
        console.log(res.data)
        if (res.data) {
          setOrders(res.data as Array<Order>)
        }
      } catch (error) {
        console.log(error)
      }
    }

    fetchOrders()
  }, [])

  const removeOrder = (id: number) => {
    const newOrders = orders.filter((order) => order.id != id)
    setOrders(newOrders)
  }
  const addOrder = (order: Order) => {
    setOrders([...orders, order])
  }

  return (
    <OrderContext.Provider value={{ orders, removeOrder, addOrder }}>
      {children}
    </OrderContext.Provider>
  )
}

export default OrderProvider
export {
  OrderContext
}

export type {
  Order, OrderItem
}
