import { useEffect, useContext, useState } from "react";
import { AuthContext } from "@/store/providers/AuthProvider"
import { axiosAuth, axiosCart, axiosProduct } from "@/utils/axios";
import refreshTokenAPI from "@/utils/refreshTokenAPI";
import { AxiosInstance } from "axios";

const useAxiosPrivate = (service: string) => {
  const { user, updateUser } = useContext(AuthContext);
  const [axiosPrivate] = useState<AxiosInstance>(() => {
    if (service === "product") return axiosProduct
    if (service === "cart") return axiosCart
    return axiosAuth
  })

  useEffect(() => {

    const resIntercept = axiosPrivate.interceptors.response.use(
      response => response,
      async (error) => {
        const prevRequest = error?.config;
        if (error?.response?.status === 401 && !prevRequest?.sent) {
          prevRequest.sent = true
          await refreshTokenAPI()
          return axiosPrivate(prevRequest)
        }

        return Promise.reject(error)
      }
    )
    return () => {
      axiosPrivate.interceptors.response.eject(resIntercept)
    }
  }, [user, updateUser, axiosPrivate])

  return axiosPrivate
}

export default useAxiosPrivate
