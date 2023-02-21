import { User } from "@/store/providers/AuthProvider"
import { axios } from "./axios"


type LoginSuccessResponse = {
  user: User
  error: null
}

type LoginFailResponse = {
  user: null
  error: Error
}

const loginAPI = async (email: string, password: string): Promise<LoginSuccessResponse | LoginFailResponse> => {
  try {
    const res = await axios.post("/v1/user/login", { email, password }, { withCredentials: true })
    const user: User = res.data
    return {
      user,
      error: null
    }
  } catch (error) {
    return {
      user: null,
      error: error as Error
    }
  }
}

const logoutAPI = async (): Promise<Error | null> => {
  try {
    await axios.post("/v1/user/logout", {}, { withCredentials: true })
    return null
  } catch (error) {
    console.log(error)
    return error as Error
  }
}

const signUpAPI = async (username: string, email: string, password: string): Promise<Error | null> => {
  try {
    await axios.post("/v1/user", { username, email, password })
    return null
  } catch (error) {
    console.log(error)
    return error as Error
  }
}

export {
  loginAPI,
  logoutAPI,
  signUpAPI,
}


