import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import { useCart } from "@/hooks/useCart";
import { useOrder } from "@/hooks/useOrder";
import { getImageUrl } from "@/utils/image";
import { TrashIcon } from "@heroicons/react/24/outline";
import { NextPage } from "next";
import Image from "next/image";
import React, { Fragment, useState } from "react";
import { toast } from "react-hot-toast";

const CartsPage: NextPage = () => {
  const { products, removeProduct, removeBatchProducts } = useCart()
  const axiosCart = useAxiosPrivate("cart")
  const axiosOrder = useAxiosPrivate("order")
  const [checkedState, setCheckedState] = useState<Array<number>>([])
  const [selectAll, setSelectAll] = useState(false)
  const { addOrder } = useOrder()

  const handleOnCheck = (id: number) => {
    const isChecked = checkedState.includes(id)
    if (!isChecked) {
      setCheckedState([...checkedState, id])
    } else {
      const newCheckState = checkedState.filter((i) => i != id)
      setSelectAll(false)
      setCheckedState(newCheckState)
    }
  }

  const handleSelectAll = () => {
    if (selectAll) {
      setCheckedState([])
    } else {
      const newState = products.map((product) => product.id)
      setCheckedState(newState)
    }
    setSelectAll(prev => !prev)
  }

  const handleDeleteCart = async (id: number) => {
    const toastId = toast.loading("Deleting...")
    try {
      await axiosCart.delete(`/v1/cart?id=${id}`, {
        withCredentials: true
      })
      removeProduct(id)
      toast.success("Success", { id: toastId })
    } catch (error) {
      toast.error("Fail", { id: toastId })
    }
  }

  const handleCreateOrder = async (e: React.FormEvent) => {
    e.preventDefault()
    const toastId = toast.loading("Creating...")
    const items = products.filter((product) => checkedState.includes(product.id)).map((product) => {
      return {
        pid: product.pid,
        amount: product.amount
      }
    })

    try {
      console.log(items)
      const res = await axiosOrder.post("/v1/order", { items }, { headers: { "Content-Type": "application/json" } })
      removeBatchProducts(checkedState)
      setCheckedState([])
      setSelectAll(false)
      addOrder(res.data)
      toast.success("Success", { id: toastId })
    } catch (error) {
      console.log(error)
      toast.error("Fail", { id: toastId })
    }

  }

  return (
    <div className="max-w-screen-md w-full mx-auto min-h-screen">
      <form className="overflow-auto mt-10" onSubmit={handleCreateOrder}>
        <button
          type="submit"
          className="text-md py-2 my-5 px-4 bg-secondary-100 
          hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md">
          Create Order
        </button>
        <div className="min-w-[640px] ">
          <div className="grid grid-cols-8 gap-2 items-center">
            <div className="col-start-1 col-end-2 flex items-center justify-center">
              <input
                type="checkbox"
                name={"all"}
                value={"all"}
                checked={selectAll}
                onChange={handleSelectAll}
              />
            </div>
            <div className="col-start-2 col-end-3">
              Image
            </div>
            <div className="col-start-3 col-end-5 text-lg text-center">
              Title
            </div>
            <div className="col-start-5 col-end-6 text-center">
              Amount
            </div>
            <div className="col-start-6 col-end-7 text-center">
              Total Price
            </div>
            <div className="col-start-7 col-end-9 flex justify-center">
              Delete
            </div>
          </div>
          <hr className="my-2 border-primary-500" />
          <div className="grid grid-cols-8 gap-2 gap-y-5 items-center">
            {products.map((product) => {
              const url = getImageUrl(product.image_name)
              return (
                <Fragment key={product.id}>
                  <div className="col-start-1 col-end-2 flex items-center justify-center">
                    <input
                      type="checkbox"
                      id={`cart-checkbox-${product.id}`}
                      name={product.title}
                      value={product.title}
                      checked={checkedState.includes(product.id)}
                      onChange={() => handleOnCheck(product.id)}
                    />
                  </div>
                  <div className="col-start-2 col-end-3">
                    <div className="w-full aspect-square rounded-md relative overflow-hidden">
                      <Image src={url} alt={product.title} fill />
                    </div>
                  </div>
                  <div className="col-start-3 col-end-5 text-lg text-center">
                    {product.title}
                  </div>
                  <div className="col-start-5 col-end-6 text-center">
                    {product.amount}
                  </div>
                  <div className="col-start-6 col-end-7 text-center">
                    ${product.amount * product.price}
                  </div>
                  <div className="col-start-7 col-end-9 flex justify-center">
                    <TrashIcon
                      onClick={() => handleDeleteCart(product.id)}
                      className="w-6 h-6 cursor-pointer" />
                  </div>
                </Fragment>
              )
            })}
          </div>
        </div>
      </form>
    </div>
  )
}

export default CartsPage
