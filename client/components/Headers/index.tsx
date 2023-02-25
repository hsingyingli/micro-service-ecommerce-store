import { SquaresPlusIcon } from "@heroicons/react/24/outline"
import Link from "next/link"
import React from "react"
import { AccountMenu } from "./Account"
import { OrderMenu } from "./Orders"

const Header: React.FC = () => {

  return (
    <header className="fixed top-0 left-0 w-screen border-b-[1px] border-primary-500 bg-primary-200">
      <div className="max-w-screen-xl mx-auto p-2 h-14
        flex items-center justify-between">
        <Link href={"/"}>
          <h1 className="font-bold text-3xl">Shop</h1>
        </Link>
        <div className="flex items-center gap-4">
          <Link href={"/sells"} className="inline-flex w-full justify-center 
            rounded-md bg-secondary-400 p-2 hover:bg-secondary-500 transition-colors duration-150">
            <SquaresPlusIcon
              className="h-5 w-5 text-violet-200 hover:text-violet-100"
              aria-hidden="true"
            />
          </Link>
          <OrderMenu />
          <AccountMenu />
        </div>
      </div>
    </header>
  )
}

export default Header
