import Link from "next/link"
import React from "react"

const Header: React.FC = () => {
  return (
    <header className="fixed top-0 left-0 w-screen border-b-[1px] border-primary-500 bg-primary-200">
      <div className="max-w-screen-xl mx-auto p-2 h-14
        flex items-center">
        <Link href={"/"}>
          <h1 className="font-bold text-3xl">Shop</h1>
        </Link>
      </div>
    </header>
  )
}

export default Header
