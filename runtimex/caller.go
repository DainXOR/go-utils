package runtimex

import (
	"path/filepath"
	"runtime"
)

type callerInfo struct {
	FunctionName string
	FileName     string
	FileLine     int
}

func CallerInfo(depth int) callerInfo {
	pc, file, line, ok := runtime.Caller(depth)

	if !ok {
		return callerInfo{"unknown", "unknown", 0}
	}

	fn := runtime.FuncForPC(pc)
	name := "unknown"
	if fn != nil {
		name = fn.Name()
	}

	return callerInfo{name, filepath.Base(file), line}
}
