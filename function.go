package jsutil

import (
	"fmt"
	"syscall/js"
)

// Invoke invokes a type-checked function with the given arguments.
func Invoke(fn js.Value, args ...any) (js.Value, error) {
	if fn.IsUndefined() {
		return js.Undefined(), fmt.Errorf("function was undefined at invocation time")
	}
	if fn.Type() != js.TypeFunction {
		return js.Undefined(), fmt.Errorf("js object was not a function at invocation time")
	}
	return fn.Invoke(args...), nil
}

// InvokeFunction invokes a type-checked function from the global scope with the given arguments.
func InvokeFunction(expr string, args ...any) (js.Value, error) {
	fn, err := Get(expr)
	if err != nil {
		return js.Undefined(), fmt.Errorf("could not get function '%s': %v", expr, err)
	}
	return Invoke(fn, args)
}
