import { useAuth } from "@/hooks/useAuth";
import React from "react";

const Order: React.FC = () => {
  const { user } = useAuth()
  return (
    <div>O</div>
  )
}

export {
  Order
}
