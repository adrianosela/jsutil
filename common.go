package jsutil

import (
	"fmt"
	"strings"
	"syscall/js"
)

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

// Get retrieves an expression from the global scope.
func Get(expr string) (js.Value, error) {
	return GetFromScope(js.Global(), expr)
}

// AssertTypeEquals returns nil if a given JavaScript value conforms to the given type.
func AssertTypeEquals(jsValue js.Value, jsType js.Type) error {
	if jsValue.Type() != jsType {
		return fmt.Errorf("not a %s", jsType.String())
	}
	return nil
}

// AssertTypeOneOf returns nil if a given JavaScript value conforms to one the given types
func AssertTypeOneOf(jsValue js.Value, jsTypes ...js.Type) error {
	for _, jsType := range jsTypes {
		if jsValue.Type() == jsType {
			return nil
		}
	}
	return fmt.Errorf("not any of %v", jsTypes)
}
