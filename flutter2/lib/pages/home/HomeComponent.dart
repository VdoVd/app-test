import 'package:flutter/material.dart';

class HomeContent extends StatefulWidget {
  @override
  Widget build(BuildContext context) {
    return Center(child: Text("Home"));
  }

  @override
  State<StatefulWidget> createState() => _Home();

  @override
  void initState() {
    print("ok-------------------------");
  }
}

class _Home extends State<HomeContent> {
  bool s = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        titleSpacing: 0,
        elevation: 0,
      ),
      body: Column(
        children: [
          TextField(
            decoration: new InputDecoration(hintText: '搜索'),
          )
        ],
      ),
    );
  }
}
