import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

Future<void> fetchAlbum() async {
  print("ablium.............");
  var res =
      await http.get(Uri.parse('http://192.168.246.1:8888/article/getList'));

  print(res.reasonPhrase.toString());
}

class FindContent extends StatefulWidget {
  const FindContent({super.key});

  @override
  State<FindContent> createState() => _Find();
}

class _Find extends State<FindContent> {
  String val = '';
  @override
  Widget build(BuildContext context) {
    return Center(child: Text(val));
  }

  @override
  void initState() {
    setState(() {
      val = "change ";
    });
    print('lllllllllllllll');
    fetchAlbum();
  }
}
