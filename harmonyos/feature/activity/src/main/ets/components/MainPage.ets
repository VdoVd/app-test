

@Component
export struct ActivityComponent {
  @State message: string = 'Hello World';

  build() {
    Row() {
      Column() {
        Text('Activity')
      }
      .width('100%')
    }
    .height('100%')
  }
}
