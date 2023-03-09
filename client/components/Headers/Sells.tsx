import { SquaresPlusIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import React from "react";
import { Tooltip } from "../Tooltip";


const SellsLink = () => {
  return (
    <Tooltip tip="Sell Product">
      <Link href={"/sells"} className="inline-flex w-full justify-center 
            rounded-md bg-secondary-400 p-2 hover:bg-secondary-500 transition-colors duration-150">

        <SquaresPlusIcon
          className="h-5 w-5 text-violet-200 hover:text-violet-100"
          aria-hidden="true"
        />
      </Link>
    </Tooltip>
  )
}

export {
  SellsLink
}
