import Link from "next/link"
import React from "react"
import { Tooltip } from "../Tooltip"
import { AccountMenu } from "./Account"
import { CartMenu } from "./Carts"
import { OrderMenu } from "./Orders"
import { SellsLink } from "./Sells"

const Header: React.FC = () => {

  return (
    <header className="fixed z-50 top-0 left-0 w-screen border-b-[1px] border-primary-500 bg-primary-200">
      <div className="max-w-screen-xl mx-auto p-2 h-14
        flex items-center justify-between">
        <Link href={"/"}>
          <h1 className="font-bold text-3xl">Shop</h1>
        </Link>
        <div className="flex items-center gap-4">
          {/* Sell */}
          <SellsLink />
          <CartMenu />
          {/* Order */}
          <OrderMenu />
          {/* Account */}
          <AccountMenu />
        </div>
      </div>
    </header>
  )
}

export default Header
