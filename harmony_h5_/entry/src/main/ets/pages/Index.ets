import { webview } from '@kit.ArkWeb'
@Entry
@Component
struct Index {

  ctl:webview.WebviewController=new webview.WebviewController()
  responseResource: WebResourceResponse = new WebResourceResponse();
  localMockUrl: string = 'http://myapp.local';
  aboutToAppear() {
    // 配置Web开启调试模式
    webview.WebviewController.setWebDebuggingAccess(true)
  }
  build() {

      Column(){
        Web({
          controller: this.ctl,
          src: `https://www.example.com/dist/index.html`,
        })
          .domStorageAccess(true)
          .imageAccess(true)
          .javaScriptAccess(true)
          .fileAccess(true)
          .mediaPlayGestureAccess(false)
          .databaseAccess(true)
          .mixedMode(MixedMode.All)
          .onPageEnd(() => {
            // this.loadHandle()
          })
          .onInterceptRequest((event) => {
            if (!event) {
              return
            }


            const requestUrl = event.request.getRequestUrl()
            console.log(`requestUrl ${requestUrl}`)
            if (requestUrl.startsWith("https://www.example.com/")) {
              let rawfileName = requestUrl.replace("https://www.example.com/", "")
              let mimeType = "application/octet-stream"
              if (rawfileName.endsWith("js")) {
                mimeType = "text/javascript"
              } else if (rawfileName.endsWith("css")) {
                mimeType = "text/css"
              } else if (rawfileName.endsWith("woff")) {
                mimeType = "font/woff"
              } else if (rawfileName.endsWith("html")) {
                mimeType = "text/html"
              }
              if(!rawfileName.endsWith("html")){
                rawfileName='dist/'+rawfileName
              }
              console.log(`rawfileName ${rawfileName}`)
              let response = new WebResourceResponse();
              // 构造响应数据，如果本地文件在rawfile下，可以通过如下方式设置
              response.setResponseData($rawfile(rawfileName));
              // response.setResponseEncoding('utf-8');
              response.setResponseMimeType(mimeType);
              response.setResponseCode(200);
              // response.setReasonMessage('OK');
              response.setResponseIsReady(true);
              return response;
            }

            return null
          })
          .layoutWeight(1)
      }
  }

}