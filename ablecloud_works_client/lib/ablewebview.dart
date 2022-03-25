
import 'package:flutter/services.dart';
import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';
import 'package:webview_windows/webview_windows.dart';

import 'main.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Welcome to Flutter'),
      ),
      body: const SafeArea(
        child: WebView(
          initialUrl: 'https://ablecloud.io',
          javascriptMode: JavascriptMode.unrestricted,
        ),
      ),
    );
  }
}
final navigatorKey = GlobalKey<NavigatorState>();

class ExampleBrowser extends StatefulWidget {
  const ExampleBrowser(String uri, {Key? key}) : super(key: key);

  @override
  State<ExampleBrowser> createState() => _ExampleBrowser();
}

class _ExampleBrowser extends State<ExampleBrowser> {
  final _controller = WebviewController();
  final _textController = TextEditingController();
  bool _isWebviewSuspended = true;

  @override
  void initState() {
    super.initState();
    initPlatformState();
  }

  Future<void> initPlatformState() async {
    // Optionally initialize the webview environment using
    // a custom user data directory
    // and/or a custom browser executable directory
    // and/or custom chromium command line flags
    //await WebviewController.initializeEnvironment(
    //    additionalArguments: '--show-fps-counter');

    try {
      await _controller.initialize();

      await _controller.setBackgroundColor(Colors.transparent);
      await _controller.setPopupWindowPolicy(WebviewPopupWindowPolicy.deny);
      await _controller.loadUrl('https://works.ablecloud.io');

      if (!mounted) return;
      await _controller.suspend();
      setState(() {});
    } on PlatformException catch (e) {
      WidgetsBinding.instance?.addPostFrameCallback((_) {
        showDialog(
            context: context,
            builder: (_) => AlertDialog(
              title: const Text('Error'),
              content: Column(
                mainAxisSize: MainAxisSize.min,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text('Code: ${e.code}'),
                  Text('Message: ${e.message}'),
                ],
              ),
              actions: [
                TextButton(
                  child: const Text('Continue'),
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                )
              ],
            ));
      });
    }
  }

  Widget compositeView() {
    if (!_controller.value.isInitialized) {
      return const Text(
        'Not Initialized',
        style: TextStyle(
          fontSize: 24.0,
          fontWeight: FontWeight.w900,
        ),
      );
    } else {
      return Stack(
        children: [
          Webview(
            _controller,
            permissionRequested: _onPermissionRequested,
          ),
          StreamBuilder<LoadingState>(
              stream: _controller.loadingState,
              builder: (context, snapshot) {
                if (snapshot.hasData &&
                    snapshot.data == LoadingState.loading) {
                  return const LinearProgressIndicator();
                } else {
                  return const SizedBox();
                }
              }),
        ],
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        tooltip: _isWebviewSuspended ? 'Resume webview' : 'Suspend webview',
        onPressed: () async {
          if (_isWebviewSuspended) {
            await _controller.resume();
          } else {
            await _controller.suspend();
          }
          setState(() {
            _isWebviewSuspended = !_isWebviewSuspended;
          });
        },
        child: _isWebviewSuspended ? const Text('connect') : const Icon( Icons.pause),
      ),
      appBar: AppBar(
          title: Text(uri)
          // StreamBuilder<String>(
          //   stream: _controller.title,
          //   builder: (context, snapshot) {
          //     return Text(
          //         snapshot.hasData ? snapshot.data! : 'WebView (Windows) Example');
          //   },
          // )
      ),
      body: Center(
        child: compositeView(),
      ),
    );
  }

  Future<WebviewPermissionDecision> _onPermissionRequested(
      String url, WebviewPermissionKind kind, bool isUserInitiated) async {
    final decision = await showDialog<WebviewPermissionDecision>(
      context: navigatorKey.currentContext!,
      builder: (BuildContext context) => AlertDialog(
        title: const Text('WebView permission requested'),
        content: Text('WebView has requested permission \'$kind\''),
        actions: <Widget>[
          TextButton(
            onPressed: () =>
                Navigator.pop(context, WebviewPermissionDecision.deny),
            child: const Text('Deny'),
          ),
          TextButton(
            onPressed: () =>
                Navigator.pop(context, WebviewPermissionDecision.allow),
            child: const Text('Allow'),
          ),
        ],
      ),
    );

    return decision ?? WebviewPermissionDecision.none;
  }
}