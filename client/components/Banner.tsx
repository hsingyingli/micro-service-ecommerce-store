import Image from "next/image";
import React from "react";


const Banner: React.FC = () => {
  return (
    <section className="w-full">
      <div className="w-[max()]">
        <h1 className="uppercase font-bold text-4xl">The Marketplace for Everyone</h1>
        <p className="text-lg font-medium mt-2">Experience Hassle-Free Online Shopping and Selling</p>
      </div>
      <div>
        <Image
          height={"167"}
          width={"121"}
          alt={"banner-mark"}
          src={"/assets/banner-mark.png"} />
      </div>
    </section>
  )
}

export {
  Banner
}
