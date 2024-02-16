package jsutils

import (
	"fmt"
	"strings"
	"syscall/js"
)

// TruthyTypes is a set (slice) of all js.Type which can be treated as a boolean.
var TruthyTypes = []js.Type{
	js.TypeUndefined, js.TypeNull, js.TypeBoolean, js.TypeNumber,
	js.TypeString, js.TypeSymbol, js.TypeFunction, js.TypeObject,
}

// Function represents a JavaScript function.
type Function struct {
	fn js.Value
}

// Invoke invokes a JavaScript function with the given arguments.
func (f Function) Invoke(args ...any) js.Value {
	return f.fn.Invoke(args...)
}

// Get retrieves an expression from the global scope.
func Get(expr string) (js.Value, error) {
	return GetFromScope(js.Global(), expr)
}

// GetFromScope retrieves an expression from the given scope.
func GetFromScope(scope js.Value, expr string) (js.Value, error) {
	value := scope
	for _, part := range strings.Split(expr, ".") {
		if value.IsUndefined() {
			return js.Undefined(), fmt.Errorf("cannot read properties of undefined (reading '%s' during parsing of '%s')", part, expr)
		}
		value = value.Get(part)
	}
	return value, nil
}

// Invoke invokes a type-checked function from the global scope with the given arguments.
func Invoke(expr string, args ...any) (js.Value, error) {
	fn, err := GetFunction(expr)
	if err != nil {
		return js.Undefined(), fmt.Errorf("could not get function '%s': %v", expr, err)
	}
	return fn.Invoke(args...), nil
}

// GetInt retrieves an type-checked integer from the global scope.
func GetInt(expr string) (int, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = assertType(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Int(), nil
}

// GetFloat retrieves a type-checked float from the global scope.
func GetFloat(expr string) (float64, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return 0, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = assertType(jsValue, js.TypeNumber); err != nil {
		return 0, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Float(), nil
}

// GetString retrieves a type-checked string from the global scope.
func GetString(expr string) (string, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return "", fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = assertType(jsValue, js.TypeString); err != nil {
		return "", fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.String(), nil
}

// GetBoolean retrieves a type-checked boolean from the global scope.
func GetBoolean(expr string) (bool, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return false, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = assertType(jsValue, js.TypeBoolean); err != nil {
		return false, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Bool(), nil
}

// GetFunction retrieves a type-checked function from the global scope.
func GetFunction(expr string) (*Function, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return nil, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if jsValue.IsUndefined() {
		return nil, fmt.Errorf("js object '%s' is undefined", expr)
	}
	if err = assertType(jsValue, js.TypeFunction); err != nil {
		return nil, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return &Function{jsValue}, nil
}

// GetFunction retrieves a type-checked (truthy) boolean from the global scope.
func GetTruthyBoolean(expr string) (bool, error) {
	jsValue, err := Get(expr)
	if err != nil {
		return false, fmt.Errorf("could not get js object '%s': %v", expr, err)
	}
	if err = assertTypes(jsValue, TruthyTypes...); err != nil {
		return false, fmt.Errorf("js object type is not compatible: %v", err)
	}
	return jsValue.Truthy(), nil
}

func assertType(jsValue js.Value, jsType js.Type) error {
	if jsValue.Type() != jsType {
		return fmt.Errorf("not a %s", jsType.String())
	}
	return nil
}

func assertTypes(jsValue js.Value, jsTypes ...js.Type) error {
	for _, jsType := range jsTypes {
		if jsValue.Type() == jsType {
			return nil
		}
	}
	return fmt.Errorf("not a any of %v", jsTypes)
}
