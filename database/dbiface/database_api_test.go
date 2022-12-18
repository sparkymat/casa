package dbiface_test

import (
	"reflect"
	"testing"

	"git.orion.home/oxhead/casa/database"
	"git.orion.home/oxhead/casa/database/dbiface"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	structType := reflect.TypeOf(&database.Service{})
	interfaceType := reflect.TypeOf((*dbiface.DatabaseAPI)(nil)).Elem()

	interfaceMethodMap := map[string]interface{}{}

	for i := 0; i < interfaceType.NumMethod(); i++ {
		interfaceMethodMap[interfaceType.Method(i).Name] = struct{}{}
	}

	for i := 0; i < structType.NumMethod(); i++ {
		if _, found := interfaceMethodMap[structType.Method(i).Name]; !found {
			t.Errorf("'%s' not found on interface '%s'", structType.Method(i).Name, interfaceType.Name())
		}
	}
}
