// Copyright 2018 The Flutter team. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import 'dart:async';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'ablerdp.dart';
// import 'ablewebview.dart';
// import 'package:system_tray/system_tray.dart';
import 'package:bitsdojo_window/bitsdojo_window.dart';

String uri="";

void main(List<String> args) async {
  WidgetsFlutterBinding.ensureInitialized();
  // runApp(const MyApp());
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
  ls.writeRDP(uri);
  rdpLaunch(ls).then((value) => exit(0));

  if (kDebugMode) {
    print(args);
  }
  // for (var val in args) {
  //   uri += val + "
  // }
}

//
// class MyApp extends StatefulWidget {
//   const MyApp({Key? key}) : super(key: key);
//
//   @override
//   State<MyApp> createState() => _MyAppState();
// }
//
// class _MyAppState extends State<MyApp> {
//   final SystemTray _systemTray = SystemTray();
//   final AppWindow _appWindow = AppWindow();
//
//   Timer? _timer;
//   bool _toogleTrayIcon = true;
//
//   @override
//   void initState() {
//     super.initState();
//     initSystemTray();
//   }
//
//   @override
//   void dispose() {
//     super.dispose();
//     _timer?.cancel();
//   }
//
//   @override
//   Widget build(BuildContext context) {
//     return MaterialApp(
//       //title: 'Flutter Demo',
//       theme: ThemeData(
//         primarySwatch: Colors.blue,
//       ),
//       home: Platform.isWindows? ExampleBrowser(uri) : const MyHomePage(),
//     );
//   }
//
//   Future<void> initSystemTray() async {
//     String path = Platform.isWindows ? 'assets/app_icon.ico' : 'assets/app_icon.png';
//     _appWindow.hide;
//     final menu = [
//       MenuItem(label: 'Show', onClicked: _appWindow.show),
//       MenuItem(label: 'Hide', onClicked: _appWindow.hide),
//       MenuItem(label: 'Exit', onClicked: _appWindow.close),
//     ];
//
//     // We first init the systray menu and then add the menu entries
//     await _systemTray.initSystemTray(
//       title: "system tray",
//       iconPath: path,
//     );
//
//     await _systemTray.setContextMenu(menu);
//
//     // handle system tray event
//     _systemTray.registerSystemTrayEventHandler((eventName) {
//       debugPrint("eventName: $eventName");
//       if (eventName == "leftMouseDown") {
//       } else if (eventName == "leftMouseUp") {
//         _systemTray.popUpContextMenu();
//       } else if (eventName == "rightMouseDown") {
//       } else if (eventName == "rightMouseUp") {
//         _appWindow.show();
//       }
//     });
//   }
// }
