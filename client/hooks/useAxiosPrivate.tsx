import { useEffect, useContext, useState } from "react";
import { AuthContext } from "@/store/providers/AuthProvider"
import { axiosAuth, axiosProduct } from "@/utils/axios";
import refreshTokenAPI from "@/utils/refreshTokenAPI";
import { AxiosInstance } from "axios";

const useAxiosPrivate = (service: string) => {
  const { user, updateUser } = useContext(AuthContext);
  const [axiosPrivate, setAxiosPrivate] = useState<AxiosInstance>(() => {
    if (service === "product") return axiosProduct
    return axiosAuth
  })

  useEffect(() => {
    const reqIntercept = axiosPrivate.interceptors.request.use(
      (config) => {
        if (config.headers && !config.headers["Authorization"]) {
          config.headers["Authorization"] = `Bearer ${user?.accessToken}`
        }
        return config
      }, (error) => Promise.reject(error)
    )

    const resIntercept = axiosPrivate.interceptors.response.use(
      response => response,
      async (error) => {
        const prevRequest = error?.config;
        if (error?.response?.status === 401 && !prevRequest?.sent) {
          prevRequest.sent = true
          const user = await refreshTokenAPI()
          const accessToken = user?.accessToken
          prevRequest.headers[`Authorization`] = `Bearer ${accessToken}`
          updateUser(user)
          return axiosPrivate(prevRequest)
        }

        return Promise.reject(error)
      }
    )
    return () => {
      axiosPrivate.interceptors.request.eject(reqIntercept)
      axiosPrivate.interceptors.response.eject(resIntercept)
    }
  }, [user, updateUser, axiosPrivate])

  return axiosPrivate
}

export default useAxiosPrivate
