package keys

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

// KeyMapToSlice converts a struct of key.Binding fields to a slice
// Useful for iterating through all keys in a keymap struct
func KeyMapToSlice(t any) (bindings []key.Binding) {
	typ := reflect.TypeOf(t)
	if typ.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < typ.NumField(); i++ {
		v := reflect.ValueOf(t).Field(i)
		bindings = append(bindings, v.Interface().(key.Binding))
	}
	return
}
