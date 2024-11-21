import 'dart:io';

import 'package:http/http.dart' as http; //导入前需要配置

const base_url = '192.168.246.1:8888';

class Request {
  static get(String url) async {
    var client = http.Client();
    var uri = Uri.parse(base_url + url);
    http.Response response = await client.get(uri);

    if (response.statusCode == HttpStatus.ok) {
      print(response.body);
    } else {
      print('请求失败 code 码为${response.statusCode}');
    }
  }
}
