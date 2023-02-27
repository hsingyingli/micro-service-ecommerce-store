import { Product } from "@/store/providers/SellProvider";
import Image from "next/image";
import React from "react";

interface Props {
  product: Product
}

const ProductCard: React.FC<Props> = ({ product }) => {
  return (
    <div className="w-full h-full ">
      <div className="relative w-full aspect-[4/3] ">
        <Image
          fill alt={product.title} src={`data:image/${product.image_type};base64,${product.image_data}`} />
      </div>
      <div className="p-2">
        <h1 className="text-lg text-secondary-600">{product.title} <span className="text-sm ml-2 text-secondary-500">({product.amount})</span></h1>
        <p className="text-secondary-500 text-sm">NT$: {product.price}</p>
      </div>
    </div>
  )
}

export {
  ProductCard
}
