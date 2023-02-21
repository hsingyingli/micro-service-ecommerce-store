import refreshTokenAPI from "@/utils/refreshTokenAPI"
import { loginAPI } from "@/utils/userAPI"
import React, { useEffect, useState, createContext } from "react"
import { useRouter } from "next/router";
import { Loading } from "@/components/Loading";

type User = {
  id: number
  username: string
  email: string
  accessToken: string
}

interface AuthContextInterface {
  user: User | null
  updateUser: (user: User | null) => void
  login: (email: string, password: string) => Promise<Error | null>
}

const initState: AuthContextInterface = {
  user: null,
  updateUser: (user) => { },
  login: async (email: string, password: string) => {
    console.log("default login")
    return null
  }
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
    const fetchUser = async () => {
      const user = await refreshTokenAPI()
      setUser(user)
      setIsLoading(false)
    }
    fetchUser()
  }, [])

  useEffect(() => {
    setIsLoading(true)
    const isNotRequiredAuth = path.includes("/login") || path.includes("/signup")
    const initUser = async () => {
      const user = await refreshTokenAPI()
      setUser(user)

      if (user === null && isNotRequiredAuth) {
        setIsLoading(false)
      } else if (user === null && !isNotRequiredAuth) {
        router.push("/login")
      }
    }

    if (user === null) {
      initUser()
      return
    }

    if (user !== null && isNotRequiredAuth) {
      router.push("/")
    } else {
      setIsLoading(false)
    }

  }, [path, user])


  return (
    !isLoading ?
      <Loading />
      : (
        <AuthContext.Provider value={{ user, updateUser, login }}>
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
