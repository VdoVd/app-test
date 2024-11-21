export interface getCodeReq{
  email:string
}
export interface Res<T>{
  code:number
  data:T
  success:boolean
  session:string|null|undefined|number
  message:string|null|undefined|number
}
export interface getCodeRes{

}
export interface RegisterReq{

}
export interface LoginData extends time_m{
  id:number
  name:string
  email:string
  email_verified_at?:string|null|undefined|number|boolean
  password:string
  remember_token?:string|null|undefined|number|boolean

}
export interface time_m{
  created_at?:number|null|undefined
  updated_at?:number|null|undefined
}
export interface PhpResponse {
  responseCode: number
  cookies: string
  header: Header
  result: string
  resultType?: number
  performanceTiming?: PerformanceTiming
}

export interface Header {
  "cache-control": string
  connection: string
  "content-type": string
  date: string
  host: string
  "x-powered-by": string
  "set-cookie": string[]
}

export interface PerformanceTiming {
  dnsTiming: number
  firstReceiveTiming: number
  firstSendTiming: number
  redirectTiming: number
  responseBodyTiming: number
  responseHeaderTiming: number
  tcpTiming: number
  tlsTiming: number
  totalFinishTiming: number
  totalTiming: number
}
export interface box{
  img:Resource
  title:string
}

export interface GetMessData{
  play_num:number
  collect_num:number
  follow_num:number
}