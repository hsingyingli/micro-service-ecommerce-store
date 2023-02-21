import { NextPage } from "next";
import { useRouter } from "next/router";
import { FormEvent, useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";

const LoginPage: NextPage = () => {
  const { login } = useAuth()
  const [isSuccess, setIsSuccess] = useState(false)
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const router = useRouter()

  const handleOnLogin = async (e: FormEvent) => {
    e.preventDefault()
    const err = await login(email, password)

    if (err === null) {
      setEmail("")
      setPassword("")
      router.push("/")
    } else {
      setIsSuccess(false)
    }
  }

  return (
    <div className="max-w-sm w-full mx-auto flex flex-col items-center">
      <form
        className="border-gray-400 border-[1px] rounded-md my-5 mx-2 w-full"
        onSubmit={handleOnLogin}
      >
        <div className="m-5 flex flex-col">
          <h1 className="font-semibold text-2xl my-5">Login</h1>
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
            className="text-md p-2 bg-secondary-50 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
          >
            Continue
          </button>
        </div>
      </form>
      <p className="text-gray-600 relative">New?</p>
      <Link href="/signup"
        className="text-md py-2 px-4 bg-secondary-50 hover:bg-secondary-200 transition-colors duration-150 mt-4 rounded-md"
      >
        Create your Amazon account
      </Link>
    </div>
  )
}

export default LoginPage
