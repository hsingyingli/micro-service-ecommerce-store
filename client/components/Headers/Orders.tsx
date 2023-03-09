import React from "react";
import { ShoppingBagIcon } from '@heroicons/react/24/outline'
import Link from "next/link";
import { Tooltip } from "../Tooltip";
import { TopRightNumber } from "../TopRightNumber";
import { useOrder } from "@/hooks/useOrder";

const OrderMenu: React.FC = () => {
  const { orders } = useOrder()
  return (
    <Tooltip tip="Order list">
      <TopRightNumber number={orders.length}>
        <Link href={"/orders"} className="inline-flex w-full justify-center 
            rounded-md bg-secondary-400 p-2 hover:bg-secondary-500 transition-colors duration-150">
          <ShoppingBagIcon
            className="h-5 w-5 text-violet-200 hover:text-violet-100"
            aria-hidden="true"
          />
        </Link>
      </TopRightNumber>
    </Tooltip>
  )
}

export {
  OrderMenu
}
