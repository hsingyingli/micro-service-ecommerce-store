import React, { ReactNode } from "react";

interface Props {
  children: ReactNode,
  number: number
}

const TopRightNumber: React.FC<Props> = ({ children, number }) => {
  return (
    <div className="relative">
      {children}
      <p className={`absolute text-xs rounded-full bg-red-600 w-4 h-4 text-white
         top-[-4px] right-[-4px] text-center ${number === 0 ? "hidden" : null}`}>
        {number}
      </p>
    </div>
  )
}


export {
  TopRightNumber
}
