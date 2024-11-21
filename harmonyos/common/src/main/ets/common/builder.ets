@Builder
export function tabBuilder(img:Resource,title:string){
  Column(){
    Image(img)
      .width('20')
      .fillColor(title=='首页'?Color.Black:Color.Black)
    Text(title)
  }.width('100%')
  .justifyContent(FlexAlign.Center)
  .alignItems(HorizontalAlign.Center)
}