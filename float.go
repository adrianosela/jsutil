package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetFloat retrieves a type-checked float from the global scope.
func GetFloat(expr string) (float64, bool) {
	float, err := GetFloatE(expr)
	if err != nil {
		return 0, false
	}
	return float, true
}

// GetFloatE retrieves a type-checked float from the global scope.
func GetFloatE(expr string) (float64, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Float(), nil
}
