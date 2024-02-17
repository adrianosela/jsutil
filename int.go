package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetInt retrieves an type-checked integer from the global scope.
func GetInt(expr string) (int, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js property '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js property type is not compatible: %v", err)
	}
	return jsValue.Int(), nil
}
