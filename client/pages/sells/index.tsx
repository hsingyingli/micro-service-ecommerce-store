import { ProductCard } from "@/components/ProductCard";
import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import { useSellProduct } from "@/hooks/useSellProduct";
import { Product } from "@/store/providers/SellProvider";
import { ArrowTopRightOnSquareIcon, PencilSquareIcon, TrashIcon } from "@heroicons/react/24/outline";
import { NextPage } from "next";
import Link from "next/link";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";

const SellsPage: NextPage = () => {
  const [searchText, setSearchText] = useState("")
  const [displayProducts, setDisplayProducts] = useState<Array<Product>>([])
  const { products, removeProduct } = useSellProduct()
  const axiosPrivate = useAxiosPrivate("product")

  useEffect(() => {
    const selectedProduct = products.filter((product) => product.title.includes(searchText))
    setDisplayProducts(selectedProduct)
  }, [searchText, products])


  const handleOnDelete = async (id: number) => {
    const toastId = toast.loading("Deleting...")
    try {
      await axiosPrivate.delete(`/v1/product/auth?id=${id}`)
      toast.success('Deleted', {
        id: toastId,
      });
      removeProduct(id)
    } catch (error) {
      toast.error('Try Later', {
        id: toastId,
      });
    }
  }

  return (
    <div className="p-2">
      <h1 className="text-4xl text-secondary-500 font-medium my-10">Experience Hassle-Free Online Selling</h1>
      <div className="flex justify-between flex-col sm:flex-row gap-2">
        <input className="py-1 px-2 rounded" placeholder="Search" value={searchText} onChange={(e) => setSearchText(e.target.value)} />
        <Link href={"/sells/create"} className="text-white py-1 px-4 rounded
          bg-primary-700 hover:bg-primary-800 transition-colors duration-150">Create</Link>
      </div>
      <div className="mt-10 flex gap-5 flex-wrap">
        {
          displayProducts.map((product) => (
            <ProductCard key={product.id} product={product}>
              <div className="flex flex-col sm:flex-row sm:justify-start justify-center gap-2 px-2 pb-1">
                <Link href={`/products/${product.id}`}>
                  <ArrowTopRightOnSquareIcon
                    className="w-6 h-6 text-secondary-500 hover:text-secondary-700"
                  />
                </Link>
                <button onClick={() => handleOnDelete(product.id)}>
                  <TrashIcon
                    className="w-6 h-6 text-secondary-500 hover:text-secondary-700"
                  />
                </button>
              </div>
            </ProductCard>
          ))
        }
      </div >
    </div >
  )
}

export default SellsPage
