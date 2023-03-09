import { ItemCard } from "@/components/ItemCard"
import useAxiosPrivate from "@/hooks/useAxiosPrivate"
import { useOrder } from "@/hooks/useOrder"
import { Disclosure } from "@headlessui/react"
import { ChevronUpIcon } from "@heroicons/react/24/outline"
import { NextPage } from "next"
import Link from "next/link"
import { toast } from "react-hot-toast"

const OrderPage: NextPage = () => {
  const { orders, removeOrder } = useOrder()
  const axiosPayment = useAxiosPrivate("payment")
  const axiosOrder = useAxiosPrivate("order")
  console.log(orders)
  const handlePayment = async (oid: number) => {
    const toastId = toast.loading("...")
    try {
      await axiosPayment.post("/v1/payment", { oid }, { headers: { "Content-Type": "application/json" } })
      removeOrder(oid)
      toast.success("Success", { id: toastId })
    } catch (error) {
      console.log(error)
      toast.error("Fail", { id: toastId })
    }
  }

  const handleDeleteOrder = async (oid: number) => {
    const toastId = toast.loading("...")
    try {
      await axiosOrder.delete(`/v1/order?id=${oid}`)
      removeOrder(oid)
      toast.success("Success", { id: toastId })
    } catch (error) {
      console.log(error)
      toast.error("Fail", { id: toastId })
    }
  }

  return (
    <div className="w-full px-4 pt-16">
      <div className="mx-auto w-full max-w-xl rounded-2xl bg-white p-2">
        {
          orders.map((order) => {
            const total = order.items.reduce((value, curr) => {
              value += curr.price * curr.amount
              return value
            }, 0)
            return (
              <Disclosure key={order.id}>
                {({ open }) => (
                  <>
                    <Disclosure.Button className="flex w-full my-1 justify-between rounded-lg bg-primary-300 px-4 py-2 
                      text-left text-sm font-medium text-secondary-600 hover:bg-primary-500 ">
                      <span>In total: $ {total}</span>
                      <ChevronUpIcon
                        className={`${open ? 'rotate-180 transform' : ''
                          } h-5 w-5 text-primary-900`}
                      />
                    </Disclosure.Button>
                    <Disclosure.Panel className="px-4 pt-4 pb-2 text-sm text-gray-500">
                      {
                        order.items.map((item) => (
                          <Link
                            key={`${order.id}-${item.id}`}
                            href={`/products/${item.pid}`}
                            className="rounded-md hover:bg-secondary-50"
                          >
                            <div className="h-14">
                              <ItemCard item={item} />
                            </div>
                          </Link>
                        ))
                      }
                      <hr className="border-primary-500" />
                      <div className="my-2 flex justify-end gap-2">
                        <button
                          onClick={() => handleDeleteOrder(order.id)}
                          className="text-md py-2 px-4 bg-secondary-300 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
                        >Delete</button>
                        <button
                          onClick={() => handlePayment(order.id)}
                          className="text-md py-2 px-4 bg-secondary-100 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
                        >Buy now</button>
                      </div>
                    </Disclosure.Panel>
                  </>
                )}
              </Disclosure>
            )
          })
        }
      </div>
    </div>
  )
}

export default OrderPage
