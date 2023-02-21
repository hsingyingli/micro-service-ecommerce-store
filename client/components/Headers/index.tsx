import { useAuth } from "@/hooks/useAuth"
import Link from "next/link"
import React from "react"
import { Account } from "./Account"
import { Order } from "./Orders"

const Header: React.FC = () => {

  return (
    <header className="fixed top-0 left-0 w-screen border-b-[1px] border-primary-500 bg-primary-200">
      <div className="max-w-screen-xl mx-auto p-2 h-14
        flex items-center justify-between">
        <Link href={"/"}>
          <h1 className="font-bold text-3xl">Shop</h1>
        </Link>
        <div className="flex items-center gap-4">
          <Order />
          <Account />
        </div>
      </div>
    </header>
  )
}

export default Header
