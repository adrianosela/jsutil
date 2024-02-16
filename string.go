package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetString retrieves a type-checked string from the global scope.
func GetString(expr string) (string, bool) {
	str, err := GetStringE(expr)
	if err != nil {
		return "", false
	}
	return str, true
}

// GetStringE retrieves a type-checked string from the global scope.
func GetStringE(expr string) (string, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return "", fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeString); err != nil {
		return "", fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.String(), nil
}
