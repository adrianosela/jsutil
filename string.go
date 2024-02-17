package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetString retrieves a type-checked string from the global scope.
func GetString(expr string) (string, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return "", fmt.Errorf("could not get js property '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeString); err != nil {
		return "", fmt.Errorf("js property type is not compatible: %v", err)
	}
	return jsValue.String(), nil
}
