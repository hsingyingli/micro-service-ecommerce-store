import { CartItem } from "@/store/providers/CartProvider";
import { OrderItem } from "@/store/providers/OrderProvicer";
import { getImageUrl } from "@/utils/image";
import Image from "next/image";
import React from "react";

interface Props {
  item: CartItem | OrderItem
}

const ItemCard: React.FC<Props> = ({ item }) => {
  const imageUrl = getImageUrl(item.image_name)

  return (
    <div className="flex flex-row w-full h-full p-2">
      <div className="relative h-full aspect-square overflow-hidden rounded-md">
        <Image src={imageUrl} alt={item.title} fill />
      </div>
      <div className="px-2 flex-grow grid grid-cols-6 gap-2 items-center overflow-x-auto">
        <h1 className="text-sm sm:text-lg col-start-1 col-end-4 truncate">{item.title}</h1>
        <p className="col-start-4 col-end-5">{item.amount}</p>
        <p className="col-start-5 col-end-7">${item.price}</p>
      </div>
    </div>
  )
}

export {
  ItemCard
}
