import { GetMessData, PhpResponse, Res } from "../../../../Index"
import httpRequest from "./http"

export const GetMessApi = async (id:number) =>{
  let res=await httpRequest.get<PhpResponse>('/GetMess?id='+id.toString())
  let result = JSON.parse(res.result) as Res<GetMessData>
  return result
}

export const register = () =>{

}