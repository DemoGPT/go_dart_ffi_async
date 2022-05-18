import 'dart:ffi';
import 'dart:isolate';
import 'package:ffi/ffi.dart';

typedef StartWorkType = Void Function(Int64 port, Pointer<Utf8> goFuncName);
typedef StartWorkFunc = void Function(int port, Pointer<Utf8> goFuncName);

class Work extends Struct {
  Pointer<Utf8> name;
}

class NativeGoFunc {
  static const String goInit = 'goInit';
  static const String goNetEnv = 'goNetEnv';
}

void main() async {
  print('-----------111--------');
  testGo().then((value) {
    print('------$value---------');
    final utf = Pointer<Utf8>.fromAddress(value as int);
    print('------${utf.toDartString()}---------');

    // final work = Pointer<Work>.fromAddress(value as int);
    // print(work.ref.name.toDartString());
  });
  print('-----------222--------');
}

Future<dynamic> testGo() {
  final lib = DynamicLibrary.open("./godart.so");

  final initializeApi =
      lib.lookupFunction<IntPtr Function(Pointer<Void>), int Function(Pointer<Void>)>(
          "InitializeDartApi");
  if (initializeApi(NativeApi.initializeApiDLData) != 0) {
    throw "Failed to initialize Dart API";
  }

  // final interactiveCppRequests = ReceivePort()
  //   ..listen((data) {
  //     print('Received: ${data} from Go');
  //     receivePort.close();
  //     return data;
  //   });
  // final int nativePort = interactiveCppRequests.sendPort.nativePort;

  final ReceivePort receivePort = ReceivePort();
  final int nativePort = receivePort.sendPort.nativePort;

  print("nativePort: ${nativePort}");

  final StartWorkFunc startWork =
      lib.lookup<NativeFunction<StartWorkType>>("StartWork").asFunction();

  startWork(nativePort, NativeGoFunc.goInit.toNativeUtf8());
  return receivePort.first;
}
