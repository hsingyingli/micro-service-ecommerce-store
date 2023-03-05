import refreshTokenAPI from "@/utils/refreshTokenAPI"
import { loginAPI, logoutAPI } from "@/utils/userAPI"
import React, { useEffect, useState, createContext } from "react"
import { useRouter } from "next/router";
import { Loading } from "@/components/Loading";

type User = {
  id: number
  username: string
  email: string
}

interface AuthContextInterface {
  user: User | null
  updateUser: (user: User | null) => void
  login: (email: string, password: string) => Promise<Error | null>
  logout: () => Promise<Error | null>
}

const initState: AuthContextInterface = {
  user: null,
  updateUser: (user) => { },
  login: async (email: string, password: string) => null,
  logout: async () => null,
}

const AuthContext = createContext<AuthContextInterface>(initState)

interface Props {
  children: React.ReactNode
}

const AuthProvider: React.FC<Props> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const router = useRouter()
  const path = router.asPath

  const updateUser = (user: User | null) => {
    setUser(user)
  }

  const logout = async () => {
    const error = await logoutAPI()
    if (error === null) {
      setUser(null)
      router.push("/")
    }
    return error
  }

  const login = async (email: string, password: string) => {
    const { user, error } = await loginAPI(email, password)
    setUser(user)
    return error
  }

  useEffect(() => {
    const handleFinishRedirect = () => setIsLoading(false)
    router.events.on("routeChangeComplete", handleFinishRedirect)
    return () => router.events.off("routeChangeComplete", handleFinishRedirect)
  }, [])

  useEffect(() => {
    setIsLoading(true)
    const initUser = async () => {
      const user = await refreshTokenAPI()
      setUser(user)
      setIsLoading(false)
    }
    initUser()
  }, [])

  return (
    isLoading ?
      <Loading />
      : (
        <AuthContext.Provider value={{ user, updateUser, login, logout }}>
          {children}
        </AuthContext.Provider >
      )
  )
}

export default AuthProvider
export { AuthContext }
export type {
  User
}
