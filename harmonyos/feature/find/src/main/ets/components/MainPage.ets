@Component
export struct FindComponent {
  @State message: string = 'Hello World';

  build() {
    Row() {
      Column() {
        Text('发现')
      }
      .width('100%')
    }
    .height('100%')
  }
}
