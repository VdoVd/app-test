import { ZRouter } from '@hzw/zrouter'
import { tabBuilder } from 'common'
import { Home } from 'home/src/main/ets/components/MainPage'
import { ActivityComponent } from 'activity/src/main/ets/components/MainPage'
import { FindComponent } from 'find/src/main/ets/components/MainPage'
import { MineComponent } from 'mine/src/main/ets/components/MainPage'
import { promptAction } from '@kit.ArkUI'
import CryptoJS from '@ohos/crypto-js'

@Entry
@Component
struct Index {
  aboutToAppear(): void {

  }
  @State currentIdx:number=0
  build() {
    Navigation(ZRouter.getNavStack()){
      Tabs({index:this.currentIdx}){
        TabContent(){
          Home()
        }.tabBar(tabBuilder($r('app.media.home'),'首页'))
        TabContent(){
          FindComponent()
        }.tabBar(tabBuilder($r('app.media.find'),'发现'))
        TabContent(){
          ActivityComponent()
        }.tabBar(tabBuilder($r('app.media.activity'),'活动'))
        TabContent(){
          MineComponent()
        }.tabBar(tabBuilder($r('app.media.mine'),'我的'))
      }.barPosition(BarPosition.End).width('100%').height('100%+ ')
    }.height('100%').width('100%').title('首页').titleMode(NavigationTitleMode.Mini)
    .hideTitleBar(true)


  }
}