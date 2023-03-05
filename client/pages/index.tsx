import { ProductCard } from "@/components/ProductCard";
import { Product } from "@/store/providers/SellProvider";
import axios from "axios";
import { NextPage } from "next";
import Image from "next/image";
import Link from "next/link";

interface Props {
  products: Array<Product>
}

const HomePage: NextPage<Props> = ({ products }) => {
  return (
    <div className="mt-10">
      <div className="relative w-full aspect-[2/1] sm:aspect-[3/1] rounded-md overflow-hidden">
        <Image src={"/assets/banner.png"} alt="banner" fill />
      </div>
      <section className="flex flex-wrap gap-5 justify-start mt-10 p-2">
        {products.map((product) => (
          <Link href={`/product/${product.id}`} key={product.id}>
            <ProductCard
              key={product.id}
              product={product}>
            </ProductCard>
          </Link>
        ))}
      </section>
    </div>
  )
}

export async function getStaticProps() {
  let products: Array<Product> = []
  try {
    const res = await axios.get('http://localhost:9011/v1/product/all', {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    products = res.data
  } catch (error) {
    console.log(error)
  }

  return {
    props: {
      products
    },
    revalidate: 30,
  }

}

export default HomePage
