import { useCart } from "@/hooks/useCart";
import { Product } from "@/store/providers/SellProvider";
import axios from "axios";
import { GetStaticProps, NextPage } from "next";
import Image from "next/image";
import { useState } from "react";

interface Props {
  product: Product
}

const ProductPage: NextPage<Props> = ({ product }) => {
  const [amount, setAmount] = useState("")
  const { addProduct } = useCart()

  const handleAddCart = (e: React.FormEvent) => {
    e.preventDefault()
    addProduct(product)
  }

  return (
    <div className="mt-10 p-5">
      <section className="flex gap-5 flex-wrap">
        <div className="relative w-full aspect-square sm:w-1/2 overflow-hidden rounded-md">
          <Image
            fill alt={product.title} src={`data:image/${product.image_type};base64,${product.image_data}`} />
        </div>
        <div className="flex-grow">
          <h1 className='text-2xl md:text-4xl text-secondary-700 my-2'>{product.title}</h1>
          <h2 className="text-xl md:text-3xl text-secondary-600 my-2">NT$ {product.price}</h2>
          <div className="mt-5">
            <h2 className="text-lg text-secondary-600">Description</h2>
            <p className="text-secondary-600">{product.description}</p>
          </div>
          <form
            onSubmit={handleAddCart}
            className="flex flex-col mt-5 max-w-sm"
          >
            <input
              required
              className="py-1 px-2 rounded"
              type={"number"} min={0} max={product.amount}
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
            />
            <button
              type="submit"
              className="text-md p-2 bg-secondary-100 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
            > Add to Cart </button>
          </form>
        </div>
      </section>
    </div>
  )
}

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const productId = params?.productId

  const res = await axios.get(`http://localhost:9011/v1/product?id=${productId}`)
  const product = res.data

  return {
    props: {
      product,
    },
    revalidate: 10, // In seconds
  }
}

export async function getStaticPaths() {
  const res = await axios.get('http://localhost:9011/v1/product/all?limit=50')
  const products: Array<Product> = res.data

  const paths = products.map((product) => ({
    params: { productId: product.id.toString() },
  }))

  return { paths, fallback: 'blocking' }
}

export default ProductPage
