// Copyright 2018 The Flutter team. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// import 'dart:async';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'ablerdp.dart';

// import 'ablewebview.dart';
// import 'package:system_tray/system_tray.dart';
import 'package:bitsdojo_window/bitsdojo_window.dart';

String uri = "";
var reg;
var rdp;
String Status = "";

void main(List<String> args) async {
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
  uri = args.join(" ");
  // var file = File('c:\\ablecloud\\file.txt');
  // var sink = file.openWrite();
  // sink.write('FILE ACCESSED ${DateTime.now()}\n');
  Status += 'FILE ACCESSED ${DateTime.now()}\n';

  ls.writeReg(uri);
  Status += 'reg writed ${DateTime.now()}\n';

  final path = Uri.parse(uri);
  String hashedpassword = 'qwer';
  if (path.queryParameters.containsKey('password')){
    var password = path.queryParameters['password'];
    rdpPassword(password!).then((value) {
      hashedpassword=value;
      Status += 'password $hashedpassword ${DateTime.now()}\n';
      //if(path.queryParameters.containsKey('password 51')){
      Map<String,List<String>> map = Map.from(path.queryParametersAll);
      print("map: $map");
      map.remove('password');
      print("map: $map");
      map['password 51'] = [hashedpassword];
      Uri p = Uri(
        scheme: path.scheme,
        path: path.path,
        queryParameters: map,
        host: path.host
      );
      //uri += '&password 51=$hashedpassword';
      uri = p.toString();
      print("p: $p");
      ls.writeRDP(uri);
      Status += 'rdp writed ${DateTime.now()}\n';

      regAdd(ls);
      Status += 'reg added ${DateTime.now()}\n';

      rdpLaunch(ls);
      Status += 'rdp launched ${DateTime.now()}\n';
      // ls.deleteRDP();
    });
  }
  else{
    ls.writeRDP(uri);
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
