import { useImageUrl } from "@/hooks/useImageUrl";
import { CartItem } from "@/store/providers/CartProvider";
import Image from "next/image";
import React from "react";

interface Props {
  item: CartItem
}

const CartItemCard: React.FC<Props> = ({ item }) => {
  const imageUrl = useImageUrl(item.image_name)
  return (
    <div className="flex flex-row w-full h-full p-2">
      <div className="relative h-full aspect-square overflow-hidden rounded-md">
        <Image src={imageUrl} alt={item.title} fill />
      </div>
      <div className="px-2 flex-grow grid grid-cols-5 items-center">
        <h1 className="text-xl col-start-1 col-end-4 truncate">{item.title}</h1>
        <p className="col-start-4 col-end-5">{item.amount}</p>
        <p className="col-start-5 col-end-6">${item.price}</p>
      </div>
    </div>
  )
}

export {
  CartItemCard
}
