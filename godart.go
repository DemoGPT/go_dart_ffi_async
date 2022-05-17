package main

import (
	"C"
	"fmt"
	"github.com/mraleph/go_dart_ffi_example/dart_api_dl"
	"time"
	"unsafe"
)

//export InitializeDartApi
func InitializeDartApi(api unsafe.Pointer) {
	dart_api_dl.Init(api)
}

//export StartWork
func StartWork(port int64, goFuncName *C.char) {
    funcName := C.GoString(goFuncName)
	fmt.Println("Go: Starting some asynchronous work")
	switch funcName {
        case "goInit":
            go goInit(port)
        case "goNetEnv":
            go goNetEnv(port)
        default:
            fmt.Printf("GO: goFuncName[%s] does not exist",funcName)
            dart_api_dl.SendToPort(port,  "default")
    }

	fmt.Println("Go: Returning to Dart")
}

func goInit(port int64) {
	var counter int64 = 50
    time.Sleep(2 * time.Second)
    fmt.Printf("GO: 2 seconds passed, goInit back: %d\n",counter)
    dart_api_dl.SendToPort(port,  "goInit")
}

func goNetEnv(port int64) {
    var counter int64 = 60
    time.Sleep(2 * time.Second)
    fmt.Printf("GO: 2 seconds passed, goNetEnv back: %d\n",counter)
    dart_api_dl.SendToPort(port,  "goNetEnv")
}



// func work(port int64) {
// 	var counter int64
// 	for {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("GO: 2 seconds passed")
// 		counter++
// 		dart_api_dl.SendToPort(port, counter)
// 	}
// }

// Unused
func main() {}
