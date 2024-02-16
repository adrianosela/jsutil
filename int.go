package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetInt retrieves an type-checked integer from the global scope.
func GetInt(expr string) (int, bool) {
	i, err := GetIntE(expr)
	if err != nil {
		return 0, false
	}
	return i, true
}

// GetIntE retrieves an type-checked integer from the global scope.
func GetIntE(expr string) (int, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Int(), nil
}
