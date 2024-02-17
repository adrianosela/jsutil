package jsutil

import (
	"fmt"
	"syscall/js"
)

// GetFloat retrieves a type-checked float from the global scope.
func GetFloat(expr string) (float64, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js property '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js property type is not compatible: %v", err)
	}
	return jsValue.Float(), nil
}
