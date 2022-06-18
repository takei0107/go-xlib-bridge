package xlib

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

type Display C.Display

func XOpenDisplay() *Display {
	display := C.XOpenDisplay(nil)
	return (*Display)(display)
}
