import { ZRouter } from "@hzw/zrouter";
import { box, GetMessApi, GetMessData, LoginData } from "common";
import { URL } from "common/src/main/ets/common/constance";
import { promptAction } from "@kit.ArkUI";
import { ClickUtil, PreferencesUtil } from "@pura/harmony-utils";

@Component
export struct MineComponent {
  @State message: string = 'Hello World';
  @StorageLink("user") user:LoginData|undefined=undefined
  @StorageLink("user_mess") user_mess:GetMessData|undefined=undefined
  @State isRefreshing:boolean=false
  boxs:box[]=[
    {title:'喜欢',img:$r('app.media.love')},
    {title:'消息',img:$r('app.media.message')},
    {title:'收藏',img:$r('app.media.collection')},
    {title:'历史',img:$r('app.media.history')},
    {title:'设置',img:$r('app.media.setting')},
    {title:'意见反馈',img:$r('app.media.feeback')},
    {title:'积分商城',img:$r('app.media.shop')},
    {title:'新手入驻',img:$r('app.media.news_fill')},
  ]

  async aboutToAppear(): Promise<void> {
    let user=await PreferencesUtil.getString("user")
    this.user=await JSON.parse(user) as LoginData
    await this.getdata();
  }

  private async getdata() {
    if(this.user!=undefined){
      try {
        let user_mess = await GetMessApi(this.user.id);
        this.user_mess = user_mess.data;
      } catch (e) {
        promptAction.showDialog({
          message:`err:${JSON.stringify(e)}`
        })

      }

    }
    setTimeout(()=>{
      this.isRefreshing=false

    },100)
  }

  build() {
      Column({space:15}) {
        Refresh({refreshing:$$this.isRefreshing}){
          Scroll(){
            Column(){
              Column({space:15}){
                Row({space:15}){
                  Image($r('app.media.mine')).width(40).height(40)
                    .backgroundColor(`rgba(0,0,0,0.1)`)
                    .borderRadius(20)
                  Column(){
                    Text(this.user==undefined?'点击登录':this.user.name).fontSize(20).fontWeight(FontWeight.Bold)
                      .textAlign(TextAlign.Start).width('100%')
                    if(this.user==undefined)
                      Text('登录解锁更多功能').fontSize(15).fontColor(Color.Gray)
                        .textAlign(TextAlign.Start).width('100%')
                  }.layoutWeight(1)
                }.width('100%')
                .onClick(()=>{
                  if(this.user==undefined)
                    ZRouter.push(URL.login_page)
                  else {

                  }
                })
                Row(){
                  this.TopPart('播放',0)
                  this.TopPart('获赞',1)
                  this.TopPart('关注',2)
                }.width('100%').justifyContent(FlexAlign.SpaceEvenly)

              }.backgroundColor('rgba(0,0,0,0.05)')
              .padding(15)
              Column(){
                Text('常用功能').textAlign(TextAlign.Start).fontWeight(FontWeight.Bold).width('100%')
                  .padding({left:30})
                  .margin({bottom:15,top:15})
                  .onClick(()=>{
                    promptAction.showDialog({
                      message:`${JSON.stringify(this.user)}`
                    })
                  })
                Flex({wrap:FlexWrap.Wrap}){
                  ForEach(this.boxs,(item:box,index:number)=>{
                    Column({space:10}){
                      Image(item.img).width(30)
                      Text(item.title).fontSize(15)
                        .fontWeight(FontWeight.Regular)
                    }.width('25%').justifyContent(FlexAlign.Center)
                    .alignItems(HorizontalAlign.Center)
                    .margin({bottom:15})
                  })
                }
              }
            }
          }

        }.onRefreshing(()=>{
          ClickUtil.debounce(()=>{
            this.getdata()
          },1000)
        })

      }
      .width('100%')
    .height('100%')
    .justifyContent(FlexAlign.Start)


  }

  @Builder
  TopPart(content:string,curIdx:number) {
    Column() {
      if(this.user_mess==undefined){
        Text('0-')
      }else if(this.user_mess!=undefined){
        if(curIdx==0){
          Text(this.user_mess?.play_num.toString())
        }else if(curIdx==1){
          Text(this.user_mess?.collect_num.toString())
        }else{
          Text(this.user_mess?.follow_num.toString())
        }
      }
      Text(curIdx===0?'播放':(curIdx===1?'获赞':'关注'))
    }
  }
}
