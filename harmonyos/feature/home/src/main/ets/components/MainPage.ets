import { ZRouter } from "@hzw/zrouter";
import { URL } from "common";
import avPlayManage from "common/src/main/ets/common/utils/AvPlayManager";

@Component
export struct Home {
  @State value:string=""
  @State message: string = 'Hello World'
  private controller: TabsController = new TabsController()
  private controller1: TabsController = new TabsController()
  @State selfFadingFade: boolean = true;
  private surfaceId: string = '';
  private timeout: number = 0; // Timer ID
  private xComponentController: XComponentController = new XComponentController();
  @State avPlayManage: avPlayManage = new avPlayManage();
  @State isSwiping: boolean = false;
  @State isClickScreen: boolean = false;
  @State flag: boolean = true; // Pause Playback
  @State XComponentFlag: boolean = false;
  @State speedSelect: number = 0;
  @State videoListSelect: number = 0;
  @State durationTime: number = 0;
  @State currentTime: number = 0;
  @State surfaceW: number = 0;
  @State surfaceH: number = 0;
  @State show: boolean = false; // Indicates whether the videoPanel component is displayed.
  @State videoSelect: number = 0;
  @State percent: number = 0;
  @State windowWidth: number = 300;
  @State windowHeight: number = 300;
  @State isCalcWHFinished: boolean = false;
  @StorageLink('videoName') videoName: Resource = $r('app.string.video_res_1');
  @StorageLink('videoIndex') videoIndex: number = 0;
  build() {
      Column() {
        Row(){
          Image($r('app.media.shop'))
            .width(25)
            .height(25)
          TextInput({placeholder:'请输入关键字',text:this.value}).layoutWeight(1)
        }.width('100%').padding({left:10,right:10})
        Tabs({ barPosition: BarPosition.Start, controller: this.controller }) {
          TabContent() {
            Scroll(){
              Column(){

                Row(){
                  Text('健康科普大讲堂').fontWeight(FontWeight.Medium).fontSize(25)
                  Text('更多 >').fontColor(Color.Gray).fontSize(17)
                }.width('100%').justifyContent(FlexAlign.SpaceBetween)

                Text('video').onClick(()=>{
                  ZRouter.push(URL.video_page)
                })
              }.padding(20)
            }.height('100%')
            .width('100%')

          }.tabBar('直播')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Yellow)
          }.tabBar('精选')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Blue)
          }.tabBar('24小时')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Green)
          }.tabBar('热门')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Green)
          }.tabBar('青年KOL')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Green)
          }.tabBar('i广东')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Green)
          }.tabBar('湾澳台')

          TabContent() {
            Column().width('100%').height('100%').backgroundColor(Color.Green)
          }.tabBar('融媒号')
        }
        .vertical(false)
        .scrollable(true)
        .barMode(BarMode.Scrollable)
        .barHeight(40)
        .animationDuration(400)
        .onChange((index: number) => {
          console.info(index.toString())
        })
        .fadingEdge(this.selfFadingFade)
        .layoutWeight(1)
        .width('100%')
      }
      .width('100%')
      .height('100%')
    .justifyContent(FlexAlign.Start)
  }
}
