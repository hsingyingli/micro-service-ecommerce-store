import { NextPage } from "next";
import { useRouter } from "next/router";
import { FormEvent, useState } from "react";
import Link from "next/link";
import { signUpAPI } from "@/utils/userAPI";
import toast from "react-hot-toast"

const SignUpPage: NextPage = () => {
  const [isLoading, setIsLoading] = useState(false)
  const [email, setEmail] = useState("")
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const router = useRouter()

  const handleOnSignUp = async (e: FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    const toastId = toast.loading('Loading...');
    const error = await signUpAPI(username, email, password)
    setIsLoading(false)
    if (error === null) {
      toast.success('Create Account Success', {
        id: toastId,
      });
      setUsername("")
      setEmail("")
      setPassword("")
      router.push("/login")
      return
    }
    toast.error('Fail to Create Account', {
      id: toastId,
    });
  }

  return (
    <div className="max-w-sm w-full mx-auto flex flex-col items-center">
      <form
        className="border-gray-400 border-[1px] rounded-md my-5 w-full"
        onSubmit={handleOnSignUp}
      >
        <div className="m-5 flex flex-col">
          <h1 className="font-semibold text-2xl my-5">Sign up</h1>
          <label className="text-md font-semibold my-2">Username</label>
          <input
            className="text-sm border-2 border-secondary-50 rounded-md p-2"
            type={"text"}
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <label className="text-md font-semibold my-2">Email</label>
          <input
            className="text-sm border-2 border-secondary-50 rounded-md p-2"
            type={"email"}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <label className="text-md font-semibold my-2">Password</label>
          <input
            className="text-sm border-2 border-secondary-50 rounded-md p-2"
            type={"password"}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button
            type="submit"
            disabled={isLoading}
            className="text-md p-2 bg-secondary-50 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
          >
            Continue
          </button>
        </div>
      </form>
      <p className="text-gray-600 relative">Has account?</p>
      <Link href="/login"
        className="text-md py-2 px-4 bg-secondary-50 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
      >
        Login
      </Link>
    </div>
  )
}

export default SignUpPage
