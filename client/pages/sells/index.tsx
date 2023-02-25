import { useSellProduct } from "@/hooks/useSellProduct";
import { NextPage } from "next";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

const SellsPage: NextPage = () => {
  const [searchText, setSearchText] = useState("")
  const { products } = useSellProduct()

  return (
    <div className="p-2">
      <h1 className="text-2xl font-medium my-10">Experience Hassle-Free Online Selling</h1>
      <div className="flex justify-between">
        <input className="py-1 px-2 rounded" placeholder="Search" value={searchText} onChange={(e) => setSearchText(e.target.value)} />
        <Link href={"/sells/create"} className="text-white py-1 px-4 rounded
          bg-primary-700 hover:bg-primary-800 transition-colors duration-150">Create</Link>
      </div>
      <div>
        {
          products.map((product) => (
            <Image width={100} height={100} alt={product.title} src={`data:image/${product.image_type};base64,${product.image_data}`} />
          ))
        }
      </div>
    </div>
  )
}

export default SellsPage
