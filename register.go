package jsutil

import "syscall/js"

// SetGlobalProperties sets propreties on the global object.
func SetGlobalProperties(kv map[string]any) {
	scope := js.Global()
	for k, v := range kv {
		safely(func() { scope.Set(k, js.ValueOf(v)) })
	}
}

func safely(fn func(), onPanic ...func(r any)) {
	defer func() {
		if r := recover(); r != nil {
			for _, x := range onPanic {
				x(r)
			}
		}
	}()

	fn()
}
