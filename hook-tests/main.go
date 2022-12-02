package main

import (
	"fmt"
	"hooks"
	"os"
	"path"
	"runtime"
	"strings"
	o "testmodule/os"

	"github.com/laurentsimon/godep1"
)

type testHook struct{}

func (l *testHook) Getenv(key string) {
	fmt.Println("hook called with ", key)
	fmt.Println("stack info:")
	retrieveCallInfo2()
}

var manager testHook

func main() {
	fmt.Println("Hello World")

	hooks.SetManager(&manager)

	os.Getenv("MYKEY")
	o.Test()

	godep1.TestEnv()

	// retrieveCallInfo2()
}

func retrieveCallInfo2() {
	pc := make([]uintptr, 1000)
	// Skip this function and the runtime.caller itself.
	n := runtime.Callers(2, pc)
	if n == 0 {
		panic("!zero!")
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	rpc := make([]uintptr, n)
	for i := 0; i <= len(pc)-1; i++ {
		rpc[i] = pc[len(pc)-(i+1)]
	}
	frames := runtime.CallersFrames(rpc)
	var prev *runtime.Frame
	for {
		curr, more := frames.Next()

		// Process this frame.
		//
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		// if !strings.Contains(frame.File, "runtime/") {
		// 	break
		// }

		packageName := getPackageName(curr)
		fmt.Printf("- %s | %s | %s:%d \n", packageName, curr.Function, curr.File, curr.Line)

		// Check for package.
		if prev != nil && prev.Function == "main.main" {
			fmt.Println("direct dep is ", curr.Function)
			break
		}
		prev = &curr

		// Check whether there are more frames to process after this one.
		if !more {
			break
		}

	}
}

func getPackageName(frame runtime.Frame) string {
	f := frame.Func
	if f == nil {
		return "__LOCAL__"
	}
	parts := strings.Split(f.Name(), ".")
	pl := len(parts)

	if parts[pl-2][0] == '(' {
		return strings.Join(parts[0:pl-2], ".")
	} else {
		return strings.Join(parts[0:pl-1], ".")
	}
	return "invalid"
}

func retrieveCallInfo() {
	i := 1
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, fileName := path.Split(file)
		parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		pl := len(parts)
		packageName := ""
		funcName := parts[pl-1]

		if parts[pl-2][0] == '(' {
			funcName = parts[pl-2] + "." + funcName
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}

		fmt.Println(i, " ",
			packageName, fileName, funcName, line)
		i++
	}
}
