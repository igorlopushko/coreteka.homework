// Package color describes console output markup.
package color

import "runtime"

var Reset = "\033[0m"
var Red = "\033[31m"

// Resets values for windows since it does not support markup.
func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
	}
}
