import axios from "axios";

const BASE_URL = process.env.NEXT_PUBLIC_AUTH_BASE_URL || ""
const PRODUCT_URL = process.env.NEXT_PUBLIC_PRODUCT_BASE_URL || ""

const base = axios.create({
  baseURL: BASE_URL
})

const axiosAuth = axios.create({
  baseURL: BASE_URL,
  withCredentials: true,
  headers: {
    'Content-type': 'application/json',
  }
})

const axiosProduct = axios.create({
  baseURL: PRODUCT_URL,
  withCredentials: true,
})

export { axiosAuth, axiosProduct, base as axios }
