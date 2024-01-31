import axios from "axios"

const apiUrl = window.palServerHelper.apiUrl
// const apiUrl = "http://localhost:8311"
// console.log("api", apiUrl)

const request = axios.create({
  baseURL: apiUrl,
  timeout: 180000,
})

const errorHandler = async (error: any) => {
  return Promise.reject(error)
}

request.interceptors.request.use((config) => {
  const token = window.localStorage.getItem("server-helper-token")

  if (token) {
    config.headers.Authorization = token
  }
  return config
}, errorHandler)

request.interceptors.response.use((response) => {
  return response.data
}, errorHandler)

export default request

export { request as axios }
