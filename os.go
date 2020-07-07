package getOS

import "runtime"

// GetOS ...
func GetOS() string {
	return runtime.GOOS
}
