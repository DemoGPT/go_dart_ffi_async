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
//		char *name;
//		int64_t age;
//	}WorkStruct;
//
//	int64_t GetWork(void **ppWork, char* name, int64_t age) {
//		WorkStruct *pWork = (WorkStruct *)malloc(sizeof(WorkStruct));
//		pWork->name=name;
//		pWork->age=age;
//
//		*ppWork = pWork;
//
//		int64_t ptr = (int64_t)pWork;
//
//		return ptr;
//	}
//
//	void clearWorkStructMemory(WorkStruct pWork) {
//		free(pWork.name);
//		free(&pWork.age);
//		free(&pWork);
//	}
//
//
//  int64_t GetMessage(void **ppStr, char* msg, int64_t length){
//		// Allocates native memory in C.
//		char *reversed_str = (char *)malloc((length + 1) * sizeof(char));
//		for (int i = 0; i < length; i++){
//			reversed_str[length - i - 1] = msg[i];
//		}
//		reversed_str[length] = '\0';
//
//		*ppStr = reversed_str;
//		int64_t ptr = (int64_t)reversed_str;
//
//		return ptr;
//	}
//
//	void clearMessageMemory(char* str) {
//		free(str);
//	}
import "C"
import "unsafe"

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, msg string) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64

	var pWork unsafe.Pointer
	ptrAddress := C.GetWork(&pWork, C.CString(msg), C.int64_t(26))

	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrAddress
	C.GoDart_PostCObject(C.int64_t(port), &obj)
}

func FreeWorkStructMemory(pointer *int64) {
	ptr := (*C.struct_WorkStruct)(unsafe.Pointer(pointer))
	C.clearWorkStructMemory(*ptr)
}

//
//
//

func SendToPortStr(port int64, msg string) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64

	var pStr unsafe.Pointer
	ptrStr := C.GetMessage(&pStr, C.CString(msg), C.int64_t(len(msg)))

	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrStr
	C.GoDart_PostCObject(C.int64_t(port), &obj)

}

func FreeMessageMemory(pointer *int64) {
	ptr := (*C.char)(unsafe.Pointer(pointer))
	C.clearMessageMemory(ptr)
}

//
//
//

func SendToPortInt(port int64, msg int64) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64
	// cgo does not support unions so we are forced to do this
	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = C.int64_t(msg)
	C.GoDart_PostCObject(C.int64_t(port), &obj)
}
