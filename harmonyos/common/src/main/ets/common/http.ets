import axios, { AxiosError, InternalAxiosRequestConfig } from '@ohos/axios'
import { base_url } from '../../../../Index'
import { http } from '@kit.NetworkKit'
//
// const axios_instance = axios.create({
//   baseURL:base_url,
//   timeout:1000,
// })
// axios_instance.interceptors.request.use((config:InternalAxiosRequestConfig)=>{
//
//   return config
// },(err:AxiosError)=>{
//
// })
//
// export default axios_instance

class HttpRequest{

   async post<T>(url:string):Promise<T>{
    let req = http.createHttp()
     let params=base_url+url
     console.log(`url:${params}`)
    let res=await req.request(`${params}`,{
      method:http.RequestMethod.POST,
      header:{
        'Content-Type':'application/x-www-form-urlencoded',
      }
    })
    console.log(`res----------:${JSON.stringify(res)}`)
     let result= JSON.parse(JSON.stringify(res)) as T
    return result
   }

  async get<T>(url:string):Promise<T>{
    let req = http.createHttp()
    let params=base_url+url
    console.log(`url:${params}`)
    let res=await req.request(`${params}`,{
      method:http.RequestMethod.GET,
    })
    console.log(`res----------:${JSON.stringify(res)}`)
    const result= JSON.parse(JSON.stringify(res)) as T
    return result
  }

}
const httpRequest=new HttpRequest()
export default httpRequest