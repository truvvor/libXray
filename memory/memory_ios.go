//go:build ios

package memory

import (
	"runtime/debug"

	"time"
)

const (
	interval  = 30
	maxMemory = 45 * 1024 * 1024
)

func forceFree(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			debug.FreeOSMemory()
		}
	}()
}

func InitForceFree() {
	debug.SetGCPercent(50)
	debug.SetMemoryLimit(maxMemory)
	duration := time.Duration(interval) * time.Second
	forceFree(duration)
}
