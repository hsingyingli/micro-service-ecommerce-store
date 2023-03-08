import React, { Fragment, useEffect, useState } from "react";
import { Popover, Transition } from '@headlessui/react'
import { useAuth } from "@/hooks/useAuth";
import { ShoppingCartIcon } from '@heroicons/react/24/outline'
import Link from "next/link";
import { Tooltip } from "../Tooltip";
import { useCart } from "@/hooks/useCart";
import { TopRightNumber } from "../TopRightNumber";
import { CartItemCard } from "../CartItemCard";

const CartMenu: React.FC = () => {
  const { user } = useAuth()
  const { products } = useCart()
  return (
    <div className="relative z-30">
      <Popover className="relative">
        {({ open }) => (
          <>
            <Tooltip tip="Shopping Cart">
              <TopRightNumber number={products.length}>
                <Popover.Button className={`inline-flex w-full justify-center 
            rounded-md ${open ? "bg-secondary-500" : "bg-secondary-400"} p-2 hover:bg-secondary-500 transition-colors duration-150`}>

                  <ShoppingCartIcon
                    className="h-5 w-5 text-violet-200 hover:text-violet-100"
                    aria-hidden="true"
                  />

                </Popover.Button>
              </TopRightNumber>
            </Tooltip>
            <Transition
              as={Fragment}
              enter="transition ease-out duration-200"
              enterFrom="opacity-0 translate-y-1"
              enterTo="opacity-100 translate-y-0"
              leave="transition ease-in duration-150"
              leaveFrom="opacity-100 translate-y-0"
              leaveTo="opacity-0 translate-y-1"
            >
              <Popover.Panel className="absolute right-0 z-30 mt-3 w-screen max-w-[15rem] sm:max-w-xs px-4 sm:px-0 ">
                <div className="bg-white overflow-hidden rounded-md shadow-lg">
                  <div className="flex flex-col p-1 gap-1">
                    {user === null ?
                      <Link
                        href={"/login"}
                        className="flex w-full items-center rounded-md p-2 
                      hover:bg-secondary-100"
                      >
                        Login
                      </Link>
                      :
                      <>
                        {products.map((product) => (
                          <Link
                            key={product.id}
                            href={`/product/${product.id}`}
                            className="h-14 w-full rounded-md hover:bg-secondary-50"
                          >
                            <CartItemCard item={product} />
                          </Link>
                        ))}
                        <Link href="/carts" className="text-center underline">
                          View All
                        </Link>
                      </>
                    }
                  </div>
                </div>
              </Popover.Panel>

            </Transition>
          </>
        )}



      </Popover>
    </div>
  )
}

export {
  CartMenu
}
