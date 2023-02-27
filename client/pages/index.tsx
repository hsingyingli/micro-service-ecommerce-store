import { ProductCard } from "@/components/ProductCard";
import { Product } from "@/store/providers/SellProvider";
import axios from "axios";
import { NextPage } from "next";
import Link from "next/link";

interface Props {
  products: Array<Product>
}

const HomePage: NextPage<Props> = ({ products }) => {
  return (
    <div className="mt-10 flex flex-wrap gap-5 justify-center">
      {products.map((product) => (
        <Link href={`/product/${product.id}`} key={product.id}>
          <div key={product.id}
            className="rounded-md border-2 border-primary-500 overflow-hidden
              w-[13rem] aspect-[6/7] hover:shadow-xl hover:shadow-primary-500">
            <ProductCard product={product} />
          </div>
        </Link>
      ))}
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
