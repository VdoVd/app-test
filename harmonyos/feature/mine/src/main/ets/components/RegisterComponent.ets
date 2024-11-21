import { Route, ZRouter } from '@hzw/zrouter';
import { ClickUtil } from '@pura/harmony-utils';
import axios_instance from 'common/src/main/ets/common/http';
import { promptAction } from '@kit.ArkUI';
import httpRequest from 'common/src/main/ets/common/http';
import { PhpResponse, Res, URL } from 'common';
import CryptoJS from '@ohos/crypto-js';

@Route({name:"register_page"})
@Component
export struct RegisterComponent{

  @State email:string=''
  @State pass:string=''
  @State code:string=''
  @State sec:number=60
  @State old_pass:string=''
  @State name:string=''
  private async submit(pass:string){
    const res=await httpRequest.get<PhpResponse>(`/register?email=${this.email}&code=${this.code}&pass=${pass}&name=${this.name}`)
    return res
  }
  private async sendCode(){
    const res=await httpRequest.get<Res<boolean>>(`/getCode?email=${this.email}`)
    promptAction.showDialog({
      message:`${JSON.stringify(res)}`
    })
  }
  build() {
    NavDestination(){
      Column({space:15}){
        TextInput({placeholder:'email',text:$$this.email})
        TextInput({placeholder:'name',text:$$this.name})
        Row({space:15}){
          TextInput({placeholder:'code',text:$$this.code}).layoutWeight(1)
          Button(this.sec==60?'验证码':`还剩${this.sec}秒`)
            .onClick(()=>{
              ClickUtil.throttle(()=>{
                this.sendCode()
                let sid=setInterval(()=>{
                  this.sec--
                  if(this.sec==0){
                    this.sec=60
                    clearInterval(sid)

                  }
                },1000)
              },1000*60)
            })

        }
        TextInput({placeholder:'pass',text:$$this.pass})
        TextInput({placeholder:'pass',text:$$this.old_pass})
        Button('注册')
          .onClick(async ()=>{
            if(this.pass!=this.old_pass){
              promptAction.showToast({
                message:'密码不一致，请重新填写'
              })
            }else if(this.email==''){
              promptAction.showToast({
                message:'邮箱未填写'
              })
            }else{
              let pass:string= CryptoJS.MD5(this.pass).toString()
              let res = await this.submit(pass)
              promptAction.showDialog({
                message:`${JSON.stringify(res)}`
              })
              if(res.responseCode==200){
                promptAction.showToast({
                  message:'注册成功'
                })
                ZRouter.push(URL.login_page)
              }
            }
          })
      }
    }.hideTitleBar(true)

  }

}