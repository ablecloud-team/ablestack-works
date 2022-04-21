// import 'dart:convert';

import 'dart:convert';
import 'dart:io';

// import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:path_provider/path_provider.dart';

import 'package:shell/shell.dart';

class LocalStorage {
  Future<String> get _localPath async {
    final directory = await getTemporaryDirectory();

    return directory.path;
  }

  Future<File> get _localRDPFile async {
    final path = await _localPath;
    return File('$path/tmp.rdp');
  }

  Future<File> get _localRegFile async {
    final path = await _localPath;
    return File('$path/tmp.reg');
  }

  void writeReg(String path) {
    _localRegFile.then((file) {
      String hashString = "";
      String addrString = "";
      String userString = "";
      String domString = "";
      String regString = "";
      if (Platform.isWindows) {
        final uri = Uri.parse(path);
        for (var query in uri.queryParameters.entries) {
          // if (kDebugMode) {
          //   print("${query.key}: ${query.value}");
          // }
          switch (query.key) {
            case "hash":
              hashString = query.value;
              break;
            case "full address":
              addrString = query.value;
              break;
            case "username":
              userString = query.value;
              break;
            case "domain":
              domString = query.value;
              break;
          }
        }
        String hasht = hashString;
        String hash2 = "";
        while (hasht.length > 2) {
          hash2 += hasht.substring(0, 2) + ",";
          hasht = hasht.substring(2);
        }
        hash2 += hasht;
        if (uri.queryParameters.containsKey("hash")) {
          regString = """
            Windows Registry Editor Version 5.00
            
            [HKEY_CURRENT_USER\\Software\\Microsoft\\Terminal Server Client\\LocalDevices]
            "$addrString"=dword:0000004c
            
            [HKEY_CURRENT_USER\\Software\\Microsoft\\Terminal Server Client\\Servers\\$addrString]
            @=""
            "CertHash"=hex:$hash2
            "UsernameHint"="$userString@$domString"
        """
              .trim();
        }
        // String cmd0 =
        //     "add \"HKCU\\Software\\Microsoft\\Terminal Server Client\\LocalDevices\" /f /v \"$addrString\" /d REG_DWORD 0000004c";
        // String cmd1 =
        //     "add \"HKCU\\Software\\Microsoft\\Terminal Server Client\\Servers\\$addrString\" /f ";
        // String cmd2 =
        //     "add \"HKCU\\Software\\Microsoft\\Terminal Server Client\\Servers\\$addrString\" /f /v CertHash /t REG_BINARY /d $hashString";
        // String cmd3 =
        //     "add \"HKCU\\Software\\Microsoft\\Terminal Server Client\\Servers\\$addrString\" /f /v UsernameHint /t REG_SZ /d $userString@$domString";
        // String cmd4 =
        //     "import ${_localRegFile.then((value) => value.absolute)}";
      }
      // 파일 쓰기
      file.writeAsString(regString);
    });
    return;
  }

  void writeRDP(String path) {
    _localRDPFile.then((file) {
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
      if (Platform.isWindows) {
        final uri = Uri.parse(path);
        for (var query in uri.queryParameters.entries) {
          // if (kDebugMode) {
          //   print("${query.key}: ${query.value}");
          // }
          if (query.key == "password 51") {
            rdpstring += "${query.key}:b:${query.value}\n";
          } else {
            try {
              var i = int.parse(query.value, radix: 10);
              // if (kDebugMode) {
              //   print("i: $i");
              // }
              rdpstring += "${query.key}:i:${query.value}\n";
            } catch (e) {
              // if (kDebugMode) {
              //   print("s: $e");
              //   print("s: $query.value");
              // }
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
          //   print("uri: $uri");
          //   print("scheme: " + uri.scheme);
          //   print("path: " + uri.path);
          //   print("query: " + uri.query);
          print("queries: " + uri.queryParameters.toString());
          //   print("rdpstring: " + rdpstring);
        }
      }
      // 파일 쓰기
      file.writeAsString(rdpstring);
    });
    return;
  }

  Future<bool> deleteRDP() async {
    final file = await _localRDPFile;
    try {
      file.delete();
      if (kDebugMode) {
        print("deleted rdp");
      }
      return true;
    } catch (e) {
      if (kDebugMode) {
        print("delete rdp failed");
      }
      return false;
    }
  }

  Future<bool> deleteReg() async {
    final file = await _localRegFile;
    try {
      file.delete();
      if (kDebugMode) {
        print("deleted reg");
      }
      return true;
    } catch (e) {
      if (kDebugMode) {
        print("delete reg failed");
      }
      return false;
    }
  }
}

void rdpLaunch(LocalStorage ls) {
  //var fs = const LocalFileSystem();
  var shell = Shell();

  if (Platform.isWindows) {
    //windows

    ls._localRDPFile.then((file) {
      var path = file.path;
      shell.start('mstsc', arguments: [path]).then((open) {
        //var ret =
        open.expectExitCode(Iterable<int>.generate(255).toList());
        //var pwd =
        open.stdout.readAsString();
        if (kDebugMode) {
          print("ret: ${open.exitCode}");
        }

        ls.deleteRDP().then((value){
            exit(0);
        });

        // print('cwd: $pwd, ret: $ret');
      });
    });
  } else if (Platform.isAndroid) {
    //Android
    return;
  }
}

void regAdd(LocalStorage ls) {
  //var fs = const LocalFileSystem();
  // var shell = Shell();

  if (Platform.isWindows) {
    //windows

    ls._localRegFile.then((file) {
      var path = file.path;
      // var open = await Future.wait([
      //   Process.run('C:\\Windows\\System32\\reg.exe', ['import', path],
      // stdoutEncoding: const Utf8Codec(allowMalformed: true),
      // stderrEncoding: const Utf8Codec(allowMalformed: true)).then((value) => ls.deleteReg())
      // ]);
      Process.run('C:\\Windows\\System32\\reg.exe', ['import', path],
              stdoutEncoding: const Utf8Codec(allowMalformed: true),
              stderrEncoding: const Utf8Codec(allowMalformed: true),
              runInShell: true)
          .then((open) {
        // open.then((value) => ls.deleteReg());
        // var ret =
        //await open[0].expectExitCode(Iterable<int>.generate(255).toList());
        // await ls.deleteReg();

        if (kDebugMode) {
          print("stdout reg: ${open.stdout}");
          print("stderr reg: ${open.stderr}");
          print("arg reg: ${open.stderr}");
          // print("str reg: ${open.toString()}");

          // print("ret reg: ${open.exitCode}");
        }
      });
      // print('cwd: $pwd, ret: $ret');
    });
  } else if (Platform.isAndroid) {
    //Android
    return;
  }
}

Future<String> rdpPassword(String password) async {
  var shell = Shell();
  String ret = '';

  if (Platform.isWindows) {
    //windows
    var open = await Process.run('%SystemRoot%\\system32\\WindowsPowerShell\\v1.0\\powershell.exe', ["(\"$password\" | ConvertTo-SecureString -AsPlainText -Force) | ConvertFrom-SecureString;"],
        stdoutEncoding: const Utf8Codec(allowMalformed: true),
        stderrEncoding: const Utf8Codec(allowMalformed: true),
        runInShell: true);
        // .then((open) {
      ret =  open.stdout;
      print("password1: $ret");
    // });
    return ret;

  }
  if (kDebugMode) {
    print("password: $ret");
  }

  return "not yet";
}
