import React from "react";

interface Props {
  children: React.ReactNode,
  tip: string
}


const Tooltip: React.FC<Props> = ({ children, tip }) => {


  return (
    <div className="relative group">
      {children}
      <p
        className="absolute z-50 bg-white p-1 text-sm rounded whitespace-nowrap 
                   transition-all duration-150 right-0 hidden  group-hover:block animate-fadin"
      >
        {tip}
      </p>
    </div>
  )
}

export {
  Tooltip
}
