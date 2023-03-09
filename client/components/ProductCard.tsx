import { Product } from "@/store/providers/SellProvider";
import { getImageUrl } from "@/utils/image";
import Image from "next/image";
import React from "react";

interface Props {
  product: Product
  children: React.ReactNode | null
}

const ProductCard: React.FC<Props> = ({ product, children }) => {
  const url = getImageUrl(product.image_name)
  return (
    <div className="rounded-md border-2 border-primary-500 overflow-hidden
               hover:shadow-xl hover:shadow-primary-500 flex flex-row sm:flex-col
               w-[calc(100vw_-_1rem)] aspect-[3/1]
               sm:w-48 sm:aspect-[3/4]">
      <div className="relative aspect-square h-full sm:h-auto w-auto sm:w-full ">
        <Image
          fill alt={product.title} src={url} />
      </div>
      <div className="p-2 flex-grow">
        <h1 className="text-2xl font-medium text-secondary-600">{product.title}</h1>
        <p className="text-secondary-500 text-lg font-medium mt-2">NT$: {product.price}</p>
        <p className="text-secondary-500 text-md font-medium">Remaining: {product.amount}</p>
      </div>
      {children}
    </div>
  )
}

export {
  ProductCard
}
