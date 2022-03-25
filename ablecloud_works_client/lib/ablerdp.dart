// import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:path_provider/path_provider.dart';

import 'package:shell/shell.dart';
import 'dart:io';

class LocalStorage {
  Future<String> get _localPath async {
    final directory = await getApplicationDocumentsDirectory();

    return directory.path;
  }

  Future<File> get _localFile async {
    final path = await _localPath;
    return File('$path/tmp.rdp');
  }

  Future<File> writeRDP(String path) async {
    final file = await _localFile;
    // String password = path;
    var rdpstring = '''screen mode id:i:2
use multimon:i:0
desktopwidth:i:1884
desktopheight:i:1027
session bpp:i:32
winposstr:s:0,1,220,288,1020,888
compression:i:1
keyboardhook:i:2
audiocapturemode:i:0
videoplaybackmode:i:1
connection type:i:7
networkautodetect:i:1
bandwidthautodetect:i:1
displayconnectionbar:i:1
enableworkspacereconnect:i:0
disable wallpaper:i:0
allow font smoothing:i:0
allow desktop composition:i:0
disable full window drag:i:1
disable menu anims:i:1
disable themes:i:0
disable cursor setting:i:0
bitmapcachepersistenable:i:1
audiomode:i:0
redirectprinters:i:1
redirectlocation:i:0
redirectcomports:i:0
redirectsmartcards:i:1
redirectclipboard:i:1
redirectposdevices:i:0
autoreconnection enabled:i:1
authentication level:i:2
prompt for credentials:i:1
negotiate security layer:i:1
remoteapplicationmode:i:0
alternate shell:s:
shell working directory:s:
gatewayhostname:s:
gatewayusagemethod:i:4
gatewaycredentialssource:i:4
gatewayprofileusagemethod:i:0
promptcredentialonce:i:0
gatewaybrokeringtype:i:0
use redirection server name:i:0
rdgiskdcproxy:i:0
kdcproxyname:s:
prompt for credentials:i:0
''';
    var shell = Shell();
    var stdout2 = '';
    var stderr2 = '';
    if (Platform.isWindows) {
      final uri = Uri.parse(path);
      for (var query in uri.queryParameters.entries) {
        print("${query.key}: ${query.value}");
        if (query.key == "password 51") {
          rdpstring += "${query.key}:b:${query.value}\n";
        } else {
          try {
            var i = int.parse(query.value, radix: 10);
            print("i: $i");
            rdpstring += "${query.key}:i:${query.value}\n";
          } catch (e) {
            print("s: $e");
            print("s: $query.value");
            rdpstring += "${query.key}:s:${query.value}\n";
          }
        }
        // if (!int.parse(query.value, radix: 16).isNaN){
        //   rdpstring += "${query.key}:b:${query.value}\n";
        // } else if (!int.parse(query.value, radix: 10).isNaN){
        //   rdpstring += "${query.key}:i:${query.value}";
        // } else{
        //   rdpstring += "${query.key}:s:${query.value}";
        // }
      }
      if (kDebugMode) {
        print("uri: $uri");
        print("scheme: " + uri.scheme);
        print("path: " + uri.path);
        print("query: " + uri.query);
        print("queries: " + uri.queryParameters.toString());
        print("rdpstring: " + rdpstring);
      }
    }
    // 파일 쓰기
    return file.writeAsString(rdpstring);
  }

  Future<bool> deleteRDP() async {
    final file = await _localFile;
    try {
      file.delete();
      print("deleted");
      return true;
    }
    catch(e){
      print("delete failed");
      return false;
    }

  }
}

Future<void> rdpLaunch(LocalStorage ls) async {
  //var fs = const LocalFileSystem();
  var shell = Shell();

  if (Platform.isWindows) {
    //windows

    final file = await ls._localFile;
    var path = file.path;
    var open = await Future.wait([shell.start('mstsc', arguments: [path])]);

    var ret = await open[0].expectExitCode(Iterable<int>.generate(255).toList());
    var pwd = await open[0].stdout.readAsString();
    print("ret: ${await open[0].exitCode}");
    await ls.deleteRDP();


    // print('cwd: $pwd, ret: $ret');


  } else if (Platform.isAndroid) {
    //Android
    return;
  }
}
