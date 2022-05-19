// Copyright 2018 The Flutter team. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'package:http/http.dart' as http;
import 'package:http/http.dart';
import 'ablerdp.dart';

// import 'ablewebview.dart';
// import 'package:system_tray/system_tray.dart';
import 'package:bitsdojo_window/bitsdojo_window.dart';

Uri uri = Uri();
String uriString = '';
var reg;
var rdp;
String Status = "";
String UUID = "";

void main(List<String> args) async {
  if (args.isEmpty) {
    print("need args");
    exit(0);
  }
  WidgetsFlutterBinding.ensureInitialized();
  //runApp(const MyApp());
  // doWhenWindowReady(() {
  //   final win = appWindow;
  //   const initialSize = Size(600, 450);
  //   win.minSize = initialSize;
  //   win.size = initialSize;
  //   win.alignment = Alignment.center;
  //   win.title = "How to use system tray with Flutter";
  //   // win.show();
  //   win.hide();
  // });
  appWindow.hide();
  var ls = LocalStorage();
  uriString = args.join(" ");
  // var file = File('c:\\ablecloud\\file.txt');
  // var sink = file.openWrite();
  // sink.write('FILE ACCESSED ${DateTime.now()}\n');
  Status += 'FILE ACCESSED ${DateTime.now()}\n';

  ls.writeReg(uriString);
  Status += 'reg writed ${DateTime.now()}\n';

  uri = Uri.parse(uriString);
  String hashedpassword = 'qwer';
  if (uri.queryParameters.containsKey('password')) {
    var password = uri.queryParameters['password'];
    rdpPassword(password!).then((value) {
      hashedpassword = value.trim();
      Status += 'password $hashedpassword ${DateTime.now()}\n';
      //if(path.queryParameters.containsKey('password 51')){
      Map<String, List<String>> map = Map.from(uri.queryParametersAll);
      if (map.containsKey('instanceUuid')) {
        UUID = map['instanceUuid']!.join();
      }
      print("map: $map");
      map.remove('password');
      print("map: $map");
      map['password 51'] = [hashedpassword];
      Uri p = Uri(
          scheme: uri.scheme,
          path: uri.path,
          queryParameters: map,
          host: uri.host,
          port: uri.port);
      //uri += '&password 51=$hashedpassword';
      uriString = p.toString();
      print("p: $p");
      ls.writeRDP(uriString);
      Status += 'rdp writed ${DateTime.now()}\n';

      regAdd(ls);
      Status += 'reg added ${DateTime.now()}\n';

      postHttpJson(p, UUID);
      rdpLaunch(ls);
      Status += 'rdp launched ${DateTime.now()}\n';

      // ls.deleteRDP();
    });
  } else {
    ls.writeRDP(uriString);
    Status += 'rdp writed ${DateTime.now()}\n';

    regAdd(ls);
    Status += 'reg added ${DateTime.now()}\n';

    rdpLaunch(ls);
    Status += 'rdp launched ${DateTime.now()}\n';
    ls.deleteReg();
    ls.deleteRDP();
  }
  // Close the IOSink to free system resources.
  // sink.close();
  //exit(0);
  //   }).catchError((error) {
  //     Status += 'rdp launch error $error ${DateTime.now()}\n';
  //   });
  // }).catchError((error) {
  //   Status += 'reg add error $error ${DateTime.now()}\n';
  // });
  // }).catchError((error) {
  //   Status += 'rdp write error $error ${DateTime.now()}\n';
  // });
  // }).catchError((error){
  //   Status += 'reg write error $error ${DateTime.now()}\n';
  // });
  // Status += 'rdp Writed ${DateTime.now()}\n';

  if (kDebugMode) {
    print(args);
    print(Status);
  }
  // for (var val in args) {
  //   uri += val + "
  // }
}

Future<Map> fetchHttpJson() async {
  var response = await http
      .get(Uri.parse('https://jsonplaceholder.typicode.com/albums/1'));
  return json.decode(response.body);
}

// ignore: non_constant_identifier_names
Future<Map> postHttpJson(Uri uri, String UUID) async {
  var client = http.Client();
  Uri uriWorks = Uri(scheme:'http', host:uri.host, port: uri.port,
      path: '/api/client');

  Map ret = {'return': ''};
  try {
    var response = await client.post(uriWorks,
        headers: <String, String> {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
        body: {
      'instanceUuid': UUID,
      'userName': uri.queryParameters['username']
    });
    ret = json.decode(response.body);
  } finally {
    client.close();
  }
  return ret;
}

// ignore: non_constant_identifier_names
Future<Map> deleteHttpJson(Uri uri, String UUID) async {
  var client = http.Client();
  Uri uriWorks = Uri(scheme:'http', host:uri.host, port: uri.port,
      path: '/api/client/$UUID/${uri.queryParameters['username']}');
  Map ret = {'return': ''};
  try {
    var response = await client.delete(uriWorks,
        headers: <String, String> {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: {
          'instanceUuid': UUID,
          'userName': uri.queryParameters['username']
        });
    ret = json.decode(response.body);
  } finally {
    client.close();
  }
  return ret;
}

class MyApp extends StatefulWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  // final SystemTray _systemTray = SystemTray();
  // final AppWindow _appWindow = AppWindow();
  //
  // Timer? _timer;
  // bool _toogleTrayIcon = true;

  @override
  void initState() {
    super.initState();
    // initSystemTray();
  }

  @override
  void dispose() {
    super.dispose();
    // _timer?.cancel();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        //title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: Text('${Status}')
        // Platform.isWindows? ExampleBrowser(uri) : const MyHomePage(),
        );
  }

// Future<void> initSystemTray() async {
//   String path = Platform.isWindows ? 'assets/app_icon.ico' : 'assets/app_icon.png';
//   // _appWindow.hide;
//   final menu = [
//     MenuItem(label: 'Show', onClicked: _appWindow.show),
//     MenuItem(label: 'Hide', onClicked: _appWindow.hide),
//     MenuItem(label: 'Exit', onClicked: _appWindow.close),
//   ];
//
//   // We first init the systray menu and then add the menu entries
//   await _systemTray.initSystemTray(
//     title: "system tray",
//     iconPath: path,
//   );
//
//   await _systemTray.setContextMenu(menu);
//
//   // handle system tray event
//   _systemTray.registerSystemTrayEventHandler((eventName) {
//     debugPrint("eventName: $eventName");
//     if (eventName == "leftMouseDown") {
//     } else if (eventName == "leftMouseUp") {
//       _systemTray.popUpContextMenu();
//     } else if (eventName == "rightMouseDown") {
//     } else if (eventName == "rightMouseUp") {
//       _appWindow.show();
//     }
//   });
// }
}
