package runtimex

import (
	"path/filepath"
	"runtime"
)

type funcName = string
type fileName = string
type fileLine = int

func CallerInfo(depth int) (funcName, fileName, fileLine) {
	pc, file, line, ok := runtime.Caller(depth)

	if !ok {
		return "unknown", "unknown", 0
	}

	fn := runtime.FuncForPC(pc)
	name := "unknown"
	if fn != nil {
		name = fn.Name()
	}

	return name, filepath.Base(file), line
}
