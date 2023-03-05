import { useAuth } from "@/hooks/useAuth";
import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import { useCart } from "@/hooks/useCart";
import { useImageUrl } from "@/hooks/useImageUrl";
import { CartItem } from "@/store/providers/CartProvider";
import { Product } from "@/store/providers/SellProvider";
import axios from "axios";
import { GetStaticProps, NextPage } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import { useState } from "react";
import { toast } from "react-hot-toast";

interface Props {
  product: Product
}

const ProductPage: NextPage<Props> = ({ product }) => {
  const { user } = useAuth()
  const [amount, setAmount] = useState("")
  const router = useRouter()
  const axiosPrivate = useAxiosPrivate("cart")
  const imageUrl = useImageUrl(product.image_name)
  const { addProduct } = useCart()

  const handleAddCart = async (e: React.FormEvent) => {
    e.preventDefault()
    if (user === null) {
      toast.error("Please login")
      router.push("/login")
      return
    }

    try {
      const res = await axiosPrivate.post("/v1/cart", {
        pid: product.id,
        amount: parseInt(amount)
      }, {
        headers: {
          "Content-Type": "application/json"
        }
      })
      addProduct({
        ...res.data,
        image_name: product.image_name
      } as CartItem)
      toast.success("Success")
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <div className="mt-10 p-5 grid grid-cols-1 md:grid-cols-2">
      <div className="relative w-[min(100%,500px)] aspect-square overflow-hidden rounded-md">
        <Image
          fill alt={product.title} src={imageUrl} />
      </div>
      <div className="mt-10 md:mt-0">
        <h1 className='text-2xl md:text-4xl text-secondary-700 my-2'>{product.title}</h1>
        <div className="mt-5">
          <p className="text-secondary-600 break-words line-clamp-4">{product.description}</p>
        </div>
        <h2 className="text-xl font-medium md:text-3xl text-secondary-600 my-5">
          NT$ {product.price}
          <span> (Remaining: {product.amount})</span>
        </h2>
        {
          user?.id === product.uid ? null :
            (<form
              onSubmit={handleAddCart}
              className="flex flex-col mt-5 max-w-sm"
            >
              <input
                required
                className="py-1 px-2 rounded"
                placeholder="amount"
                type={"number"} min={0} max={product.amount}
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
              />
              <button
                disabled={user?.id === product.uid}
                type="submit"
                className="text-md p-2 bg-secondary-100 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
              > Add to Cart </button>
            </form>)
        }
      </div>
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
