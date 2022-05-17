package dart_api_dl

// #include "stdint.h"
// #include "stdlib.h"
// #include "include/dart_api_dl.c"
//
// // Go does not allow calling C function pointers directly. So we are
// // forced to provide a trampoline.
// bool GoDart_PostCObject(Dart_Port_DL port, Dart_CObject* obj) {
//   return Dart_PostCObject_DL(port, obj);
// }
//
//	typedef struct WorkStruct{
//		char *str;
//	}WorkStruct;
//
//	int64_t GetWork(void **ppWork, char* str) {
//		WorkStruct *pWork = (WorkStruct *)malloc(sizeof(WorkStruct));
//		pWork->str=str;
//
//		*ppWork = pWork;
//
//		int64_t ptr = (int64_t)pWork;
//
//		return ptr;
//	}
import "C"
import "unsafe"

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, str string) {
    var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64

	var pWork unsafe.Pointer
	ptrAddress := C.GetWork(&pWork, C.CString(str))

	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrAddress
	C.GoDart_PostCObject(C.int64_t(port), &obj)


// 	var obj C.Dart_CObject
// 	obj._type = C.Dart_CObject_kInt64
// 	// cgo does not support unions so we are forced to do this
// 	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = C.int64_t(msg)
//  C.GoDart_PostCObject(C.int64_t(port), &obj)

}
