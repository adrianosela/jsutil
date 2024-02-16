package jsutil

import (
	"fmt"
	"syscall/js"
)

// TruthyTypes is a set (slice) of all js.Type
// which can be treated as a boolean.
var TruthyTypes = []js.Type{
	js.TypeUndefined,
	js.TypeNull,
	js.TypeBoolean,
	js.TypeNumber,
	js.TypeString,
	js.TypeSymbol,
	js.TypeFunction,
	js.TypeObject,
}

// GetBoolean retrieves a type-checked boolean from the global scope.
func GetBoolean(expr string) (bool, bool) {
	b, err := GetBooleanE(expr)
	if err != nil {
		return false, false
	}
	return b, true
}

// GetTruthyBoolean retrieves a type-checked boolean from the global scope.
func GetTruthyBoolean(expr string) (bool, bool) {
	b, err := GetTruthyBooleanE(expr)
	if err != nil {
		return false, false
	}
	return b, true
}

// GetBooleanE retrieves a type-checked boolean from the global scope.
func GetBooleanE(expr string) (bool, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return false, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeEquals(jsValue, js.TypeBoolean); err != nil {
		return false, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Bool(), nil
}

// GetTruthyBooleanE retrieves a type-checked (truthy) boolean from the global scope.
func GetTruthyBooleanE(expr string) (bool, error) {
	jsValue, err := GetE(expr)
	if err != nil {
		return false, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = AssertTypeOneOf(jsValue, TruthyTypes...); err != nil {
		return false, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Truthy(), nil
}
