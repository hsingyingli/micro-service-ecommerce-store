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
    <div className="flex flex-row w-full h-full">
      <div className="relative h-full aspect-square">
        <Image src={imageUrl} alt={item.title} fill />
      </div>
      <div className="px-2 ">
        hello
      </div>
    </div>
  )
}

export {
  CartItemCard
}
