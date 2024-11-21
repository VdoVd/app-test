import { Route, ZRouter } from '@hzw/zrouter';
import { URL } from 'common/src/main/ets/common/constance';
import httpRequest from 'common/src/main/ets/common/http';
import { LoginData, PhpResponse, Res } from 'common/src/main/ets/common/model';
import { promptAction } from '@kit.ArkUI';
import CryptoJS from '@ohos/crypto-js';
import { PreferencesUtil } from '@pura/harmony-utils';

@Route({name:"login_page"})
@Component
export struct LoginComponent{

  @State email:string=''
  @State pass:string=''
  @StorageLink("user") user:LoginData|undefined=undefined

  async aboutToAppear(): Promise<void> {

  }
  private async submit(){
    let pass = CryptoJS.MD5(this.pass)
    let res = await httpRequest.get<PhpResponse>(`/login?email=${this.email}&pass=${pass}`)

    let tmp=JSON.parse(res.result) as Res<LoginData>
    this.user=tmp.data
    PreferencesUtil.putSync("user",JSON.stringify(this.user))
    promptAction.showDialog({
      message:`${JSON.stringify(this.user)}`
    })

    console.log(`res5:${(JSON.parse(res.result) as Res<LoginData>).data.name}`)
    if(res.responseCode==200){
      console.log(`user:${JSON.stringify(this.user)}`)
      ZRouter.clear()
    }
  }
  build() {
    NavDestination(){
      Column({space:20}){
        TextInput({placeholder:'邮箱',text:$$this.email})
        TextInput({placeholder:'密码',text:$$this.pass})
        Button('登录').width('100%')
          .onClick(()=>{
            if(this.email==''||this.pass==''){
              promptAction.showToast({
                message:'邮箱或者密码为空'
              })
            }else {
              this.submit()
            }
          })
        Row(){
          Text('注册')
            .onClick(()=>{
              ZRouter.push(URL.register_page)
            })
        }.justifyContent(FlexAlign.End)
      }.padding(20)
    }.hideTitleBar(true)
    .onShown(()=>{
      
    })
  }

}