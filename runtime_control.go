package libXray

import "runtime/debug"

func LibXrayFreeOSMemory() {
	debug.FreeOSMemory()
}

func LibXraySetMemoryLimit(bytes int64) {
	debug.SetMemoryLimit(bytes)
}

func LibXraySetGCPercent(percent int) {
	debug.SetGCPercent(percent)
}
