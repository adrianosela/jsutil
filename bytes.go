package jsutil

import (
	"fmt"
	"syscall/js"
)

// JsToBytes converts a Go byte slice to a JavaScript Uint8Array.
func BytesToJs(data []byte) js.Value {
	n := len(data)
	jsValue := js.Global().Get("Uint8Array").New(n)
	js.CopyBytesToJS(jsValue, data[:n])
	return jsValue
}

// JsToBytes converts a JavaScript Uint8Array or Uint8ClampedArray to a Go byte slice.
func JsToBytes(data js.Value) ([]byte, error) {
	if err := AssertTypeEquals(data, js.TypeObject); err != nil {
		return nil, fmt.Errorf("js property type is not compatible: %v", err)
	}
	length := data.Length()
	buf := make([]byte, length)
	copied := js.CopyBytesToGo(buf, data)
	if copied < length {
		return buf, fmt.Errorf("bytes copied less than bytes available (%d < %d)", copied, length)
	}
	return buf, nil
}

// JsCopyBytes loads a Uint8Array or Uint8ClampedArray to a byte slice. It returns the
// number of bytes copied, which will be the minimum of the lengths of data and dst.
func JsCopyBytes(data js.Value, dst []byte) (int, error) {
	if err := AssertTypeEquals(data, js.TypeObject); err != nil {
		return 0, fmt.Errorf("js property type is not compatible: %v", err)
	}
	copied := js.CopyBytesToGo(dst, data)
	return copied, nil
}
