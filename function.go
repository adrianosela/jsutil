package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetFunction retrieves a type-checked function from the global scope.
func GetFunction(expr string) (js.Value, bool) {
	fn, err := GetFunctionE(expr)
	if err != nil {
		return js.Undefined(), false
	}
	return fn, true
}

// GetFunctionE retrieves a type-checked function from the global scope.
func GetFunctionE(expr string) (js.Value, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return js.Undefined(), fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeFunction); err != nil {
		return js.Undefined(), fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue, nil
}

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
	fn, err := GetE(expr)
	if err != nil {
		return js.Undefined(), fmt.Errorf("could not get function '%s': %v", expr, err)
	}
	return Invoke(fn, args)
}
