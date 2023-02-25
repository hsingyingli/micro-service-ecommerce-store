import useAxiosPrivate from "@/hooks/useAxiosPrivate"
import { NextPage } from "next"
import Image from "next/image"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import toast from "react-hot-toast"

const CreateProductPage: NextPage = () => {
  const [isLoading, setIsLoading] = useState(false)
  const [image, setImage] = useState<FileList>()
  const [imageUrl, setImageUrl] = useState("")
  const [title, setTitle] = useState("")
  const [price, setPrice] = useState("")
  const [amount, setAmount] = useState("")
  const [description, setDescription] = useState("")
  const axiosPrivate = useAxiosPrivate("product")
  const router = useRouter()

  useEffect(() => {
    if (image === undefined) return

    const imageArray = Array.from(image)
    if (imageArray.length < 1) return

    const newImageUrl = URL.createObjectURL(imageArray[0])
    setImageUrl(newImageUrl)

  }, [image])

  const handleOnImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault()
    if (!e.target.files) {
      return
    }
    setImageUrl("")
    setImage(e.target.files)
  }

  const handleOnSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (image === undefined) return
    setIsLoading(true)
    const toastId = toast.loading('Loading...');

    const formData = new FormData()
    formData.append("image", image[0])
    formData.append('title', title)
    formData.append('price', price)
    formData.append('amount', amount)
    formData.append('description', description)

    try {
      const res = await axiosPrivate.post('/v1/product', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })

      toast.success('Success', {
        id: toastId,
      });
      router.push("/sells")
    } catch (error) {
      toast.success('Fail', {
        id: toastId,
      });
    }

    setIsLoading(false)

  }

  return (
    <div className="flex flex-col items-center mt-10 mx-2">
      <form
        className="flex flex-col gap-2"
        onSubmit={handleOnSubmit}
      >
        <div className="relative overflow-hidden mx-auto w-60 aspect-square 
          border-2 border-primary-500 rounded-xl">
          {imageUrl.length < 1 ? null :
            <Image fill src={imageUrl} alt={"product image"} />
          }
        </div>
        <input
          required
          type={"file"} accept={"image/*"}
          onChange={handleOnImageUpload}
        />
        <label className="mt-5">Title</label>
        <input
          required
          className="py-1 px-2 rounded"
          type={"text"}
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <label>Price</label>
        <input
          required
          className="py-1 px-2 rounded"
          type={"number"} min={0}
          value={price}
          onChange={(e) => setPrice(e.target.value)}
        />
        <label>Amount</label>
        <input
          required
          className="py-1 px-2 rounded"
          type={"number"} min={0}
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />
        <label>Description</label>
        <textarea
          required
          className="py-1 px-2 rounded"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button
          type="submit"
          className="text-md p-2 bg-secondary-100 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
          disabled={isLoading}
        >
          Create
        </button>
      </form>
    </div>
  )
}

export default CreateProductPage
